// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/tiku/internal/biz"
	"edu/service/tiku/internal/conf"
	"edu/service/tiku/internal/dao"
	"edu/service/tiku/internal/server"
	"edu/service/tiku/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.Data, *conf.App, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
