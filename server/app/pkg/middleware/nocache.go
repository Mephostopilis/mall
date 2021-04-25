package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transportgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

var (
	// Unix epoch time.
	// nolint:gochecknoglobals
	epoch = time.Unix(0, 0).Format(time.RFC1123)

	// Taken from https://github.com/mytrile/nocache
	// nolint:gochecknoglobals
	noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, no-store, no-transform, must-revalidate, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
		"Last-Modified":   time.Now().UTC().Format(http.TimeFormat),
	}

	// ETag headers array.
	// nolint:gochecknoglobals
	etagHeaders = [6]string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	}
)

// NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent
// a router (or subrouter) from being cached by an upstream proxy and/or client.
//
// As per http://wiki.nginx.org/HttpProxyModule - NoCache sets:
//      Expires: Thu, 01 Jan 1970 00:00:00 UTC
//      Cache-Control: no-cache, no-store, no-transform, must-revalidate, private, max-age=0
//      Pragma: no-cache (for HTTP/1.0 proxies/clients)
//      X-Accel-Expires: 0
func NoCache(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("middleware/nocache", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Info("transport: %+v", tr)
			}
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
				for _, v := range etagHeaders {
					if h.Request.Header.Get(v) != "" {
						h.Request.Header.Del(v)
					}
				}

				// Set our NoCache headers
				for k, v := range noCacheHeaders {
					h.Response.Header().Set(k, v)
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
