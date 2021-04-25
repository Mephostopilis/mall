package di

import (
	"edu/comet/internal/conf"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
)

func newApp(id uuid.UUID, srv *conf.Service, logger log.Logger, gs *grpc.Server, r *registry.Registry) *kratos.App {
	app := kratos.New(
		kratos.ID(id.String()),
		kratos.Name(srv.Name),
		kratos.Version(srv.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		// kratos.Registrar(r),
	)
	return app
}
