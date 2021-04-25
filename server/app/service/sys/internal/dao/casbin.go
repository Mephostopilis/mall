package dao

import (
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

// 获取TbUsers
func (d *dao) GetCasbinRule(e *gormAdapter.CasbinRule) (gormAdapter.CasbinRule, error) {
	var doc gormAdapter.CasbinRule
	table := d.orm.Table("casbin_rule")

	if e.ID != 0 {
		table = table.Where("id = ?", e.ID)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取TbUsers带分页
func (d *dao) GetTbUsersPage(e *gormAdapter.CasbinRule, pageSize int, pageIndex int) ([]gormAdapter.CasbinRule, int64, error) {
	var doc []gormAdapter.CasbinRule
	table := d.orm.Table("casbin_rule")

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	// dataPermission := new(model.DataPermission)
	// dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	// table, err := d.GetDataScope(table, "casbin_rule", dataPermission)
	// if err != nil {
	// 	return nil, 0, err
	// }
	var count int64
	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 创建TbUsers
func (d *dao) CreateCasbinRule(e *gormAdapter.CasbinRule) (gormAdapter.CasbinRule, error) {
	var doc gormAdapter.CasbinRule
	result := d.orm.Table("casbin_rule").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 更新TbUsers
func (d *dao) UpdateCasbinRule(e *gormAdapter.CasbinRule, id int) (update gormAdapter.CasbinRule, err error) {
	if err = d.orm.Table("casbin_rule").Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table("casbin_rule").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除TbUsers
func (d *dao) DeleteCasbinRule(e *gormAdapter.CasbinRule, id int) (success bool, err error) {
	if err = d.orm.Table("casbin_rule").Where("uid = ?", id).Delete(&gormAdapter.CasbinRule{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCasbinRule(e *gormAdapter.CasbinRule, id []int) (Result bool, err error) {
	if err = d.orm.Table("casbin_rule").Where("uid in (?)", id).Delete(&gormAdapter.CasbinRule{}).Error; err != nil {
		return
	}
	Result = true
	return
}
