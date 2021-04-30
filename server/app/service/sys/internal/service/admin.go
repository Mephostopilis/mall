package service

import (
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/biz"
	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.AdminUsecase
}

func NewAdminService(genConf *conf.App, logger log.Logger, adao dao.Dao, uc *biz.AdminUsecase) *AdminService {
	log := log.NewHelper("service/admin", logger)
	return &AdminService{
		log: log,
		uc:  uc,
	}
}
