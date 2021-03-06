package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ssopb "edu/api/sso/v1"
	"edu/pkg/ecode"
	"edu/pkg/meta"
	"edu/pkg/tools"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"gopkg.in/yaml.v2"
)

// Config is the identify config model.
type Config struct {
	// csrf switch.
	DisableCSRF bool
	WhiteList   []string
}

type TokenList struct {
	List []string
}

// Auth is the authorization middleware
type Auth struct {
	ssoClient ssopb.SsoClient
	conf      *Config
	log       *log.Helper
	whiteList map[string]bool
}

// authFunc will return mid and error by given context
type authFunc func(context.Context, *transporthttp.Transport) (*meta.DataPermission, error)

var _defaultConf = &Config{
	DisableCSRF: false,
}

// New is used to create an authorization middleware
func New(conf *Config, logger log.Logger, ssoClient ssopb.SsoClient) *Auth {
	if conf == nil {
		conf = _defaultConf
	}
	auth := &Auth{
		conf:      conf,
		log:       log.NewHelper(log.With(logger, "module", "middleware/auth")),
		ssoClient: ssoClient,
		whiteList: make(map[string]bool),
	}

	f, err := tools.Open("../../configs/tokenlist.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var tl TokenList
	if err = yaml.Unmarshal(bytes, &tl); err != nil {
		panic(err)
	}
	for i := 0; i < len(tl.List); i++ {
		url := tl.List[i]
		auth.whiteList[url] = true
	}
	return auth
}

// User is used to mark path as access required.
// If `access_token` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) User(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		tp, ok := transport.FromServerContext(ctx)
		if ok {
			if tp.Kind() == transport.KindHTTP {
				h := tp.(*transporthttp.Transport)
				if a.whiteList[h.Request().URL.Path] {
					a.log.Info("----------------------------------auth")
					if h.RequestHeader().Get("access_token") == "" {
						ctx, err = a.UserWeb(ctx, h)
						if err != nil {
							return
						}
					} else {
						ctx, err = a.UserMobile(ctx, h)
						if err != nil {
							return
						}
					}
				}
			}
		}
		return handler(ctx, req)
	}
}

// UserWeb is used to mark path as web access required.
func (a *Auth) UserWeb(ctx context.Context, info *transporthttp.Transport) (context.Context, error) {
	return a.midAuth(ctx, info, a.authBearerAuth)
}

// UserMobile is used to mark path as mobile access required.
func (a *Auth) UserMobile(ctx context.Context, info *transporthttp.Transport) (context.Context, error) {
	return a.midAuth(ctx, info, a.authToken)
}

// Guest is used to mark path as guest policy.
// If `access_token` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) Guest(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		tp, ok := transport.FromServerContext(ctx)
		if ok {
			if tp.Kind() == transport.KindHTTP {
				h := tp.(*transporthttp.Transport)
				if h.RequestHeader().Get("access_token") == "" {
					a.GuestWeb(ctx, h)
					return
				}
				a.GuestMobile(ctx, h)
			}
		}
		return handler(ctx, req)
	}
}

// GuestWeb is used to mark path as web guest policy.
func (a *Auth) GuestWeb(ctx context.Context, info *transporthttp.Transport) (context.Context, error) {
	a.log.Info("http: [%s] [%s]", info.Request().Method, info.Request().URL.Path)
	return a.guestAuth(ctx, info, a.authCookie)
}

// GuestMobile is used to mark path as mobile guest policy.
func (a *Auth) GuestMobile(ctx context.Context, info *transporthttp.Transport) (context.Context, error) {
	a.log.Info("http: [%s] %s", info.Request().Method, info.Request().URL.Path)
	return a.guestAuth(ctx, info, a.authToken)
}

// authToken is used to authorize request by token
func (a *Auth) authToken(ctx context.Context, info *transporthttp.Transport) (*meta.DataPermission, error) {
	req := info.Request()
	key := req.Form.Get("access_token")
	if key == "" {
		return nil, ecode.Unauthorized("ErrAccessToken", "access_token is null")
	}
	// NOTE: ??????????????????????????????????????????????????????id
	// TODO: get mid from some code
	// var mid uint64
	resp, err := a.ssoClient.Introspect(ctx, &ssopb.IntrospectReq{
		AccessToken: key,
	})
	if err != nil {
		return nil, err
	}

	return &meta.DataPermission{
		UserId:    resp.UserId,
		RoleId:    resp.RoleId,
		RoleKey:   resp.RoleKey,
		DataScope: resp.DataScope,
	}, nil
}

// authCookie is used to authorize request by cookie
func (a *Auth) authCookie(ctx context.Context, h *transporthttp.Transport) (*meta.DataPermission, error) {
	req := h.Request()
	session, _ := req.Cookie("SESSION")
	if session == nil {
		return nil, ecode.Unauthorized("ErrSession", "session is null")
	}
	// NOTE: ??????????????????????????????????????????????????????id
	// var mid uint64
	// TODO: get mid from some code
	a.log.Info("---------------------")

	// check csrf
	clientCsrf := req.FormValue("csrf")
	if a.conf != nil && !a.conf.DisableCSRF && req.Method == "POST" {
		// NOTE: ???????????????CSRF???????????????CSRF??????????????????????????????csrf
		var csrf string // TODO: get csrf from some code
		if clientCsrf != csrf {
			return nil, ecode.Unauthorized("ErrCsrf", "csrf err")
		}
	}

	return nil, nil
}

// BearerAuth parse bearer token
func (a *Auth) authBearerAuth(ctx context.Context, info *transporthttp.Transport) (*meta.DataPermission, error) {
	auth := info.RequestHeader().Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = info.RequestHeader().Get("access_token")
	}
	resp, err := a.ssoClient.Introspect(ctx, &ssopb.IntrospectReq{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	a.log.Infof("------------------auth:%v", resp)
	return &meta.DataPermission{
		UserId:    resp.UserId,
		RoleId:    resp.RoleId,
		RoleKey:   resp.RoleKey,
		DataScope: resp.DataScope,
	}, nil
}

func (a *Auth) midAuth(ctx context.Context, h *transporthttp.Transport, auth authFunc) (context.Context, error) {
	mid, err := auth(ctx, h)
	if err != nil {
		return nil, err
	}
	return setMid(ctx, mid)
}

func (a *Auth) guestAuth(ctx context.Context, h *transporthttp.Transport, auth authFunc) (context.Context, error) {
	mid, err := auth(ctx, h)
	// no error happened and mid is valid
	if err == nil {
		return setMid(ctx, mid)
	}
	return ctx, nil
}

// set mid into context
// NOTE: This method is not thread safe.
func setMid(ctx context.Context, resp *meta.DataPermission) (context.Context, error) {
	dp, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToClientContext(ctx, "x-md-global-dp", string(dp))
	md, ok := metadata.FromClientContext(ctx)
	if !ok {
		return ctx, ecode.Unauthorized("ErrMD", "md is nil")
	}
	dpp := md.Get("x-md-global-dp")
	fmt.Printf("dp: %v", dpp)
	return ctx, nil
}
