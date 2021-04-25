// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/admin/internal/conf"
	"edu/admin/internal/server"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
