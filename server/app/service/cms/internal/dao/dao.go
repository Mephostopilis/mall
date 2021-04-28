package dao

import (
	"context"

	"edu/service/cms/internal/conf"
	"edu/service/cms/internal/model"

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

	GetCmsHelpCategoryPage(ctx context.Context, e *model.CmsHelpCategory, pageSize int, pageIndex int) (docs []model.CmsHelpCategory, total int64, err error)
	GetCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory) (model.CmsHelpCategory, error)
	CreateCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory) (model.CmsHelpCategory, error)
	UpdateCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory, id uint64) (update model.CmsHelpCategory, err error)
	BatchDeleteCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory, id []uint64) (Result bool, err error)

	GetCmsHelpPage(ctx context.Context, e *model.CmsHelp, pageSize int, pageIndex int) (docs []model.CmsHelp, total int64, err error)
	GetCmsHelp(ctx context.Context, e *model.CmsHelp) (model.CmsHelp, error)
	CreateCmsHelp(ctx context.Context, e *model.CmsHelp) (model.CmsHelp, error)
	UpdateCmsHelp(ctx context.Context, e *model.CmsHelp, id uint64) (update model.CmsHelp, err error)
	BatchDeleteCmsHelp(ctx context.Context, e *model.CmsHelp, id []uint64) (Result bool, err error)

	GetCmsSubjectPage(ctx context.Context, e *model.CmsSubject, pageSize int, pageIndex int) (docs []model.CmsSubject, total int64, err error)
	GetCmsSubject(ctx context.Context, e *model.CmsSubject) (model.CmsSubject, error)
	CreateCmsSubject(ctx context.Context, e *model.CmsSubject) (model.CmsSubject, error)
	UpdateCmsSubject(ctx context.Context, e *model.CmsSubject, id uint64) (update model.CmsSubject, err error)
	BatchDeleteCmsSubject(ctx context.Context, e *model.CmsSubject, id []uint64) (Result bool, err error)
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
