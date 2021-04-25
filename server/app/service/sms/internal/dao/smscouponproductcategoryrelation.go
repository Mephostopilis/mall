package dao

import (
	"context"
	"edu/service/sms/internal/model"
)

//CreateSmsCouponProductCategoryRelation 创建SmsCouponProductCategoryRelation
func (d *dao) CreateSmsCouponProductCategoryRelation(ctx context.Context, e *model.SmsCouponProductCategoryRelation) (model.SmsCouponProductCategoryRelation, error) {
	var doc model.SmsCouponProductCategoryRelation
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SmsCouponProductCategoryRelation
func (d *dao) GetSmsCouponProductCategoryRelation(ctx context.Context, e *model.SmsCouponProductCategoryRelation) (model.SmsCouponProductCategoryRelation, error) {
	var doc model.SmsCouponProductCategoryRelation
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SmsCouponProductCategoryRelation带分页
func (d *dao) GetSmsCouponProductCategoryRelationPage(ctx context.Context, e *model.SmsCouponProductCategoryRelation, pageSize int, pageIndex int) (docs []model.SmsCouponProductCategoryRelation, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新SmsCouponProductCategoryRelation
func (d *dao) UpdateSmsCouponProductCategoryRelation(ctx context.Context, e *model.SmsCouponProductCategoryRelation, id int) (update model.SmsCouponProductCategoryRelation, err error) {
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

// 删除SmsCouponProductCategoryRelation
func (d *dao) DeleteSmsCouponProductCategoryRelation(ctx context.Context, e *model.SmsCouponProductCategoryRelation, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SmsCouponProductCategoryRelation{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSmsCouponProductCategoryRelation(ctx context.Context, e *model.SmsCouponProductCategoryRelation, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SmsCouponProductCategoryRelation{}).Error; err != nil {
		return
	}
	Result = true
	return
}
