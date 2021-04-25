package dao

import (
	"context"

	"edu/service/sys/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(New)

// dao dao.
type dao struct {
	orm *gorm.DB
	log *log.Helper
}

// New new a dao and return.
func New(c *conf.Data, logger log.Logger) (d Dao, err error) {
	return newDao(c, logger)
}

func newDao(c *conf.Data, logger log.Logger) (d *dao, err error) {
	log := log.NewHelper("dao/admin", logger)
	db, err := gorm.Open(mysql.Open(c.Admin.Source), &gorm.Config{})
	if err != nil {
		log.Error("get account fail")
		return
	}
	// db.DB().SetMaxIdleConns(cfg.Idle)
	// db.DB().SetMaxOpenConns(cfg.Active)

	d = &dao{
		orm: db,
		log: log,
	}
	return
}

// Close close the resource.
func (d *dao) Close() {
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
