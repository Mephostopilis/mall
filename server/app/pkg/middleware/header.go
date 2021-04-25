package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transportgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("middleware/options", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Info("transport: %+v", tr)
			}
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
				if h.Request.Method == "OPTIONS" {
					header := h.Response.Header()
					header.Set("Access-Control-Allow-Origin", "*")
					header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
					header.Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
					header.Set("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
					header.Set("Content-Type", "application/json")
				}
			}
			g, ok := transportgrpc.FromServerContext(ctx)
			if ok {
				log.Info("grpc: %s", g.FullMethod)
			}
			return handler(ctx, req)
		}
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("middleware/secure", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Info("transport: %+v", tr)
			}
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
				header := h.Response.Header()
				header.Set("Access-Control-Allow-Origin", "*")
				//c.Header("X-Frame-Options", "DENY")
				header.Set("X-Content-Type-Options", "nosniff")
				header.Set("X-XSS-Protection", "1; mode=block")
				// if c.Request.TLS != nil {
				// 	header.Set("Strict-Transport-Security", "max-age=31536000")
				// }

				// Also consider adding Content-Security-Policy headers
				// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
			}
			g, ok := transportgrpc.FromServerContext(ctx)
			if ok {
				log.Info("grpc: %s", g.FullMethod)
			}
			return handler(ctx, req)
		}
	}
}
