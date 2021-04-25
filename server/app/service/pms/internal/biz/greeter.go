package biz

import (
	"context"
	"edu/service/pms/internal/conf"
	"edu/service/pms/internal/dao"

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

func (uc *GreeterUsecase) GenID(ctx context.Context) (id uint64, err error) {
	return
}
