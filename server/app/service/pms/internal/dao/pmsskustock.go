package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsSkuStock 创建PmsSkuStock
func (d *dao) CreatePmsSkuStock(ctx context.Context, e *model.PmsSkuStock) (model.PmsSkuStock, error) {
	var doc model.PmsSkuStock
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsSkuStock
func (d *dao) GetPmsSkuStock(ctx context.Context, e *model.PmsSkuStock) (model.PmsSkuStock, error) {
	var doc model.PmsSkuStock
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsSkuStock带分页
func (d *dao) GetPmsSkuStockPage(ctx context.Context, e *model.PmsSkuStock, pageSize int, pageIndex int) (docs []model.PmsSkuStock, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsSkuStock
func (d *dao) UpdatePmsSkuStock(ctx context.Context, e *model.PmsSkuStock, id int) (update model.PmsSkuStock, err error) {
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

// 删除PmsSkuStock
func (d *dao) DeletePmsSkuStock(ctx context.Context, e *model.PmsSkuStock, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsSkuStock{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsSkuStock(ctx context.Context, e *model.PmsSkuStock, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsSkuStock{}).Error; err != nil {
		return
	}
	Result = true
	return
}
