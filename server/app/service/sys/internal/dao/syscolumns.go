package dao

// func (d *dao) GetSysColumnsList(e *model.SysColumns) ([]model.SysColumns, error) {
// 	var doc []model.SysColumns
// 	table := d.orm.Table("sys_columns")
// 	table = table.Where("table_id = ?", e.TableId)
// 	if err := table.Find(&doc).Error; err != nil {
// 		return nil, err
// 	}
// 	return doc, nil
// }

// func (d *dao) CreateSysColumns(e *model.SysColumns) (model.SysColumns, error) {
// 	var doc model.SysColumns
// 	result := d.orm.Table("sys_columns").Create(&e)
// 	if result.Error != nil {
// 		err := result.Error
// 		return doc, err
// 	}
// 	doc = *e
// 	return doc, nil
// }

// func (d *dao) UpdateSysColumns(e *model.SysColumns) (update model.SysColumns, err error) {
// 	if err = d.orm.Table("sys_columns").First(&update, e.ColumnId).Error; err != nil {
// 		return
// 	}

// 	//参数1:是要修改的数据
// 	//参数2:是修改的数据
// 	if err = d.orm.Table("sys_columns").Model(&update).Updates(&e).Error; err != nil {
// 		return
// 	}
// 	return
// }
