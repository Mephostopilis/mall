// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/uuid/internal/biz"
	"edu/service/uuid/internal/conf"
	"edu/service/uuid/internal/server"
	"edu/service/uuid/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.App, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
