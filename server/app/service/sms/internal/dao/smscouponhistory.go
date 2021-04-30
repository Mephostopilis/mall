package dao

import (
	"context"
	"edu/service/sms/internal/model"
)

//CreateSmsCouponHistory 创建SmsCouponHistory
func (d *dao) CreateSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory) (model.SmsCouponHistory, error) {
	var doc model.SmsCouponHistory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SmsCouponHistory
func (d *dao) GetSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory) (model.SmsCouponHistory, error) {
	var doc model.SmsCouponHistory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SmsCouponHistory带分页
func (d *dao) GetSmsCouponHistoryPage(ctx context.Context, e *model.SmsCouponHistory, pageSize int, pageIndex int) (docs []model.SmsCouponHistory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新SmsCouponHistory
func (d *dao) UpdateSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory, id uint64) (update model.SmsCouponHistory, err error) {
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

// 删除SmsCouponHistory
func (d *dao) DeleteSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SmsCouponHistory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSmsCouponHistory(ctx context.Context, e *model.SmsCouponHistory, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SmsCouponHistory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
