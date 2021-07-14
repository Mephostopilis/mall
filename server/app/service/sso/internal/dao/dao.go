package dao

import (
	"context"

	"edu/service/sso/internal/conf"
	"edu/service/sso/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4"
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

	// GetByID(ctx context.Context, id string) (cli oauth2.ClientInfo, err error)
	GetByID(ctx context.Context, id string) (cli oauth2.ClientInfo, err error)

	Create(info oauth2.TokenInfo) (err error)
	RemoveByCode(code string) error
	RemoveByAccess(access string) error
	RemoveByRefresh(refresh string) error
	GetByCode(code string) (oauth2.TokenInfo, error)
	GetByAccess(access string) (oauth2.TokenInfo, error)
	GetByRefresh(refresh string) (oauth2.TokenInfo, error)

	CreateSsoApp(ctx context.Context, e *model.SsoApp) (model.SsoApp, error)
	GetSsoApp(ctx context.Context, e *model.SsoApp) (model.SsoApp, error)
	GetSsoAppPage(ctx context.Context, e *model.SsoApp, pageSize int, pageIndex int) ([]model.SsoApp, int64, error)
	UpdateSsoApp(ctx context.Context, e *model.SsoApp, id string) (update model.SsoApp, err error)
	DeleteSsoApp(ctx context.Context, e *model.SsoApp, id string) (success bool, err error)
	BatchDeleteSsoApp(ctx context.Context, e *model.SsoApp, id []int) (Result bool, err error)

	CreateSsoToken(ctx context.Context, e *model.SsoToken) (model.SsoToken, error)
	GetSsoToken(ctx context.Context, e *model.SsoToken) (model.SsoToken, error)
	GetSsoTokenPage(ctx context.Context, e *model.SsoToken, pageSize int, pageIndex int) ([]model.SsoToken, int64, error)
	UpdateSsoToken(ctx context.Context, e *model.SsoToken, id string) (update model.SsoToken, err error)
	DeleteSsoToken(ctx context.Context, e *model.SsoToken, id string) (success bool, err error)
	BatchDeleteSsoToken(ctx context.Context, e *model.SsoToken, id []int) (Result bool, err error)

	GetSysUserPage(e *model.SysUser, pageSize int, pageIndex int) ([]model.SysUser, int64, error)
	GetSysUserInfo(e *model.SysUser) (SysUserView model.SysUser, err error)
	InsertSysUser(e *model.SysUser) (id uint64, err error)
	UpdateSysUser(e *model.SysUser, id uint64) (update model.SysUser, err error)
	BatchDeleteSysUser(e *model.SysUser, id []int) (Result bool, err error)

	GetSysRolePage(role *model.SysRole, pageSize int, pageIndex int) ([]model.SysRole, int64, error)
	GetSysRole(role *model.SysRole) (SysRole model.SysRole, err error)
	GetSysRoleList(role *model.SysRole) (SysRole []model.SysRole, err error)
	InsertSysRole(role *model.SysRole) (id int, err error)
	GetRoleDeptId(role *model.SysRole) ([]int, error)
	UpdateSysRole(role *model.SysRole, id int) (update model.SysRole, err error)
	BatchDeleteSysRole(role *model.SysRole, ids []int64) (Result bool, err error)

	GetLoginLogPage(e *model.SysLoginLog, pageSize int, pageIndex int) ([]model.SysLoginLog, int64, error)
	GetLoginLog(e *model.SysLoginLog) (model.SysLoginLog, error)
	CreateLoginLog(e *model.SysLoginLog) (model.SysLoginLog, error)
	UpdateLoginLog(e *model.SysLoginLog, id int) (update model.SysLoginLog, err error)
	BatchDeleteLoginLog(e *model.SysLoginLog, id []int) (Result bool, err error)

	GetUser(u *model.Login) (user model.SysUser, role model.SysRole, e error)

	GetJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog) (model.SysLoginLog, error)
	GetJwtAuthLoginLogPage(ctx context.Context, e *model.SysLoginLog, pageSize int, pageIndex int) ([]model.SysLoginLog, int64, error)
	CreateJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog) (model.SysLoginLog, error)
	UpdateJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog, id int) (update model.SysLoginLog, err error)
	GetJwtAuthUser(ctx context.Context, u *model.Login) (user model.SysUser, role model.SysRole, e error)
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
	log := log.NewHelper(log.With(logger, "module", "dao/sso"))
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
