package gcasbin

import (
	"context"
	"net/http"

	"edu/admin/internal/middleware/jwtauth"

	"github.com/casbin/casbin/v2"
	lcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	transportgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

type CasbinMiddleware struct {
	enforcer *lcasbin.Enforcer
	log      *log.Helper
	subFn    SubjectFn
}

// SubjectFn is used to look up current subject in runtime.
// If it can not find anything, just return an empty string.
type SubjectFn func(w http.ResponseWriter, r *http.Request) string

// Logic is the logical operation (AND/OR) used in permission checks
// in case multiple permissions or roles are specified.
type Logic int

const (
	AND Logic = iota
	OR
)

var (
	SubFnNilErr = errors.InvalidArgument("sunfn", "subFn is nil")
)

// NewCasbinMiddleware returns a new CasbinMiddleware using Casbin's Enforcer internally.
// modelFile is the file path to Casbin model file e.g. path/to/rbac_model.conf.
// policyAdapter can be a file or a DB adapter.
// File: path/to/basic_policy.csv
// MySQL DB: mysqladapter.NewDBAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
// subFn is a function that looks up the current subject in runtime and returns an empty string if nothing found.
func newCasbinMiddleware(logger log.Logger, modelFile string, policyAdapter interface{}, subFn SubjectFn) (*CasbinMiddleware, error) {
	log := log.NewHelper("middleware/gcasbin", logger)
	if subFn == nil {
		return nil, SubFnNilErr
	}
	m, err := model.NewModelFromString(modelFile)
	if err != nil {
		log.Errorf("err = %v", err)
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, policyAdapter)
	if err != nil {
		log.Errorf("err = %v", err)
		return nil, err
	}
	return &CasbinMiddleware{
		enforcer: e,
		subFn:    subFn,
		log:      log,
	}, nil
}

// Option is used to change some default behaviors.
type Option interface {
	apply(*options)
}

type options struct {
	logic Logic
}

type logicOption Logic

func (lo logicOption) apply(opts *options) {
	opts.logic = Logic(lo)
}

// WithLogic sets the logical operator used in permission or role checks.
func WithLogic(logic Logic) Option {
	return logicOption(logic)
}

// AuthCheckRole 检查角色
func (am *CasbinMiddleware) AuthCheckRole() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				rolekey := jwtauth.GetRoleName(ctx)

				//检查权限
				res, err := am.enforcer.Enforce(rolekey, h.Request.URL.Path, h.Request.Method)
				if err != nil {
					return nil, err
				}
				if res {
					return handler(ctx, req)
				}
				return nil, errors.Unauthorized("casbin", "[%s] not unauthorized", rolekey)
			}
			g, ok := transportgrpc.FromServerContext(ctx)
			if ok {
				am.log.Infof("grpc: %s", g.FullMethod)
			}
			return handler(ctx, req)
		}
	}
}
