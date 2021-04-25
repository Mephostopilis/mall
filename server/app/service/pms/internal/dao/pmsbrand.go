package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsBrand 创建PmsBrand
func (d *dao) CreatePmsBrand(ctx context.Context, e *model.PmsBrand) (model.PmsBrand, error) {
	var doc model.PmsBrand
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsBrand
func (d *dao) GetPmsBrand(ctx context.Context, e *model.PmsBrand) (model.PmsBrand, error) {
	var doc model.PmsBrand
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsBrand带分页
func (d *dao) GetPmsBrandPage(ctx context.Context, e *model.PmsBrand, pageSize int, pageIndex int) (docs []model.PmsBrand, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsBrand
func (d *dao) UpdatePmsBrand(ctx context.Context, e *model.PmsBrand, id int) (update model.PmsBrand, err error) {
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

// 删除PmsBrand
func (d *dao) DeletePmsBrand(ctx context.Context, e *model.PmsBrand, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsBrand{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsBrand(ctx context.Context, e *model.PmsBrand, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsBrand{}).Error; err != nil {
		return
	}
	Result = true
	return
}
