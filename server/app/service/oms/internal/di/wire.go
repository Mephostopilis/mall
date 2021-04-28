// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/oms/internal/biz"
	"edu/service/oms/internal/conf"
	"edu/service/oms/internal/dao"
	"edu/service/oms/internal/server"
	"edu/service/oms/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.Data, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
