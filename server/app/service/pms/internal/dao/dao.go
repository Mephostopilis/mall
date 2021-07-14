package dao

import (
	"context"

	"edu/service/pms/internal/conf"
	"edu/service/pms/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(New)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)

	GetPmsAlbumPage(ctx context.Context, e *model.PmsAlbum, pageSize int, pageIndex int) (docs []model.PmsAlbum, total int64, err error)
	GetPmsAlbum(ctx context.Context, e *model.PmsAlbum) (model.PmsAlbum, error)
	CreatePmsAlbum(ctx context.Context, e *model.PmsAlbum) (model.PmsAlbum, error)
	UpdatePmsAlbum(ctx context.Context, e *model.PmsAlbum, id uint64) (update model.PmsAlbum, err error)
	BatchDeletePmsAlbum(ctx context.Context, e *model.PmsAlbum, id []uint64) (Result bool, err error)

	GetPmsBrandPage(ctx context.Context, e *model.PmsBrand, pageSize int, pageIndex int) (docs []model.PmsBrand, total int64, err error)
	GetPmsBrand(ctx context.Context, e *model.PmsBrand) (model.PmsBrand, error)
	CreatePmsBrand(ctx context.Context, e *model.PmsBrand) (model.PmsBrand, error)
	UpdatePmsBrand(ctx context.Context, e *model.PmsBrand, id uint64) (update model.PmsBrand, err error)
	BatchDeletePmsBrand(ctx context.Context, e *model.PmsBrand, id []uint64) (Result bool, err error)

	GetPmsProductCategoryList(ctx context.Context, e *model.PmsProductCategory) (docs []model.PmsProductCategory, err error)
	GetPmsProductCategory(ctx context.Context, e *model.PmsProductCategory) (model.PmsProductCategory, error)
	CreatePmsProductCategory(ctx context.Context, e *model.PmsProductCategory) (model.PmsProductCategory, error)
	UpdatePmsProductCategory(ctx context.Context, e *model.PmsProductCategory, id uint64) (update model.PmsProductCategory, err error)
	BatchDeletePmsProductCategory(ctx context.Context, e *model.PmsProductCategory, id []uint64) (Result bool, err error)

	GetPmsProductPage(ctx context.Context, e *model.PmsProduct, pageSize int, pageIndex int) (docs []model.PmsProduct, total int64, err error)
	GetPmsProduct(ctx context.Context, e *model.PmsProduct) (model.PmsProduct, error)
	CreatePmsProduct(ctx context.Context, e *model.PmsProduct) (model.PmsProduct, error)
	UpdatePmsProduct(ctx context.Context, e *model.PmsProduct, id uint64) (update model.PmsProduct, err error)
	BatchDeletePmsProduct(ctx context.Context, e *model.PmsProduct, id []uint64) (Result bool, err error)

	GetPmsProductAttributePage(ctx context.Context, e *model.PmsProductAttribute, pageSize int, pageIndex int) (docs []model.PmsProductAttribute, total int64, err error)
	GetPmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute) (model.PmsProductAttribute, error)
	CreatePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute) (model.PmsProductAttribute, error)
	UpdatePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute, id uint64) (update model.PmsProductAttribute, err error)
	BatchDeletePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute, id []uint64) (Result bool, err error)

	GetPmsProductAttributeCategoryPage(ctx context.Context, e *model.PmsProductAttributeCategory, pageSize int, pageIndex int) (docs []model.PmsProductAttributeCategory, total int64, err error)
	GetPmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory) (model.PmsProductAttributeCategory, error)
	CreatePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory) (model.PmsProductAttributeCategory, error)
	UpdatePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory, id uint64) (update model.PmsProductAttributeCategory, err error)
	BatchDeletePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory, id []uint64) (Result bool, err error)
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
	log := log.NewHelper(log.With(logger, "module", "dao"))
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
