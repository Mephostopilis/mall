package dao

import (
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
)

func (d *dao) GetDictDataByCode(e *model.SysDictData) (model.SysDictData, error) {
	var doc model.SysDictData

	table := d.orm.Table(e.TableName())
	if e.DictCode != 0 {
		table = table.Where("dict_code = ?", e.DictCode)
	}
	if e.DictLabel != "" {
		table = table.Where("dict_label = ?", e.DictLabel)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetDictData(e *model.SysDictData) ([]model.SysDictData, error) {
	var doc []model.SysDictData

	table := d.orm.Table(e.TableName())
	if e.DictCode != 0 {
		table = table.Where("dict_code = ?", e.DictCode)
	}
	if e.DictLabel != "" {
		table = table.Where("dict_label = ?", e.DictLabel)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}

	if err := table.Order("dict_sort").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetDictDataPage(e *model.SysDictData, pageSize int, pageIndex int) ([]model.SysDictData, int64, error) {
	var doc []model.SysDictData

	table := d.orm.Table(e.TableName())
	if e.DictCode != 0 {
		table = table.Where("dict_code = ?", e.DictCode)
	}
	if e.DictType != "" {
		table = table.Where("dict_type = ?", e.DictType)
	}
	if e.DictLabel != "" {
		table = table.Where("dict_label = ?", e.DictLabel)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	var count int64
	if err := table.Where("`deleted_at` IS NULL").Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := table.Order("dict_sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

func (d *dao) CreateDictData(e *model.SysDictData) (model.SysDictData, error) {
	var doc model.SysDictData

	i := int64(0)
	d.orm.Table(e.TableName()).Where("dict_label=? or (dict_label=? and dict_value = ?)", e.DictLabel, e.DictValue).Count(&i)
	if i > 0 {
		return doc, pb.ErrOutOfRange("out", "?????????????????????????????????????????????")
	}

	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (d *dao) UpdateDictData(e *model.SysDictData, id int) (update model.SysDictData, err error) {
	if err = d.orm.Table(e.TableName()).Where("dict_code = ?", id).First(&update).Error; err != nil {
		return
	}

	if e.DictLabel != "" && e.DictLabel != update.DictLabel {
		return update, pb.ErrOutOfRange("out", "????????????????????????")
	}

	if e.DictValue != "" && e.DictValue != update.DictValue {
		return update, pb.ErrOutOfRange("out", "????????????????????????")
	}

	//??????1:?????????????????????
	//??????2:??????????????????
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) DeleteDictData(e *model.SysDictData, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("dict_code = ?", id).Delete(&model.SysDictData{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) BatchDeleteDictData(e *model.SysDictData, ids []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("dict_code in (?)", ids).Delete(&model.SysDictData{}).Error; err != nil {
		return
	}
	Result = true
	return
}
