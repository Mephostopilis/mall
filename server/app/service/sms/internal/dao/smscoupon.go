package dao

import (
	"context"
	"edu/service/sms/internal/model"
)

//CreateSmsCoupon 创建SmsCoupon
func (d *dao) CreateSmsCoupon(ctx context.Context, e *model.SmsCoupon) (model.SmsCoupon, error) {
	var doc model.SmsCoupon
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SmsCoupon
func (d *dao) GetSmsCoupon(ctx context.Context, e *model.SmsCoupon) (model.SmsCoupon, error) {
	var doc model.SmsCoupon
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SmsCoupon带分页
func (d *dao) GetSmsCouponPage(ctx context.Context, e *model.SmsCoupon, pageSize int, pageIndex int) (docs []model.SmsCoupon, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新SmsCoupon
func (d *dao) UpdateSmsCoupon(ctx context.Context, e *model.SmsCoupon, id uint64) (update model.SmsCoupon, err error) {
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

// 删除SmsCoupon
func (d *dao) DeleteSmsCoupon(ctx context.Context, e *model.SmsCoupon, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SmsCoupon{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSmsCoupon(ctx context.Context, e *model.SmsCoupon, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SmsCoupon{}).Error; err != nil {
		return
	}
	Result = true
	return
}
