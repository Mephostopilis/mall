package dao

import (
	"context"
	"edu/service/oms/internal/model"
)

//CreateOmsOrderReturnReason 创建OmsOrderReturnReason
func (d *dao) CreateOmsOrderReturnReason(ctx context.Context, e *model.OmsOrderReturnReason) (model.OmsOrderReturnReason, error) {
	var doc model.OmsOrderReturnReason
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取OmsOrderReturnReason
func (d *dao) GetOmsOrderReturnReason(ctx context.Context, e *model.OmsOrderReturnReason) (model.OmsOrderReturnReason, error) {
	var doc model.OmsOrderReturnReason
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取OmsOrderReturnReason带分页
func (d *dao) GetOmsOrderReturnReasonPage(ctx context.Context, e *model.OmsOrderReturnReason, pageSize int, pageIndex int) (docs []model.OmsOrderReturnReason, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取OmsOrderReturnReason列表
func (d *dao) GetOmsOrderReturnReasonList(ctx context.Context, e *model.OmsOrderReturnReason) (docs []model.OmsOrderReturnReason, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新OmsOrderReturnReason
func (d *dao) UpdateOmsOrderReturnReason(ctx context.Context, e *model.OmsOrderReturnReason, id uint64) (update model.OmsOrderReturnReason, err error) {
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

// 删除OmsOrderReturnReason
func (d *dao) DeleteOmsOrderReturnReason(ctx context.Context, e *model.OmsOrderReturnReason, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.OmsOrderReturnReason{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteOmsOrderReturnReason(ctx context.Context, e *model.OmsOrderReturnReason, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.OmsOrderReturnReason{}).Error; err != nil {
		return
	}
	Result = true
	return
}
