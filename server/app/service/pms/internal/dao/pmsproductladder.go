package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsProductLadder 创建PmsProductLadder
func (d *dao) CreatePmsProductLadder(ctx context.Context, e *model.PmsProductLadder) (model.PmsProductLadder, error) {
	var doc model.PmsProductLadder
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsProductLadder
func (d *dao) GetPmsProductLadder(ctx context.Context, e *model.PmsProductLadder) (model.PmsProductLadder, error) {
	var doc model.PmsProductLadder
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsProductLadder带分页
func (d *dao) GetPmsProductLadderPage(ctx context.Context, e *model.PmsProductLadder, pageSize int, pageIndex int) (docs []model.PmsProductLadder, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsProductLadder
func (d *dao) UpdatePmsProductLadder(ctx context.Context, e *model.PmsProductLadder, id int) (update model.PmsProductLadder, err error) {
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

// 删除PmsProductLadder
func (d *dao) DeletePmsProductLadder(ctx context.Context, e *model.PmsProductLadder, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsProductLadder{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsProductLadder(ctx context.Context, e *model.PmsProductLadder, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsProductLadder{}).Error; err != nil {
		return
	}
	Result = true
	return
}
