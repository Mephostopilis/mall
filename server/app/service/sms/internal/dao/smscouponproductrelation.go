package dao

import (
	"context"
	"edu/service/sms/internal/model"
)

//CreateSmsCouponProductRelation 创建SmsCouponProductRelation
func (d *dao) CreateSmsCouponProductRelation(ctx context.Context, e *model.SmsCouponProductRelation) (model.SmsCouponProductRelation, error) {
	var doc model.SmsCouponProductRelation
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SmsCouponProductRelation
func (d *dao) GetSmsCouponProductRelation(ctx context.Context, e *model.SmsCouponProductRelation) (model.SmsCouponProductRelation, error) {
	var doc model.SmsCouponProductRelation
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SmsCouponProductRelation带分页
func (d *dao) GetSmsCouponProductRelationPage(ctx context.Context, e *model.SmsCouponProductRelation, pageSize int, pageIndex int) (docs []model.SmsCouponProductRelation, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新SmsCouponProductRelation
func (d *dao) UpdateSmsCouponProductRelation(ctx context.Context, e *model.SmsCouponProductRelation, id int) (update model.SmsCouponProductRelation, err error) {
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

// 删除SmsCouponProductRelation
func (d *dao) DeleteSmsCouponProductRelation(ctx context.Context, e *model.SmsCouponProductRelation, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SmsCouponProductRelation{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSmsCouponProductRelation(ctx context.Context, e *model.SmsCouponProductRelation, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SmsCouponProductRelation{}).Error; err != nil {
		return
	}
	Result = true
	return
}
