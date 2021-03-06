// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"edu/service/uuid/internal/biz"
	"edu/service/uuid/internal/conf"
	"edu/service/uuid/internal/server"
	"edu/service/uuid/internal/service"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func InitApp(confService *conf.Service, confServer *conf.Server, app *conf.App, logger log.Logger, registryRegistry *registry.Registry) (*kratos.App, error) {
	greeterUsecase := biz.NewGreeterUsecase(app, logger)
	uuidService := service.NewUUIDService(logger, greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, logger, uuidService)
	kratosApp := newApp(confService, logger, grpcServer, registryRegistry)
	return kratosApp, nil
}
