package dao

import (
	"context"

	ssopb "edu/api/sso/v1"
	"edu/service/sys/internal/model"
)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	// Article(c context.Context, id int64) (*model.Article, error)

	GetSysConfig(e *model.SysConfig) (model.SysConfig, error)
	GetSysConfigPage(e *model.SysConfig, pageSize int, pageIndex int) ([]model.SysConfig, int64, error)
	CreateSysConfig(ctx context.Context, e *model.SysConfig) (model.SysConfig, error)
	UpdateSysConfig(e *model.SysConfig, id int) (update model.SysConfig, err error)
	DeleteSysConfig(e *model.SysConfig, id int) (success bool, err error)
	BatchDeleteSysConfig(e *model.SysConfig, ids []int) (Result bool, err error)

	CreateDictData(e *model.SysDictData) (model.SysDictData, error)
	GetDictDataByCode(e *model.SysDictData) (model.SysDictData, error)
	GetDictData(e *model.SysDictData) ([]model.SysDictData, error)
	GetDictDataPage(e *model.SysDictData, pageSize int, pageIndex int) ([]model.SysDictData, int64, error)
	UpdateDictData(e *model.SysDictData, id int) (update model.SysDictData, err error)
	DeleteDictData(e *model.SysDictData, id int) (success bool, err error)
	BatchDeleteDictData(e *model.SysDictData, id []int) (Result bool, err error)

	// GetMenu(e *model.Menu) (m []model.Menu, err error)
	// GetMenu(e *model.Menu, permission *ssopb.DataPermission) (m []model.Menu, err error)
	// GetMenuRole(e *model.Menu, rolename string) (m []model.Menu, err error)
	// GetMenusRoleList(e *model.Menu) (Menus []model.Menu, err error)
	GetMenusByRoleName(e *model.Menu, rolename string) (Menus []model.Menu, err error)
	GetMenus(e *model.Menu) (Menus []model.Menu, err error)
	GetMenuPage(e *model.Menu) (Menus []model.Menu, err error)
	GetMenuByMenuId(e *model.Menu) (Menu model.Menu, err error)
	CreateMenu(e *model.Menu) (id int32, err error)
	UpdateMenu(e *model.Menu, id int) (update model.Menu, err error)
	DeleteMenu(e *model.Menu, id int) (success bool, err error)

	GetSysOperLog(e *model.SysOperLog) (model.SysOperLog, error)
	GetSysOperLogPage(e *model.SysOperLog, pageSize int, pageIndex int) ([]model.SysOperLog, int64, error)
	CreateSysOperLog(e *model.SysOperLog) (model.SysOperLog, error)
	UpdateSysOperLog(e *model.SysOperLog, id int) (update model.SysOperLog, err error)
	BatchDeleteSysOperLog(e *model.SysOperLog, id []int) (Result bool, err error)

	CreateSysJob(e *model.SysJob) (model.SysJob, error)
	GetSysJob(e *model.SysJob) (model.SysJob, error)
	GetSysJobPage(e *model.SysJob, pageSize int, pageIndex int) ([]model.SysJob, int64, error)
	GetSysJobList(e *model.SysJob) ([]model.SysJob, error)
	UpdateSysJob(e *model.SysJob, id int) (update model.SysJob, err error)
	RemoveSysJobAllEntryID(e *model.SysJob) (update model.SysJob, err error)
	RemoveSysJobEntryID(e *model.SysJob, entryID int) (update model.SysJob, err error)
	DeleteSysJob(e *model.SysJob, id int) (success bool, err error)
	BatchDeleteSysJob(e *model.SysJob, id []int) (Result bool, err error)

	GetRoleMenu(rm *model.SysRoleMenu, roleId int32) ([]model.SysRoleMenu, error)
	GetRoleMenuPermisFromSysMenu(rm *model.SysRoleMenu) ([]string, error)
	GetRoleMenuIDSFromSysMenu(rm *model.SysRoleMenu) ([]model.MenuPath, error)
	// DeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roleId int) (bool, error)
	DeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roleId int, roleKey string) (bool, error)
	// BatchDeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roleIds []int) (bool, error)
	BatchDeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roles []ssopb.Role) (bool, error)
	DeleteRoleMenu(rm *model.SysRoleMenu, RoleId int, MenuID int) (bool, error)
	InsertRoleMenu(rm *model.SysRoleMenu, roleId int, rolekey string, menuId []int) (bool, error)

	GetSysSetting(ctx context.Context, s *model.SysSetting) (create model.SysSetting, err error)
	UpdateSysSetting(ctx context.Context, s *model.SysSetting) (update model.SysSetting, err error)

	CreateDictType(e *model.SysDictType) (model.SysDictType, error)
	GetDictType(e *model.SysDictType) (model.SysDictType, error)
	GetDictTypeList(e *model.SysDictType) ([]model.SysDictType, error)
	GetDictTypePage(e *model.SysDictType, pageSize int, pageIndex int) ([]model.SysDictType, int64, error)
	UpdateDictType(e *model.SysDictType, id int) (update model.SysDictType, err error)
	DeleteDictType(e *model.SysDictType, id int) (success bool, err error)
	BatchDeleteDictType(e *model.SysDictType, id []int) (Result bool, err error)

	// GetSysColumnsList(e *model.SysColumns) ([]model.SysColumns, error)
	// CreateSysColumns(e *model.SysColumns) (model.SysColumns, error)
	// UpdateSysColumns(e *model.SysColumns) (update model.SysColumns, err error)

	GetSysTablesPage(e *model.SysTables, pageSize int, pageIndex int) ([]model.SysTables, int64, error)
	GetSysTables(e *model.SysTables) (model.SysTables, error)
	CreateSysTables(e *model.SysTables) (model.SysTables, error)
	UpdateSysTables(e *model.SysTables) (update model.SysTables, err error)
	DeleteSysTables(e *model.SysTables) (success bool, err error)
	BatchDeleteSysTables(e *model.SysTables, id []int) (Result bool, err error)

	GetSysResourcePage(ctx context.Context, e *model.SysResource, pageSize int, pageIndex int) (docs []model.SysResource, total int64, err error)
	GetSysResource(ctx context.Context, e *model.SysResource) (model.SysResource, error)
	CreateSysResource(ctx context.Context, e *model.SysResource) (model.SysResource, error)
}
