package biz

import (
	"context"
	"edu/service/oms/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/sony/sonyflake"
)

type GreeterUsecase struct {
	sf  *sonyflake.Sonyflake
	log *log.Helper
}

func NewGreeterUsecase(c *conf.App, logger log.Logger) *GreeterUsecase {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return uint16(c.Uuid.MachineId), nil
		},
	})
	return &GreeterUsecase{
		sf:  sf,
		log: log.NewHelper("usecase/greeter", logger),
	}
}

func (uc *GreeterUsecase) GenID(ctx context.Context) (id uint64, err error) {
	id, err = uc.sf.NextID()
	if err != nil {
		return
	}
	return
}
