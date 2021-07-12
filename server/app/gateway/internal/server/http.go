package server

import (
	"context"
	xhttp "net/http"

	cmspb "edu/api/cms"
	memberpb "edu/api/member"
	pmspb "edu/api/pms"
	ssopb "edu/api/sso"
	syspb "edu/api/sys/v1"
	tikupb "edu/api/tiku"
	"edu/gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
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

	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
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
	// tiku
	tikuc, err := tikupb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	tikupb.RegisterApiHandlerClient(ctx, mux, tikuc)
	// pms
	pmsc, err := pmspb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	pmspb.RegisterApiHandlerClient(ctx, mux, pmsc)
	// member
	memberc, err := memberpb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	memberpb.RegisterApiHandlerClient(ctx, mux, memberc)
	// cms
	cmsc, err := cmspb.NewApi(ctx, m)
	if err != nil {
		panic(err)
	}
	cmspb.RegisterApiHandlerClient(ctx, mux, cmsc)

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
	httpSrv.HandlePrefix("/api/v1", mux)

	h := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", h)

	httpSrv.HandleFunc("/api/start", func(w xhttp.ResponseWriter, r *xhttp.Request) {
		w.Write([]byte("hello kratos!"))
	})
	return httpSrv
}
