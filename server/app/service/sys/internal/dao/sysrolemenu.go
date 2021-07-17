package dao

import (
	"fmt"

	ssopb "edu/api/sso/v1"

	"edu/service/sys/internal/model"
)

func (d *dao) GetRoleMenu(rm *model.SysRoleMenu, roleId int32) (r []model.SysRoleMenu, err error) {
	table := d.orm.Table("sys_role_menu").Where("role_id = ?", roleId)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (d *dao) GetRoleMenuIDSFromSysMenu(rm *model.SysRoleMenu) ([]model.MenuPath, error) {
	var r []model.MenuPath
	table := d.orm.Select("sys_menu.path").Table("sys_role_menu")
	table = table.Joins("left join sys_role on sys_role.role_id=sys_role_menu.role_id")
	table = table.Joins("left join sys_menu on sys_menu.id=sys_role_menu.menu_id")
	table = table.Where("sys_role.role_name = ? and sys_menu.type=1", rm.RoleName)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (d *dao) DeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roleId int, roleKey string) (bool, error) {
	tx := d.orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Table("sys_role_menu").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	sql3 := "delete from casbin_rule where v0= '" + roleKey + "';"
	if err := tx.Exec(sql3).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}

	return true, nil
}

func (d *dao) InsertRoleMenu(rm *model.SysRoleMenu, roleId int, rolekey string, menuId []int) (bool, error) {
	tx := d.orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}
	var menus []model.Menu
	if err := tx.Table("sys_menu").Where("menu_id in (?)", menuId).Find(&menus).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := "INSERT INTO `sys_role_menu` (`role_id`,`menu_id`,`role_name`) VALUES "

	sql2 := "INSERT INTO casbin_rule  (`p_type`,`v0`,`v1`,`v2`) VALUES "
	for i := 0; i < len(menus); i++ {
		menu := menus[i]
		if len(menus)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,'%s');", roleId, menus[i].MenuId, rolekey)
			if menus[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s');", rolekey, menus[i].Path, menus[i].Action)
			}
		} else {
			sql += fmt.Sprintf("(%d,%d,'%s'),", roleId, menu.MenuId, rolekey)
			if menus[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s'),", rolekey, menus[i].Path, menus[i].Action)
			}
		}
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	sql2 = sql2[0:len(sql2)-1] + ";"
	if err := tx.Exec(sql2).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}
	return true, nil
}

func (d *dao) DeleteRoleMenu(rm *model.SysRoleMenu, roleId int, menuID int) (bool, error) {
	table := d.orm.Table("sys_role_menu").Where("role_id = ?", roleId)
	if menuID != 0 {
		table = table.Where("menu_id = ?", menuID)
	}
	if err := table.Delete(&rm).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (d *dao) BatchDeleteRoleMenuFromSysRole(rm *model.SysRoleMenu, roles []ssopb.Role) (bool, error) {
	// 判断角色id是否有效

	tx := d.orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	roleIds := make([]int, 0)
	if err := tx.Table("sys_role_menu").Where("role_id in (?)", roleIds).Delete(&rm).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	// TODO: fix 数据从服务返回
	// var role []model.SysRole
	// if err := tx.Table("sys_role").Where("role_id in (?)", roleIds).Find(&role).Error; err != nil {
	// 	tx.Rollback()
	// 	return false, err
	// }
	sql := ""
	for i := 0; i < len(roles); i++ {
		sql += "delete from casbin_rule where v0= '" + roles[i].RoleName + "';"
	}
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Commit().Error; err != nil {
		return false, err
	}
	return true, nil
}
