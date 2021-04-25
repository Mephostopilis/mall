package biz

import (
	"edu/pkg/jwtauth"
	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminUsecase struct {
	d      dao.Dao
	logger log.Logger
	log    *log.Helper
	mw     *jwtauth.GinJWTMiddleware
	cfg    *conf.App
}

func NewAdminUsecase(c *conf.App, logger log.Logger, repo dao.Dao) (*AdminUsecase, error) {
	log := log.NewHelper("usecase/admin", logger)
	mw, err := jwtauth.New(
		&jwtauth.Jwt{Secret: c.Jwt.Secret, Timeout: c.Jwt.Timeout.AsDuration()},
		logger,
		jwtauth.DataPermissionToClaimsFunc,
		jwtauth.ClaimsToDataPermissionFunc)
	if err != nil {
		return nil, err
	}
	return &AdminUsecase{
		d:      repo,
		mw:     mw,
		logger: logger,
		log:    log,
		cfg:    c,
	}, nil
}
