package biz

import (
	"context"
	"edu/pkg/ecode"
	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type JobUsecase struct {
	admin dao.Dao
	log   *log.Helper
	cfg   *conf.App
}

func NewJobUsecase(logger log.Logger, d dao.Dao) (uc *JobUsecase, err error) {
	log := log.NewHelper("usecase/job", logger)
	uc = &JobUsecase{
		log: log,
	}
	return
}

// AddJob 添加任务 AddJob(invokeTarget string, jobId int, jobName string, cronExpression string)
func (s *JobUsecase) AddJob(ctx context.Context, exp string) (EntryId int, err error) {
	return 0, ecode.AdminNotSupportedJob
}

func (s *JobUsecase) addHttpJob(ctx context.Context, job HttpJob) (EntryId int, err error) {
	// ch := cron.NewChain(cron.DelayIfStillRunning()).Then(&job)
	// s.sche.AddJob(ch)

	// id, err1 := s.sche.AddJob(job.CronExpression, job)
	// if err1 != nil {
	// 	log.Error("JobCore AddJob error", err)
	// 	err = err1
	// 	return
	// }
	// EntryId = int(id)
	// return
	return
}

func (s *JobUsecase) addExecJob(ctx context.Context, job ExecJob) (EntryId int, err error) {
	// id, err := s.sche.AddJob(job.CronExpression, job)
	// if err != nil {
	// 	log.Error("JobCore AddJob error: %v", err)
	// 	return 0, err
	// }
	// EntryId := int(id)
	return EntryId, nil
}

// RemoveJob 移除任务
func (s *JobUsecase) RemoveJob(ctx context.Context, entryID int) chan bool {
	ch := make(chan bool)
	go func() {
		// s.sche.Remove(cron.EntryID(entryID))
		s.log.Infof("JobCore Remove success ,info entryID :", entryID)
		ch <- true
	}()
	return ch
}
