package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductCategoryAttributeRelation 创建PmsProductCategoryAttributeRelation
func (d *dao) CreatePmsProductCategoryAttributeRelation(ctx context.Context, e *model.PmsProductCategoryAttributeRelation) (model.PmsProductCategoryAttributeRelation, error) {
	var doc model.PmsProductCategoryAttributeRelation
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductCategoryAttributeRelation
func (d *dao) GetPmsProductCategoryAttributeRelation(ctx context.Context, e *model.PmsProductCategoryAttributeRelation) (model.PmsProductCategoryAttributeRelation, error) {
	var doc model.PmsProductCategoryAttributeRelation
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductCategoryAttributeRelation带分页
func (d *dao) GetPmsProductCategoryAttributeRelationPage(ctx context.Context, e *model.PmsProductCategoryAttributeRelation, pageSize int, pageIndex int) (docs []model.PmsProductCategoryAttributeRelation, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductCategoryAttributeRelation
func (d *dao) UpdatePmsProductCategoryAttributeRelation(ctx context.Context, e *model.PmsProductCategoryAttributeRelation, id int) (update model.PmsProductCategoryAttributeRelation, err error) {
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

// 删除PmsProductCategoryAttributeRelation
func (d *dao) DeletePmsProductCategoryAttributeRelation(ctx context.Context, e *model.PmsProductCategoryAttributeRelation, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductCategoryAttributeRelation{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductCategoryAttributeRelation(ctx context.Context, e *model.PmsProductCategoryAttributeRelation, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductCategoryAttributeRelation{}).Error; err != nil {
		return
	}
	Result = true
	return
}
