package dao

import (
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
)

func (d *dao) CreateDictType(e *model.SysDictType) (model.SysDictType, error) {
	var doc model.SysDictType

	i := int64(0)
	d.orm.Table(e.TableName()).Where("dict_name=? or dict_type = ?", e.DictName, e.DictType).Count(&i)
	if i > 0 {
		return doc, pb.ErrOutOfRange("table", "字典名称或者字典类型已经存在！")
	}

	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (d *dao) GetDictType(e *model.SysDictType) (model.SysDictType, error) {
	var doc model.SysDictType

	table := d.orm.Table(e.TableName())
	if e.DictId != 0 {
		table = table.Where("dict_id = ?", e.DictId)
	}
	if e.DictName != "" {
		table = table.Where("dict_name = ?", e.DictName)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetDictTypeList(e *model.SysDictType) ([]model.SysDictType, error) {
	var doc []model.SysDictType

	table := d.orm.Table(e.TableName())
	if e.DictId != 0 {
		table = table.Where("dict_id = ?", e.DictId)
	}
	if e.DictName != "" {
		table = table.Where("dict_name = ?", e.DictName)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetDictTypePage(e *model.SysDictType, pageSize int, pageIndex int) ([]model.SysDictType, int64, error) {
	var doc []model.SysDictType

	table := d.orm.Table(e.TableName())
	if e.DictId != 0 {
		table = table.Where("dict_id = ?", e.DictId)
	}
	if e.DictName != "" {
		table = table.Where("dict_name = ?", e.DictName)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.DictName)
	}

	var total int64
	if err := table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	return doc, total, nil
}

func (d *dao) UpdateDictType(e *model.SysDictType, id int) (update model.SysDictType, err error) {
	if err = d.orm.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if e.DictName != "" && e.DictName != update.DictName {
		return update, pb.ErrOutOfRange("table", "名称不允许修改！")
	}

	if e.DictType != "" && e.DictType != update.DictType {
		return update, pb.ErrOutOfRange("table", "类型不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) DeleteDictType(e *model.SysDictType, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("dict_id = ?", id).Delete(&model.SysDictData{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) BatchDeleteDictType(e *model.SysDictType, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("dict_id in (?)", id).Delete(&model.SysDictType{}).Error; err != nil {
		return
	}
	Result = true
	return
}
