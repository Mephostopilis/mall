package gcasbin

import (
	"context"

	syspb "edu/api/sys/v1"
	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

type Option func(*options)

type options struct {
	logger    log.Logger
	sysc      syspb.SysClient
	whiteList map[string]bool
}

// WithLogger with middleware logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithSysc with middleware logger.
func WithSysc(sysc syspb.SysClient) Option {
	return func(o *options) {
		o.sysc = sysc
	}
}

// WithSysc with middleware logger.
func WithWhiteList(list []string) Option {
	return func(o *options) {
		for i := 0; i < len(list); i++ {
			url := list[i]
			o.whiteList[url] = true
		}
	}
}

type CasbinMiddleware struct {
	sysc      syspb.SysClient
	log       *log.Helper
	WhiteList map[string]bool
}

// NewCasbinMiddleware returns a new CasbinMiddleware using Casbin's Enforcer internally.
// modelFile is the file path to Casbin model file e.g. path/to/rbac_model.conf.
// policyAdapter can be a file or a DB adapter.
// File: path/to/basic_policy.csv
// MySQL DB: mysqladapter.NewDBAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
// subFn is a function that looks up the current subject in runtime and returns an empty string if nothing found.
func NewCasbinMiddleware(opts ...Option) (*CasbinMiddleware, error) {
	_options := options{
		logger:    log.DefaultLogger,
		whiteList: map[string]bool{},
	}
	for _, o := range opts {
		o(&_options)
	}
	log := log.NewHelper(log.With(_options.logger, "module", "middleware/gcasbin"))
	return &CasbinMiddleware{
		sysc:      _options.sysc,
		WhiteList: _options.whiteList,
		log:       log,
	}, nil
}

// AuthCheckRole 检查角色
func (am *CasbinMiddleware) AuthCheckRole() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			info, ok := transport.FromServerContext(ctx)
			if ok {
				if info.Kind() == transport.KindHTTP {
					h := info.(*transporthttp.Transport)

					if am.WhiteList[h.Request().URL.Path] {
						ret, err1 := am.sysc.CheckResource(ctx, &syspb.CheckResourceRequest{
							Path:   h.Request().URL.Path,
							Method: h.Request().Method,
						})
						if err1 != nil {
							err = err1
							return
						}
						if !ret.Ok {
							err = ecode.Unauthorized("ErrCheckRole", "fail")
							return
						}
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
