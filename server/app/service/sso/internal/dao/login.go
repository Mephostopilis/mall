package dao

import (
	"edu/pkg/tools"
	"edu/service/sso/internal/model"
)

func (d *dao) GetUser(u *model.Login) (user model.SysUser, role model.SysRole, e error) {
	e = d.orm.Table("ums_user").Where("username = ? ", u.Username).Find(&user).Error
	if e != nil {
		return
	}
	_, e = tools.CompareHashAndPassword(user.Password, u.Password)
	if e != nil {
		return
	}
	e = d.orm.Table("ums_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if e != nil {
		return
	}
	return
}
