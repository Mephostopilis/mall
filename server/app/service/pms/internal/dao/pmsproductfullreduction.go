package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductFullReduction 创建PmsProductFullReduction
func (d *dao) CreatePmsProductFullReduction(ctx context.Context, e *model.PmsProductFullReduction) (model.PmsProductFullReduction, error) {
	var doc model.PmsProductFullReduction
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductFullReduction
func (d *dao) GetPmsProductFullReduction(ctx context.Context, e *model.PmsProductFullReduction) (model.PmsProductFullReduction, error) {
	var doc model.PmsProductFullReduction
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductFullReduction带分页
func (d *dao) GetPmsProductFullReductionPage(ctx context.Context, e *model.PmsProductFullReduction, pageSize int, pageIndex int) (docs []model.PmsProductFullReduction, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductFullReduction
func (d *dao) UpdatePmsProductFullReduction(ctx context.Context, e *model.PmsProductFullReduction, id int) (update model.PmsProductFullReduction, err error) {
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

// 删除PmsProductFullReduction
func (d *dao) DeletePmsProductFullReduction(ctx context.Context, e *model.PmsProductFullReduction, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductFullReduction{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductFullReduction(ctx context.Context, e *model.PmsProductFullReduction, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductFullReduction{}).Error; err != nil {
		return
	}
	Result = true
	return
}
