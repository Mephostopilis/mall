package dao

import (
	"edu/pkg/tools"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
)

func (d *dao) GetMenusByRoleName(e *model.Menu, rolename string) (Menus []model.Menu, err error) {
	table := d.orm.Table(e.TableName()).Select("sys_menu.*").Joins("left join sys_role_menu on sys_role_menu.menu_id=sys_menu.menu_id")
	table = table.Where("sys_role_menu.role_name=? and menu_type in ('M','C')", rolename)
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetMenus(e *model.Menu) (Menus []model.Menu, err error) {
	table := d.orm.Table(e.TableName())
	if e.MenuName != "" {
		table = table.Where("menu_name = ?", e.MenuName)
	}
	if e.Path != "" {
		table = table.Where("path = ?", e.Path)
	}
	if e.Action != "" {
		table = table.Where("action = ?", e.Action)
	}
	if e.MenuType != "" {
		table = table.Where("menu_type = ?", e.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetMenuPage(e *model.Menu) (Menus []model.Menu, err error) {
	table := d.orm.Table(e.TableName())
	if e.MenuName != "" {
		table = table.Where("menu_name = ?", e.MenuName)
	}
	if e.Title != "" {
		table = table.Where("title = ?", e.Title)
	}
	// if e.Visible != "" {
	// 	table = table.Where("visible = ?", e.Visible)
	// }
	if e.MenuType != "" {
		table = table.Where("menu_type = ?", e.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetMenuByMenuId(e *model.Menu) (Menu model.Menu, err error) {
	table := d.orm.Table(e.TableName())
	table = table.Where("menu_id = ?", e.MenuId)
	if err = table.Find(&Menu).Error; err != nil {
		return
	}
	return
}

func (d *dao) initPaths(menu *model.Menu) (err error) {
	parentMenu := new(model.Menu)
	if int(menu.ParentId) != 0 {
		d.orm.Table("sys_menu").Where("menu_id = ?", menu.ParentId).First(parentMenu)
		if parentMenu.Paths == "" {
			err = errors.OutOfRange("table", "父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return
		}
		menu.Paths = parentMenu.Paths + "/" + tools.IntToString(int(menu.MenuId))
	} else {
		menu.Paths = "/0/" + tools.IntToString(int(menu.MenuId))
	}
	return
}

func (d *dao) CreateMenu(e *model.Menu) (id int32, err error) {
	if err = d.initPaths(e); err != nil {
		return
	}
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = e.MenuId
	return
}

func (d *dao) UpdateMenu(e *model.Menu, id int) (update model.Menu, err error) {
	if err = d.orm.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.initPaths(e); err != nil {
		return
	}
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}

func (d *dao) DeleteMenu(e *model.Menu, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("menu_id = ?", id).Delete(&model.Menu{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) GetRoleMenuPermisFromSysMenu(rm *model.SysRoleMenu) ([]string, error) {
	var r []model.Menu
	table := d.orm.Select("sys_menu.permission").Table("sys_menu").Joins("left join sys_role_menu on sys_menu.menu_id = sys_role_menu.menu_id")
	table = table.Where("role_id = ?", rm.RoleId)

	table = table.Where("sys_menu.menu_type in('F','C')")
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list, nil
}
