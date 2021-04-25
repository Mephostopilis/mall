package server

import (
	"context"
	xhttp "net/http"

	"edu/admin/internal/conf"
	cmspb "edu/api/cms"
	omspb "edu/api/oms"
	pmspb "edu/api/pms"
	smspb "edu/api/sms"
	ssopb "edu/api/sso"
	syspb "edu/api/sys/v1"
	tikupb "edu/api/tiku"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
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
func NewHTTPServer(c *conf.Server, logger log.Logger, r *registry.Registry) *http.Server {
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

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	ctx := context.Background()
	gr := newGateway(logger)

	// sso
	ssoc, err := ssopb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	ssopb.RegisterAdminHandlerClient(ctx, gr, ssoc)
	// sys
	sysc, err := syspb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	syspb.RegisterAdminHandlerClient(ctx, gr, sysc)
	// tiku
	tikuc, err := tikupb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	tikupb.RegisterAdminHandlerClient(ctx, gr, tikuc)

	// cms
	cmsc, err := cmspb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	cmspb.RegisterAdminHandlerClient(ctx, gr, cmsc)

	// pms
	pmsc, err := pmspb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	pmspb.RegisterAdminHandlerClient(ctx, gr, pmsc)

	// oms
	omsc, err := omspb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	omspb.RegisterAdminHandlerClient(ctx, gr, omsc)

	// sms
	smsc, err := smspb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	smspb.RegisterAdminHandlerClient(ctx, gr, smsc)

	httpSrv.HandlePrefix("/admin/v1", cors(gr))
	httpSrv.HandlePrefix("/admin/login", cors(gr))
	httpSrv.HandlePrefix("/admin/refresh_token", cors(gr))
	httpSrv.HandlePrefix("/admin/swagger/", xhttp.StripPrefix("/admin/swagger/", xhttp.FileServer(xhttp.Dir("../../swagger-ui/dist"))))
	httpSrv.HandlePrefix("/admin/docs/", xhttp.StripPrefix("/admin/docs/", xhttp.FileServer(xhttp.Dir("../../../api"))))
	httpSrv.HandlePrefix("/form-generator/", xhttp.StripPrefix("/form-generator/", xhttp.FileServer(xhttp.Dir("../../static/form-generator"))))

	httpSrv.HandleFunc("/admin/doc/*", swaggerServer(logger, "../../doc/"))
	httpSrv.HandleFunc("/admin/start", func(w xhttp.ResponseWriter, r *xhttp.Request) {
		w.Write([]byte("hello kratos!"))
	})
	return httpSrv
}
