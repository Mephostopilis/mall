package server

import (
	"edu/service/tiku/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

// 检测所有的服务
type CronSrv interface {
	Endpoint() (string, error)
	Start() error
	Stop() error
}

type cronSrv struct {
	cr  *cron.Cron
	log *log.Helper
	uc  *biz.TikuUsecase
}

// NewCron 新cron
func NewCron(logger log.Logger, uc *biz.TikuUsecase) (d CronSrv, err error) {
	log := log.NewHelper("server/cron", logger)
	cr := cron.New()
	s := &cronSrv{
		cr:  cr,
		log: log,
		uc:  uc,
	}
	spec := "0/1 * * * *"
	cr.AddFunc(spec, func() {
		s.syncPort()
		// s.verifyIdentity()
	})
	d = s
	return
}

func (s *cronSrv) Endpoint() (ep string, err error) {
	return "", nil
}

func (s *cronSrv) Start() error {
	s.cr.Start()
	s.log.Infof("start corn")
	return nil
}

func (s *cronSrv) Stop() error {
	s.cr.Stop()
	s.log.Infof("stop corn")
	return nil
}

func (s *cronSrv) syncPort() {
	return
}
