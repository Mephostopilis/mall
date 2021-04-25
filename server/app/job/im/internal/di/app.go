package di

import (
	"edu/job/im/internal/conf"
	"edu/job/im/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

func newApp(srv *conf.Service, logger log.Logger, gs *server.AmqpConsumer) *kratos.App {
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
