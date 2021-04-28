package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductAttributeCategory 创建PmsProductAttributeCategory
func (d *dao) CreatePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory) (model.PmsProductAttributeCategory, error) {
	var doc model.PmsProductAttributeCategory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductAttributeCategory
func (d *dao) GetPmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory) (model.PmsProductAttributeCategory, error) {
	var doc model.PmsProductAttributeCategory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductAttributeCategory带分页
func (d *dao) GetPmsProductAttributeCategoryPage(ctx context.Context, e *model.PmsProductAttributeCategory, pageSize int, pageIndex int) (docs []model.PmsProductAttributeCategory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductAttributeCategory
func (d *dao) UpdatePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory, id uint64) (update model.PmsProductAttributeCategory, err error) {
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

// 删除PmsProductAttributeCategory
func (d *dao) DeletePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductAttributeCategory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductAttributeCategory(ctx context.Context, e *model.PmsProductAttributeCategory, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductAttributeCategory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
