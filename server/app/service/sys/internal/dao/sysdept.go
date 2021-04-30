package dao

import (
	"fmt"

	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
)

func (d *dao) CreateSysDept(e *model.SysDept) (doc model.SysDept, err error) {
	table := d.orm.Begin()
	deptPath := "/" + fmt.Sprintf("%d", e.DeptId)
	if int(e.ParentId) != 0 {
		var deptP model.SysDept
		if err = table.Table(e.TableName()).Where("dept_id = ?", e.ParentId).First(&deptP).Error; err != nil {
			table.Rollback()
			return
		}
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	e.DeptPath = deptPath
	result := table.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		table.Rollback()
		err = result.Error
		return
	}

	if err = table.Commit().Error; err != nil {
		return
	}
	return
}

func (d *dao) GetSysDept(e *model.SysDept) (doc model.SysDept, err error) {
	table := d.orm.Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if err = table.First(&doc).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetSysDeptList(e *model.SysDept) ([]model.SysDept, error) {
	var doc []model.SysDept

	table := d.orm.Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetSysDeptPage(e *model.SysDept) ([]model.SysDept, error) {
	var doc []model.SysDept

	table := d.orm.Select("*").Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}
	if e.DeptPath != "" {
		table = table.Where("deptPath like %?%", e.DeptPath)
	}
	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (d *dao) UpdateSysDept(e *model.SysDept, id int) (update model.SysDept, err error) {
	if err = d.orm.Table(e.TableName()).Where("dept_id = ?", id).First(&update).Error; err != nil {
		return
	}
	deptPath := "/" + fmt.Sprintf("%d", e.DeptId)
	if int(e.ParentId) != 0 {
		var deptP model.SysDept
		d.orm.Table(e.TableName()).Where("dept_id = ?", e.ParentId).First(&deptP)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	e.DeptPath = deptPath

	if e.DeptPath != "" && e.DeptPath != update.DeptPath {
		return update, errors.OutOfRange("table", "上级部门不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据

	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) DeleteSysDept(e *model.SysDept, id int) (success bool, err error) {

	// user := model.SysUser{}
	// user.DeptId = id
	// userlist, err := d.GetSysUserList(&user)
	// if err != nil {
	// 	return
	// }
	// if len(userlist) <= 0 {
	// 	err = ecode.ErrDeptId
	// 	return
	// }

	tx := d.orm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		success = false
		return
	}

	if err = tx.Table(e.TableName()).Where("dept_id = ?", id).Delete(&model.SysDept{}).Error; err != nil {
		success = false
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		success = false
		return
	}
	success = true

	return
}
