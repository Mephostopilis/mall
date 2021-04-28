package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductCategory 创建PmsProductCategory
func (d *dao) CreatePmsProductCategory(ctx context.Context, e *model.PmsProductCategory) (model.PmsProductCategory, error) {
	var doc model.PmsProductCategory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductCategory
func (d *dao) GetPmsProductCategory(ctx context.Context, e *model.PmsProductCategory) (model.PmsProductCategory, error) {
	var doc model.PmsProductCategory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductCategory带分页
func (d *dao) GetPmsProductCategoryPage(ctx context.Context, e *model.PmsProductCategory, pageSize int, pageIndex int) (docs []model.PmsProductCategory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取PmsProductCategory带分页
func (d *dao) GetPmsProductCategoryList(ctx context.Context, e *model.PmsProductCategory) (docs []model.PmsProductCategory, err error) {
	table := d.orm.Table(e.TableName())
	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductCategory
func (d *dao) UpdatePmsProductCategory(ctx context.Context, e *model.PmsProductCategory, id uint64) (update model.PmsProductCategory, err error) {
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

// 删除PmsProductCategory
func (d *dao) DeletePmsProductCategory(ctx context.Context, e *model.PmsProductCategory, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductCategory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductCategory(ctx context.Context, e *model.PmsProductCategory, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductCategory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
