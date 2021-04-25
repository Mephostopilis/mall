// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"edu/service/member/internal/biz"
	"edu/service/member/internal/conf"
	"edu/service/member/internal/dao"
	"edu/service/member/internal/server"
	"edu/service/member/internal/service"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func InitApp(*conf.Service, *conf.Server, *conf.App, *conf.Data, log.Logger, *registry.Registry) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
