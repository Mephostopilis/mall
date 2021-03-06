// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"edu/pms/internal/biz"
	"edu/pms/internal/conf"
	"edu/pms/internal/dao"
	"edu/pms/internal/server"
	"edu/pms/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

// Injectors from wire.go:

// initApp init kratos application.
func InitApp(confService *conf.Service, confServer *conf.Server, data *conf.Data, app *conf.App, logger log.Logger, registrar registry.Registrar) (*kratos.App, error) {
	daoDao, err := dao.New(data, logger)
	if err != nil {
		return nil, err
	}
	greeterUsecase := biz.NewGreeterUsecase(app, logger, daoDao)
	apiService := service.NewApiService(greeterUsecase)
	adminService := service.NewAdminService(logger, greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, logger, apiService, adminService)
	kratosApp := newApp(confService, logger, grpcServer, registrar)
	return kratosApp, nil
}
