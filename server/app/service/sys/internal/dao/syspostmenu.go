package dao

import (
	"fmt"

	"edu/service/sys/internal/model"
)

func (d *dao) InsertSysPostMenu(rm *model.SysPostMenu, postId int, menuIds []int) (bool, error) {
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := "INSERT INTO `sys_post_menu` (`post_id`,`menu_id`) VALUES "

	for i := 0; i < len(menuIds); i++ {
		if len(menuIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", postId, menuIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", postId, menuIds[i])
		}
	}
	d.orm.Exec(sql)

	return true, nil
}

func (d *dao) DeleteSysRoleDept(rm *model.SysPostMenu, postId int) (bool, error) {
	if err := d.orm.Table("sys_post_menu").Where("post_id = ?", postId).Delete(&rm).Error; err != nil {
		return false, err
	}
	return true, nil
}
