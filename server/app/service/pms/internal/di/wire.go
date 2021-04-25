// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/pms/internal/biz"
	"edu/service/pms/internal/conf"
	"edu/service/pms/internal/dao"
	"edu/service/pms/internal/server"
	"edu/service/pms/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.Data, *conf.App, log.Logger, registry.Registrar) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
