package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductVertifyRecord 创建PmsProductVertifyRecord
func (d *dao) CreatePmsProductVertifyRecord(ctx context.Context, e *model.PmsProductVertifyRecord) (model.PmsProductVertifyRecord, error) {
	var doc model.PmsProductVertifyRecord
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductVertifyRecord
func (d *dao) GetPmsProductVertifyRecord(ctx context.Context, e *model.PmsProductVertifyRecord) (model.PmsProductVertifyRecord, error) {
	var doc model.PmsProductVertifyRecord
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductVertifyRecord带分页
func (d *dao) GetPmsProductVertifyRecordPage(ctx context.Context, e *model.PmsProductVertifyRecord, pageSize int, pageIndex int) (docs []model.PmsProductVertifyRecord, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductVertifyRecord
func (d *dao) UpdatePmsProductVertifyRecord(ctx context.Context, e *model.PmsProductVertifyRecord, id int) (update model.PmsProductVertifyRecord, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除PmsProductVertifyRecord
func (d *dao) DeletePmsProductVertifyRecord(ctx context.Context, e *model.PmsProductVertifyRecord, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductVertifyRecord{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductVertifyRecord(ctx context.Context, e *model.PmsProductVertifyRecord, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductVertifyRecord{}).Error; err != nil {
		return
	}
	Result = true
	return
}
