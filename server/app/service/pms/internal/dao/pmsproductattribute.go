package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductAttribute 创建PmsProductAttribute
func (d *dao) CreatePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute) (model.PmsProductAttribute, error) {
	var doc model.PmsProductAttribute
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductAttribute
func (d *dao) GetPmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute) (model.PmsProductAttribute, error) {
	var doc model.PmsProductAttribute
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductAttribute带分页
func (d *dao) GetPmsProductAttributePage(ctx context.Context, e *model.PmsProductAttribute, pageSize int, pageIndex int) (docs []model.PmsProductAttribute, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取PmsProductAttribute列表
func (d *dao) GetPmsProductAttributeList(ctx context.Context, e *model.PmsProductAttribute) (docs []model.PmsProductAttribute, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductAttribute
func (d *dao) UpdatePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute, id uint64) (update model.PmsProductAttribute, err error) {
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

// 删除PmsProductAttribute
func (d *dao) DeletePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductAttribute{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductAttribute(ctx context.Context, e *model.PmsProductAttribute, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductAttribute{}).Error; err != nil {
		return
	}
	Result = true
	return
}
