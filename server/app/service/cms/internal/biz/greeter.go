package biz

import (
	"edu/service/cms/internal/conf"
	"edu/service/cms/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type GreeterUsecase struct {
	d   dao.Dao
	log *log.Helper
}

func NewGreeterUsecase(c *conf.App, logger log.Logger, d dao.Dao) *GreeterUsecase {
	return &GreeterUsecase{
		d:   d,
		log: log.NewHelper("usecase/greeter", logger),
	}
}
