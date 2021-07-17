package di

import (
	"edu/service/sys/internal/conf"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func newApp(srv *conf.Service, logger log.Logger, gs *grpc.Server, hs *http.Server, r *registry.Registry) *kratos.App {
	return kratos.New(
		kratos.Name(srv.Name),
		kratos.Version(srv.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	)
}
