package dao

import (
	pb "edu/api/sso/v1"
	"edu/service/sso/internal/model"
)

func (d *dao) GetSysRolePage(role *model.SysRole, pageSize int, pageIndex int) ([]model.SysRole, int64, error) {
	var doc []model.SysRole

	table := d.orm.Table(role.TableName())
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if role.Status != "" {
		table = table.Where("status = ?", role.Status)
	}
	if role.RoleKey != "" {
		table = table.Where("role_key = ?", role.RoleKey)
	}

	// 数据权限控制
	// dataPermission := new(model.DataPermission)
	// dataPermission.UserId, _ = tools.StringToInt(role.DataScope)
	// table, err := d.GetDataScope(table, "sys_role", dataPermission)
	// if err != nil {
	// 	return nil, 0, err
	// }
	var count int64

	if err := table.Order("role_sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (d *dao) GetSysRole(role *model.SysRole) (SysRole model.SysRole, err error) {
	table := d.orm.Table(role.TableName())
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.First(&SysRole).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetSysRoleList(role *model.SysRole) (SysRole []model.SysRole, err error) {
	table := d.orm.Table(role.TableName())
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.Order("role_sort").Find(&SysRole).Error; err != nil {
		return
	}
	return
}

func (d *dao) InsertSysRole(role *model.SysRole) (id int, err error) {
	i := int64(0)
	if err = d.orm.Table(role.TableName()).Where("role_name=? or role_key = ?", role.RoleName, role.RoleKey).Count(&i).Error; err != nil {
		return
	}
	if i > 0 {
		return 0, pb.ErrOutOfRange("角色名称或者角色标识已经存在！")
	}
	role.UpdateBy = ""
	result := d.orm.Table(role.TableName()).Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.RoleId
	return
}

func (d *dao) GetRoleDeptId(role *model.SysRole) ([]int, error) {
	deptIds := make([]int, 0)
	deptList := make([]model.DeptIdList, 0)
	if err := d.orm.Table("sys_role_dept").Select("sys_role_dept.dept_id").Joins("LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id").Where("role_id = ? ", role.RoleId).Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.dept_id=sys_role_dept.dept_id where role_id =? )", role.RoleId).Find(&deptList).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}

	return deptIds, nil
}

//修改
func (d *dao) UpdateSysRole(role *model.SysRole, id int) (update model.SysRole, err error) {
	if err = d.orm.Table(role.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if role.RoleName != "" && role.RoleName != update.RoleName {
		return update, pb.ErrOutOfRange("角色名称不允许修改！")
	}

	if role.RoleKey != "" && role.RoleKey != update.RoleKey {
		return update, pb.ErrOutOfRange("角色标识不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(role.TableName()).Model(&update).Updates(&role).Error; err != nil {
		return
	}
	return
}

//批量删除
func (d *dao) BatchDeleteSysRole(role *model.SysRole, id []int64) (Result bool, err error) {
	if err = d.orm.Table(role.TableName()).Where("role_id in (?)", id).Delete(model.SysRole{}).Error; err != nil {
		return
	}
	Result = true
	return
}
