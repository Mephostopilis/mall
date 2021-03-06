package dao

import (
	"context"
	"edu/service/oms/internal/model"
)

//CreateOmsOrderItem 创建OmsOrderItem
func (d *dao) CreateOmsOrderItem(ctx context.Context, e *model.OmsOrderItem) (model.OmsOrderItem, error) {
	var doc model.OmsOrderItem
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取OmsOrderItem
func (d *dao) GetOmsOrderItem(ctx context.Context, e *model.OmsOrderItem) (model.OmsOrderItem, error) {
	var doc model.OmsOrderItem
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取OmsOrderItem带分页
func (d *dao) GetOmsOrderItemPage(ctx context.Context, e *model.OmsOrderItem, pageSize int, pageIndex int) (docs []model.OmsOrderItem, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取OmsOrderItem列表
func (d *dao) GetOmsOrderItemList(ctx context.Context, e *model.OmsOrderItem) (docs []model.OmsOrderItem, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新OmsOrderItem
func (d *dao) UpdateOmsOrderItem(ctx context.Context, e *model.OmsOrderItem, id uint64) (update model.OmsOrderItem, err error) {
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

// 删除OmsOrderItem
func (d *dao) DeleteOmsOrderItem(ctx context.Context, e *model.OmsOrderItem, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.OmsOrderItem{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteOmsOrderItem(ctx context.Context, e *model.OmsOrderItem, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.OmsOrderItem{}).Error; err != nil {
		return
	}
	Result = true
	return
}
