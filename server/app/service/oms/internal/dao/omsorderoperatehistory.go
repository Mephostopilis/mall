package dao

import (
	"context"
	"edu/service/oms/internal/model"
)

//CreateOmsOrderOperateHistory 创建OmsOrderOperateHistory
func (d *dao) CreateOmsOrderOperateHistory(ctx context.Context, e *model.OmsOrderOperateHistory) (model.OmsOrderOperateHistory, error) {
	var doc model.OmsOrderOperateHistory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取OmsOrderOperateHistory
func (d *dao) GetOmsOrderOperateHistory(ctx context.Context, e *model.OmsOrderOperateHistory) (model.OmsOrderOperateHistory, error) {
	var doc model.OmsOrderOperateHistory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取OmsOrderOperateHistory带分页
func (d *dao) GetOmsOrderOperateHistoryPage(ctx context.Context, e *model.OmsOrderOperateHistory, pageSize int, pageIndex int) (docs []model.OmsOrderOperateHistory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取OmsOrderOperateHistory列表
func (d *dao) GetOmsOrderOperateHistoryList(ctx context.Context, e *model.OmsOrderOperateHistory) (docs []model.OmsOrderOperateHistory, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新OmsOrderOperateHistory
func (d *dao) UpdateOmsOrderOperateHistory(ctx context.Context, e *model.OmsOrderOperateHistory, id uint64) (update model.OmsOrderOperateHistory, err error) {
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

// 删除OmsOrderOperateHistory
func (d *dao) DeleteOmsOrderOperateHistory(ctx context.Context, e *model.OmsOrderOperateHistory, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.OmsOrderOperateHistory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteOmsOrderOperateHistory(ctx context.Context, e *model.OmsOrderOperateHistory, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.OmsOrderOperateHistory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
