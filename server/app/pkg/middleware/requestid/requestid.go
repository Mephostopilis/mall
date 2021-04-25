package requestid

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
)

type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with middleware logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// RequestId 唯一的请求id
func RequestId(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/requestid", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Infof("transport: %+v", tr)
			}
			h, ok := http.FromServerContext(ctx)
			if ok {
				log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
				// Check for incoming header, use it if exists
				requestId := h.Request.Header.Get("X-Request-Id")

				// Create request id with UUID4
				if requestId == "" {
					u4, _ := uuid.NewUUID()
					requestId = u4.String()
				}

				// Expose it for use in the application
				ctx = context.WithValue(ctx, "X-Request-Id", requestId)

				// Set X-Request-Id header
				h.Response.Header().Set("X-Request-Id", requestId)
			}
			g, ok := grpc.FromServerContext(ctx)
			if ok {
				log.Info("grpc: %s", g.FullMethod)
			}
			return handler(ctx, req)
		}
	}
}
