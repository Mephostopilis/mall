package dao

import (
	"context"

	"edu/service/sms/internal/conf"
	"edu/service/sms/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(New)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)

	GetSmsCouponPage(ctx context.Context, e *model.SmsCoupon, pageSize int, pageIndex int) (docs []model.SmsCoupon, total int64, err error)
	GetSmsCoupon(ctx context.Context, e *model.SmsCoupon) (model.SmsCoupon, error)
	CreateSmsCoupon(ctx context.Context, e *model.SmsCoupon) (model.SmsCoupon, error)
	UpdateSmsCoupon(ctx context.Context, e *model.SmsCoupon, id uint64) (update model.SmsCoupon, err error)
	BatchDeleteSmsCoupon(ctx context.Context, e *model.SmsCoupon, id []uint64) (Result bool, err error)

	GetSmsCouponHistoryPage(ctx context.Context, e *model.SmsCouponHistory, pageSize int, pageIndex int) (docs []model.SmsCouponHistory, total int64, err error)
	GetSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory) (model.SmsCouponHistory, error)
	CreateSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory) (model.SmsCouponHistory, error)
	UpdateSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory, id uint64) (update model.SmsCouponHistory, err error)
	BatchDeleteSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory, id []uint64) (Result bool, err error)
}

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
	log := log.NewHelper("dao", logger)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Error("get account fail")
		return
	}
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
