package tools

import (
	"context"

	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(New)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)
	GetDBColumnsPage(dbname string, e *model.DBColumns, pageSize int, pageIndex int) ([]model.DBColumns, int64, error)
	GetDBColumnsList(dbname string, e *model.DBColumns) ([]model.DBColumns, error)

	GetDBTablesPage(dbname string, e *model.DBTables, pageSize int, pageIndex int) ([]model.DBTables, int64, error)
	GetDBTables(dbname string, e *model.DBTables) (model.DBTables, error)
}

// dao dao.
type dao struct {
	orm *gorm.DB
	log *log.Helper
}

// New new a dao and return.
func New(c *conf.App, logger log.Logger) (d Dao, err error) {
	return newDao(c, logger)
}

func newDao(c *conf.App, logger log.Logger) (d *dao, err error) {
	log := log.NewHelper("dao/tools", logger)
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: c.Tools.Source,
		}), &gorm.Config{})
	if err != nil {
		log.Error("open error")
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
