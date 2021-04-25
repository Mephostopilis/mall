package server

import (
	"edu/service/riot/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type CronSrv struct {
	cr  *cron.Cron
	log *log.Helper
	uc  *biz.Engine
}

// NewCronServer 新cron
func NewCronServer(logger log.Logger, uc *biz.Engine) (d *CronSrv, err error) {
	log := log.NewHelper("server/cron", logger)
	cr := cron.New()
	s := &CronSrv{
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

func (s *CronSrv) Endpoint() (ep string, err error) {
	return
}

func (s *CronSrv) Start() error {
	s.cr.Start()
	s.log.Infof("start corn")
	return nil
}

func (s *CronSrv) Stop() error {
	s.cr.Stop()
	s.log.Infof("stop corn")
	return nil
}

func (s *CronSrv) syncPort() {
	// 增加
	// s.uc.StepcaHealth(context.Background())
	// for i := 0; i < len(jobList); i++ {
	// 	if jobList[i].JobType == 1 {
	// 		j := &model.HttpJob{}
	// 		j.InvokeTarget = jobList[i].InvokeTarget
	// 		j.CronExpression = jobList[i].CronExpression
	// 		j.JobId = jobList[i].JobId
	// 		j.Name = jobList[i].JobName

	// 		sysJob.EntryId, err = s.AddJob(j)
	// 	} else if jobList[i].JobType == 2 {
	// 		j := &model.ExecJob{}
	// 		j.InvokeTarget = jobList[i].InvokeTarget
	// 		j.CronExpression = jobList[i].CronExpression
	// 		j.JobId = jobList[i].JobId
	// 		j.Name = jobList[i].JobName
	// 		sysJob.EntryId, err = s.AddJob(j)
	// 	}
	// 	_, err = d.UpdateSysJob(&sysJob, jobList[i].JobId)
	// }
	return
}
