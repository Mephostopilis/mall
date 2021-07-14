package dao

import (
	"context"

	"edu/service/tiku/internal/conf"
	"edu/service/tiku/internal/model"

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

	GetChoicesPage(ctx context.Context, e *model.Choices, pageSize int, pageIndex int) ([]model.Choices, int64, error)
	GetChoices(ctx context.Context, e *model.Choices) (model.Choices, error)
	CreateChoices(ctx context.Context, e *model.Choices) (model.Choices, error)
	UpdateChoices(ctx context.Context, e *model.Choices, id uint64) (update model.Choices, err error)
	BatchDeleteChoices(ctx context.Context, e *model.Choices, id []uint64) (Result bool, err error)

	GetExercisePage(ctx context.Context, e *model.Exercise, pageSize int, pageIndex int) ([]model.Exercise, int64, error)
	GetExercise(ctx context.Context, e *model.Exercise) (model.Exercise, error)
	CreateExercise(ctx context.Context, e *model.Exercise) (model.Exercise, error)
	UpdateExercise(ctx context.Context, e *model.Exercise, id uint64) (update model.Exercise, err error)
	BatchDeleteExercise(ctx context.Context, e *model.Exercise, id []uint64) (Result bool, err error)

	GetUserChoices(ctx context.Context, e *model.UserChoices) (model.UserChoices, error)
	GetUserChoicesPage(ctx context.Context, e *model.UserChoices, pageSize int, pageIndex int) ([]model.UserChoices, int64, error)
	CreateUserChoices(ctx context.Context, e *model.UserChoices) (model.UserChoices, error)

	GetUserExercise(ctx context.Context, e *model.UserExercise) (model.UserExercise, error)
	GetUserExercisePage(ctx context.Context, e *model.UserExercise, pageSize int, pageIndex int) ([]model.UserExercise, int64, error)
	CreateUserExercise(ctx context.Context, e *model.UserExercise) (model.UserExercise, error)
	UpdateUserExercise(ctx context.Context, e *model.UserExercise, id int) (update model.UserExercise, err error)
	DeleteUserExercise(ctx context.Context, e *model.UserExercise, id int) (success bool, err error)
	BatchDeleteUserExercise(ctx context.Context, e *model.UserExercise, id []int) (Result bool, err error)
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
