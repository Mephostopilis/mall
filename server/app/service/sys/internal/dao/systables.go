package dao

import (
	"edu/service/sys/internal/model"
)

func (d *dao) GetSysTablesPage(e *model.SysTables, pageSize int, pageIndex int) (doc []model.SysTables, count int64, err error) {
	table := d.orm.Table(e.TableName())
	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	if err = table.Where("`deleted_at` IS NULL").Count(&count).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetSysTables(e *model.SysTables) (doc model.SysTables, err error) {
	table := d.orm.Table(e.TableName())
	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableId != 0 {
		table = table.Where("table_id = ?", e.TableId)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}
	if err = table.First(&doc).Error; err != nil {
		return
	}
	// 获取所有列
	var cols []model.SysColumns
	table = d.orm.Table("sys_columns")
	table = table.Where("table_id = ?", e.TableId)
	if err = table.Find(&cols).Error; err != nil {
		return
	}
	doc.Columns = cols
	return
}

func (d *dao) CreateSysTables(e *model.SysTables) (doc model.SysTables, err error) {
	table := d.orm.Begin()

	result := table.Table("sys_tables").Create(e)
	if result.Error != nil {
		if err = table.Rollback().Error; err != nil {
			return
		}
		err = result.Error
		return
	}
	// d.log.Debugf("xx %v", result)
	// if err = table.Table("sys_tables").Where("table_name = ?", e.TBName).First(&doc).Error; err != nil {
	// 	return
	// }
	doc = *e
	for i := 0; i < len(e.Columns); i++ {
		col := e.Columns[i]
		col.TableId = doc.TableId
		result := table.Table("sys_columns").Create(&col)
		if result.Error != nil {
			if err = table.Rollback().Error; err != nil {
				return
			}
			err = result.Error
			return
		}
	}
	if err = table.Commit().Error; err != nil {
		return
	}
	return
}

func (d *dao) UpdateSysTables(e *model.SysTables) (update model.SysTables, err error) {
	table := d.orm.Begin()
	if err = table.Table("sys_tables").First(&update, e.TableId).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = table.Table("sys_tables").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	if err = table.Commit().Error; err != nil {
		return
	}
	return
}

func (d *dao) DeleteSysTables(e *model.SysTables) (success bool, err error) {
	table := d.orm.Begin()
	if err = table.Table("sys_tables").Delete(model.SysTables{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	if err = table.Table("sys_columns").Delete(model.SysColumns{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	if err = table.Commit().Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) BatchDeleteSysTables(e *model.SysTables, ids []int) (success bool, err error) {
	table := d.orm.Begin()
	if err = table.Table("sys_tables").Delete(model.SysTables{}, "table_id in ?", ids).Error; err != nil {
		success = false
		return
	}
	if err = table.Table("sys_columns").Delete(model.SysColumns{}, "table_id in ?", ids).Error; err != nil {
		success = false
		return
	}
	if err = table.Commit().Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
