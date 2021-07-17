package biz

import (
	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminUsecase struct {
	d      dao.Dao
	logger log.Logger
	log    *log.Helper
	cfg    *conf.App
}

func NewAdminUsecase(c *conf.App, logger log.Logger, repo dao.Dao) (*AdminUsecase, error) {
	log := log.NewHelper(log.With(logger, "module", "usecase/admin"))
	return &AdminUsecase{
		d:      repo,
		logger: logger,
		log:    log,
		cfg:    c,
	}, nil
}
