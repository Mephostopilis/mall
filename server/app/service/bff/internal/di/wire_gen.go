// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"edu/service/cms/internal/biz"
	"edu/service/cms/internal/conf"
	"edu/service/cms/internal/dao"
	"edu/service/cms/internal/server"
	"edu/service/cms/internal/service"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func InitApp(confService *conf.Service, confServer *conf.Server, data *conf.Data, app *conf.App, logger log.Logger, registryRegistry *registry.Registry) (*kratos.App, error) {
	daoDao, err := dao.New(data, logger)
	if err != nil {
		return nil, err
	}
	greeterUsecase, err := biz.NewGreeterUsecase(logger, daoDao, registryRegistry)
	if err != nil {
		return nil, err
	}
	apiService := service.NewCmsService(greeterUsecase)
	adminService := service.NewAdminService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, logger, apiService, adminService)
	kratosApp := newApp(confService, logger, grpcServer, registryRegistry)
	return kratosApp, nil
}