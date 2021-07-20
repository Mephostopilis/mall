package server

import (
	"context"
	xhttp "net/http"

	"edu/admin/internal/conf"
	cmspb "edu/api/cms/v1"
	memberpb "edu/api/member"
	omspb "edu/api/oms"
	pmspb "edu/api/pms"
	smspb "edu/api/sms"
	ssopb "edu/api/sso/v1"
	syspb "edu/api/sys/v1"
	tikupb "edu/api/tiku"
	"edu/pkg/middleware/auth"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
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
func NewHTTPServer(c *conf.Server, authc *conf.Auth, logger log.Logger, r *registry.Registry) *http.Server {
	ssocc, err := ssopb.NewSso(context.TODO(),
		grpc.WithMiddleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Client(),
				logging.Client(logger),
			),
		),
		grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	au := auth.New(&auth.Config{DisableCSRF: false, WhiteList: authc.WhiteList}, logger, ssocc)
	m := grpc.WithMiddleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Client(),
			logging.Client(logger),
			validate.Validator(),
			au.User,
			metadata.Client(),
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

	// member
	memberc, err := memberpb.NewAdmin(ctx, m, grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}
	memberpb.RegisterAdminHandlerClient(ctx, gr, memberc)

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
	httpSrv.HandlePrefix("/admin/v1", cors(gr))
	httpSrv.HandlePrefix("/admin/login", cors(gr))
	httpSrv.HandlePrefix("/admin/refresh_token", cors(gr))

	h := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q/", h)
	httpSrv.HandlePrefix("/form-generator/", xhttp.StripPrefix("/form-generator/", xhttp.FileServer(xhttp.Dir("../../static/form-generator"))))

	httpSrv.HandleFunc("/admin/start", func(w xhttp.ResponseWriter, r *xhttp.Request) {
		w.Write([]byte("hello kratos!"))
	})
	return httpSrv
}
