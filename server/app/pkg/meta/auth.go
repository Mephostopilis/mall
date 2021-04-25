package meta

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/metadata"
)

// Config is the identify config model.
type Config struct {
	// csrf switch.
	DisableCSRF bool
}

// Auth is the authorization middleware
type Auth struct {
	conf *Config
	log  *log.Helper
}

// authFunc will return mid and error by given context
type authFunc func(context.Context, *http.Request, metadata.MD) metadata.MD

var _defaultConf = &Config{
	DisableCSRF: false,
}

// New is used to create an authorization middleware
func New(conf *Config, logger log.Logger) *Auth {
	if conf == nil {
		conf = _defaultConf
	}
	auth := &Auth{
		conf: conf,
		log:  log.NewHelper("auth", logger),
	}
	return auth
}

// User is used to mark path as access required.
// If `access_token` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) User(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	if r.Form.Get("access_token") == "" {
		return a.UserWeb(ctx, r, md)
	}
	return a.UserMobile(ctx, r, md)
}

// BearerAuth parse bearer token
func (a *Auth) BearerAuth(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	auth := r.Header.Get("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = r.Form.Get("access_token")
	}
	md.Set("token", token)
	md.Set("ua", r.UserAgent())
	md.Set("remote_addr", r.RemoteAddr)
	md.Set("app_id", "1")
	return md
}

// UserWeb is used to mark path as web access required.
func (a *Auth) UserWeb(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	return a.midAuth(ctx, r, md, a.authCookie)
}

// UserMobile is used to mark path as mobile access required.
func (a *Auth) UserMobile(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	return a.midAuth(ctx, r, md, a.authToken)
}

// Guest is used to mark path as guest policy.
// If `access_token` is exist in request form, it will using mobile access policy.
// Otherwise to web access policy.
func (a *Auth) Guest(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	if r.Form.Get("token") == "" {
		return a.GuestWeb(ctx, r, md)
	}
	return a.GuestMobile(ctx, r, md)
}

// GuestWeb is used to mark path as web guest policy.
func (a *Auth) GuestWeb(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	a.log.Infof("http: [%s] [%s]", r.Method, r.URL.Path)
	return a.guestAuth(ctx, r, md, a.authCookie)
}

// GuestMobile is used to mark path as mobile guest policy.
func (a *Auth) GuestMobile(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	a.log.Infof("http: [%s] %s", r.Method, r.URL.Path)
	return a.guestAuth(ctx, r, md, a.authToken)

}

// authToken is used to authorize request by token
func (a *Auth) authHeader(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	key := r.Header.Get("Authorization")
	if key == "" {
		return md
	}
	md.Set("token", key)
	return md
}

// authToken is used to authorize request by token
func (a *Auth) authToken(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	key := r.Form.Get("access_token")
	if key == "" {
		return md
	}
	md.Set("token", key)
	return md
}

// authCookie is used to authorize request by cookie
func (a *Auth) authCookie(ctx context.Context, r *http.Request, md metadata.MD) metadata.MD {
	session, _ := r.Cookie("SESSION")
	if session == nil {
		return md
	}
	// NOTE: 请求登录鉴权服务接口，拿到对应的用户id
	// var mid uint64
	// TODO: get mid from some code

	// check csrf
	clientCsrf := r.FormValue("csrf")
	if a.conf != nil && !a.conf.DisableCSRF && r.Method == "POST" {
		// NOTE: 如果开启了CSRF认证，请从CSRF服务获取该用户关联的csrf
		var csrf string // TODO: get csrf from some code
		if clientCsrf != csrf {
			return md
		}
	}

	md.Set("token", session.Value)
	return md
}

func (a *Auth) midAuth(ctx context.Context, r *http.Request, md metadata.MD, auth authFunc) metadata.MD {
	return auth(ctx, r, md)
}

func (a *Auth) guestAuth(ctx context.Context, r *http.Request, md metadata.MD, auth authFunc) metadata.MD {
	return auth(ctx, r, md)
}
