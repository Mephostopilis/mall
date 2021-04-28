package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProduct 创建PmsProduct
func (d *dao) CreatePmsProduct(ctx context.Context, e *model.PmsProduct) (model.PmsProduct, error) {
	var doc model.PmsProduct
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProduct
func (d *dao) GetPmsProduct(ctx context.Context, e *model.PmsProduct) (model.PmsProduct, error) {
	var doc model.PmsProduct
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProduct带分页
func (d *dao) GetPmsProductPage(ctx context.Context, e *model.PmsProduct, pageSize int, pageIndex int) (docs []model.PmsProduct, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProduct
func (d *dao) UpdatePmsProduct(ctx context.Context, e *model.PmsProduct, id uint64) (update model.PmsProduct, err error) {
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

// 删除PmsProduct
func (d *dao) DeletePmsProduct(ctx context.Context, e *model.PmsProduct, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProduct{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProduct(ctx context.Context, e *model.PmsProduct, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProduct{}).Error; err != nil {
		return
	}
	Result = true
	return
}
