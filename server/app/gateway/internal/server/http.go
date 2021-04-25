package server

import (
	"context"
	xhttp "net/http"

	cmspb "edu/api/cms"
	ssopb "edu/api/sso"
	syspb "edu/api/sys/v1"
	tikupb "edu/api/tiku"
	"edu/gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type httpServer struct {
	mux *gwruntime.ServeMux
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *httpServer) ServeHTTP(res xhttp.ResponseWriter, req *xhttp.Request) {
	// ctx, cancel := context.WithTimeout(req.Context(), s.timeout)
	// defer cancel()
	// ctx = transport.NewContext(ctx, transport.Transport{Kind: transport.KindHTTP})
	// ctx = NewServerContext(ctx, ServerInfo{Request: req, Response: res})
	// s.router.ServeHTTP(res, req.WithContext(ctx))
	s.mux.ServeHTTP(res, req)
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger) *http.Server {
	// middleware
	// 全局中間件
	// csrf := bm.CSRF([]string{"127.0.0.1", "192.168.21.95"}, []string{"/api"})
	// engine.Use(csrf)
	// httpSrv.Use(middleware.RequestId())
	// httpTransport.Use(httpServer, middleware.Options())
	// cors := bm.CORS([]string{"127.0.0.1:9527", "127.0.0.1:9527/admin", "127.0.0.1:9527/api/v1"})
	// engine.Use(cors)

	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Address != "" {
		opts = append(opts, http.Address(c.Http.Address))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.Logger(logger))
	httpSrv := http.NewServer(opts...)

	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
			status.Client(),
			tracing.Client(),
			logging.Client(logging.WithLogger(logger)),
		),
	)

	ctx := context.Background()
	mux := newGateway()

	ssoc, err := ssopb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	ssopb.RegisterApiHandlerClient(ctx, mux, ssoc)
	sysc, err := syspb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	syspb.RegisterApiHandlerClient(ctx, mux, sysc)
	tikuc, err := tikupb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	tikupb.RegisterApiHandlerClient(ctx, mux, tikuc)
	cmsc, err := cmspb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	cmspb.RegisterApiHandlerClient(ctx, mux, cmsc)

	httpSrv.HandlePrefix("/api/v1", mux)
	httpSrv.HandlePrefix("/api/swagger/", xhttp.StripPrefix("/api/swagger/", xhttp.FileServer(xhttp.Dir("../../swagger-ui/dist"))))
	httpSrv.HandlePrefix("/api/docs/", xhttp.StripPrefix("/api/docs/", xhttp.FileServer(xhttp.Dir("../../../api"))))

	httpSrv.HandleFunc("/api/start", func(w xhttp.ResponseWriter, r *xhttp.Request) {
		w.Write([]byte("hello kratos!"))
	})
	return httpSrv
}
