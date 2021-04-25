package dao

import (
	"fmt"
	"time"

	"edu/service/sys/internal/model"
)

func (d *dao) InsertSysDeptMenu(rm *model.SysDeptMenu, uid uint64, deptId int, menuIds []int) (bool, error) {
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := "INSERT INTO `sys_dept_menu` (`role_id`,`dept_id`, `create_by`, `created_at`) VALUES "

	now := time.Now()
	for i := 0; i < len(menuIds); i++ {
		if len(menuIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,%s,%v);", deptId, menuIds[i], fmt.Sprint(uid), now)
		} else {
			sql += fmt.Sprintf("(%d,%d,%s,%v),", deptId, menuIds[i], fmt.Sprint(uid), now)
		}
	}
	d.orm.Exec(sql)

	return true, nil
}

func (d *dao) DeleteSysDeptMenu(rm *model.SysDeptMenu, deptId int) (bool, error) {
	if err := d.orm.Table("sys_dept_menu").Where("dept_id = ?", deptId).Delete(&rm).Error; err != nil {
		return false, err
	}
	return true, nil
}
