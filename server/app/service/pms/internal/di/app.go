package di

import (
	"edu/service/pms/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func newApp(srv *conf.Service, logger log.Logger, gs *grpc.Server, r registry.Registrar) *kratos.App {
	app := kratos.New(
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
