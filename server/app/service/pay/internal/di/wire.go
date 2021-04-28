// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/pay/internal/biz"
	"edu/service/pay/internal/conf"
	"edu/service/pay/internal/dao"
	"edu/service/pay/internal/server"
	"edu/service/pay/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.Data, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
