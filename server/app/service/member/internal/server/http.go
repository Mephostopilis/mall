package server

import (
	xhttp "net/http"

	pb "edu/api/member"
	"edu/pkg/middleware/handleerr"
	"edu/service/member/internal/conf"
	"edu/service/member/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, a *service.AdminService, s *service.ApiService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				handleerr.Server(),
				tracing.Server(),
				logging.Server(logger),
			),
		),
	}
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

	pb.RegisterAdminHTTPServer(httpSrv, a)
	pb.RegisterApiHTTPServer(httpSrv, s)

	httpSrv.HandleFunc("/kratos/start", func(w xhttp.ResponseWriter, r *xhttp.Request) {
		w.Write([]byte("hello kratos!"))
	})

	return httpSrv
}
