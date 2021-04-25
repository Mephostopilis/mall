// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/comet/internal/biz"
	"edu/comet/internal/conf"
	"edu/comet/internal/server"
	"edu/comet/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(uuid.UUID, *conf.Service, *conf.Server, *conf.App, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
