package dao

import (
	"context"
	"edu/service/oms/internal/model"
)

//CreateOmsOrder 创建OmsOrder
func (d *dao) CreateOmsOrder(ctx context.Context, e *model.OmsOrder) (model.OmsOrder, error) {
	var doc model.OmsOrder
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取OmsOrder
func (d *dao) GetOmsOrder(ctx context.Context, e *model.OmsOrder) (model.OmsOrder, error) {
	var doc model.OmsOrder
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取OmsOrder带分页
func (d *dao) GetOmsOrderPage(ctx context.Context, e *model.OmsOrder, pageSize int, pageIndex int) (docs []model.OmsOrder, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取OmsOrder列表
func (d *dao) GetOmsOrderList(ctx context.Context, e *model.OmsOrder) (docs []model.OmsOrder, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新OmsOrder
func (d *dao) UpdateOmsOrder(ctx context.Context, e *model.OmsOrder, id uint64) (update model.OmsOrder, err error) {
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

// 删除OmsOrder
func (d *dao) DeleteOmsOrder(ctx context.Context, e *model.OmsOrder, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.OmsOrder{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteOmsOrder(ctx context.Context, e *model.OmsOrder, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.OmsOrder{}).Error; err != nil {
		return
	}
	Result = true
	return
}
