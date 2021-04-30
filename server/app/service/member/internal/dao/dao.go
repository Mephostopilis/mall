package dao

import (
	"context"

	"edu/service/member/internal/conf"
	"edu/service/member/internal/model"

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

	GetMemberPage(ctx context.Context, e *model.Member, pageSize int, pageIndex int) (docs []model.Member, total int64, err error)
	GetMember(ctx context.Context, e *model.Member) (model.Member, error)
	CreateMember(ctx context.Context, e *model.Member) (model.Member, error)
	UpdateMember(ctx context.Context, e *model.Member, id uint64) (update model.Member, err error)
	BatchDeleteMember(ctx context.Context, e *model.Member, id []int) (Result bool, err error)

	GetMemberAssetsPage(ctx context.Context, e *model.MemberAssets, pageSize int, pageIndex int) (docs []model.MemberAssets, total int64, err error)
	GetMemberAssets(ctx context.Context, e *model.MemberAssets) (model.MemberAssets, error)
	CreateMemberAssets(ctx context.Context, e *model.MemberAssets) (model.MemberAssets, error)
	UpdateMemberAssets(ctx context.Context, e *model.MemberAssets, id uint64) (update model.MemberAssets, err error)
	BatchDeleteMemberAssets(ctx context.Context, e *model.MemberAssets, id []int) (Result bool, err error)
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
