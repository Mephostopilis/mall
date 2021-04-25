package dao

import (
	"context"
	"edu/service/sso/internal/model"
)

func (d *dao) GetJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog) (model.SysLoginLog, error) {
	return d.GetLoginLog(e)
}

func (d *dao) GetJwtAuthLoginLogPage(ctx context.Context, e *model.SysLoginLog, pageSize int, pageIndex int) ([]model.SysLoginLog, int64, error) {
	return d.GetLoginLogPage(e, pageSize, pageIndex)
}

func (d *dao) CreateJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog) (model.SysLoginLog, error) {
	return d.CreateLoginLog(e)
}

func (d *dao) UpdateJwtAuthLoginLog(ctx context.Context, e *model.SysLoginLog, id int) (update model.SysLoginLog, err error) {
	return d.UpdateLoginLog(e, id)
}

func (d *dao) GetJwtAuthUser(ctx context.Context, u *model.Login) (user model.SysUser, role model.SysRole, e error) {
	return d.GetUser(u)
}
