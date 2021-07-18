package auth

import (
	"context"

	ssopb "edu/api/sso/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Config is the identify config model.
type Config struct {
	// csrf switch.
	DisableCSRF bool
}

// Auth is the authorization middleware
type Auth struct {
	ssoClient ssopb.SsoClient
	conf      *Config
	log       *log.Helper
}

// authFunc will return mid and error by given context
type authFunc func(context.Context, *transporthttp.Transport) (uint64, error)

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
				if h.RequestHeader().Get("access_token") == "" {
					a.UserWeb(ctx, h)
					return
				}
				a.UserMobile(ctx, h)
			}
		}
		return handler(ctx, req)
	}
}

// BearerAuth parse bearer token
// func (a *Auth) BearerAuth(ctx *bm.Context) (string, bool) {
// 	auth := ctx.Request.Header.Get("Authorization")
// 	prefix := "Bearer "
// 	token := ""

// 	if auth != "" && strings.HasPrefix(auth, prefix) {
// 		token = auth[len(prefix):]
// 	} else {
// 		token = ctx.Params.ByName("access_token")
// 	}
// 	return token, token != ""
// }

// UserWeb is used to mark path as web access required.
func (a *Auth) UserWeb(ctx context.Context, info *transporthttp.Transport) (context.Context, error) {
	return a.midAuth(ctx, info, a.authCookie)
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
func (a *Auth) authToken(ctx context.Context, info *transporthttp.Transport) (uint64, error) {
	req := info.Request()
	key := req.Form.Get("access_token")
	if key == "" {
		return 0, errors.Unauthorized("ErrAccessToken", "access_token is null")
	}
	// NOTE: 请求登录鉴权服务接口，拿到对应的用户id
	// TODO: get mid from some code
	var mid uint64
	resp, err := a.ssoClient.Introspect(ctx, &ssopb.IntrospectReq{
		AccessToken: key,
	})
	if err != nil {
		return 0, err
	}
	mid = uint64(resp.UserId)
	return mid, nil
}

// authCookie is used to authorize request by cookie
func (a *Auth) authCookie(ctx context.Context, h *transporthttp.Transport) (uint64, error) {
	req := h.Request()
	session, _ := req.Cookie("SESSION")
	if session == nil {
		return 0, errors.Unauthorized("ErrSession", "session is null")
	}
	// NOTE: 请求登录鉴权服务接口，拿到对应的用户id
	var mid uint64
	// TODO: get mid from some code

	// check csrf
	clientCsrf := req.FormValue("csrf")
	if a.conf != nil && !a.conf.DisableCSRF && req.Method == "POST" {
		// NOTE: 如果开启了CSRF认证，请从CSRF服务获取该用户关联的csrf
		var csrf string // TODO: get csrf from some code
		if clientCsrf != csrf {
			return 0, errors.Unauthorized("ErrCsrf", "csrf err")
		}
	}

	return mid, nil
}

func (a *Auth) midAuth(ctx context.Context, h *transporthttp.Transport, auth authFunc) (context.Context, error) {
	mid, err := auth(ctx, h)
	if err != nil {
		return nil, errors.Unauthorized("mid", "auth not found")
	}
	return setMid(ctx, mid)
}

func (a *Auth) guestAuth(ctx context.Context, h *transporthttp.Transport, auth authFunc) (context.Context, error) {
	mid, err := auth(ctx, h)
	// no error happened and mid is valid
	if err == nil && mid > 0 {
		return setMid(ctx, mid)
	}
	return ctx, nil
}

// set mid into context
// NOTE: This method is not thread safe.
func setMid(ctx context.Context, mid uint64) (context.Context, error) {
	ctx = context.WithValue(ctx, "mid", mid)
	return ctx, nil
}
