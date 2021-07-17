package dao

import (
	pb "edu/api/sso/v1"
	"edu/pkg/ecode"

	"edu/service/sso/internal/model"
)

// 获取用户数据
func (d *dao) GetSysUserInfo(e *model.SysUser) (userView model.SysUser, err error) {
	table := d.orm.Table(e.TableName()).Select([]string{"ums_user.*", "ums_role.role_name"})
	table = table.Joins("left join ums_role on ums_user.role_id=ums_role.role_id")
	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if err = table.First(&userView).Error; err != nil {
		err = ecode.WrapError(err)
		return
	}
	return
}

func (d *dao) GetSysUserPage(e *model.SysUser, pageSize int, pageIndex int) (docs []model.SysUser, total int64, err error) {

	table := d.orm.Table(e.TableName()).Select([]string{"ums_user.*", "ums_role.role_name"})
	table = table.Joins("left join ums_role on ums_user.role_id=ums_role.role_id")

	if e.Username != "" {
		table = table.Where("ums_user.username = ?", e.Username)
	}
	if e.Status != "" {
		table = table.Where("ums_user.status = ?", e.Status)
	}

	if e.Phone != "" {
		table = table.Where("ums_user.phone = ?", e.Phone)
	}

	// if e.DeptId != 0 {
	// 	table = table.Where("sys_user.dept_id in (select dept_id from sys_dept where dept_path like ? )", "%"+tools.IntToString(e.DeptId)+"%")
	// }

	if err = table.Where("ums_user.deleted_at IS NULL").Count(&total).Error; err != nil {
		return
	}

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return nil, 0, err
	}

	return
}

//添加
func (d *dao) InsertSysUser(e *model.SysUser) (id uint64, err error) {
	if err = e.Encrypt(); err != nil {
		return
	}

	// check 用户名
	var count int64
	d.orm.Table(e.TableName()).Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = pb.ErrUsername("")
		return
	}

	//添加数据
	if err = d.orm.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}

//修改
func (d *dao) UpdateSysUser(e *model.SysUser, id uint64) (update model.SysUser, err error) {
	if e.Password != "" {
		if err = e.Encrypt(); err != nil {
			return
		}
	}
	if err = d.orm.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if e.RoleId == 0 {
		e.RoleId = update.RoleId
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) BatchDeleteSysUser(e *model.SysUser, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("user_id in (?)", id).Delete(&model.SysUser{}).Error; err != nil {
		return
	}
	Result = true
	return
}
