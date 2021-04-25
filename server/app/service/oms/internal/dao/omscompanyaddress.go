package dao

import (
	"context"
	"edu/service/oms/internal/model"
)

//CreateOmsCompanyAddress 创建OmsCompanyAddress
func (d *dao) CreateOmsCompanyAddress(ctx context.Context, e *model.OmsCompanyAddress) (model.OmsCompanyAddress, error) {
	var doc model.OmsCompanyAddress
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取OmsCompanyAddress
func (d *dao) GetOmsCompanyAddress(ctx context.Context, e *model.OmsCompanyAddress) (model.OmsCompanyAddress, error) {
	var doc model.OmsCompanyAddress
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取OmsCompanyAddress带分页
func (d *dao) GetOmsCompanyAddressPage(ctx context.Context, e *model.OmsCompanyAddress, pageSize int, pageIndex int) (docs []model.OmsCompanyAddress, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新OmsCompanyAddress
func (d *dao) UpdateOmsCompanyAddress(ctx context.Context, e *model.OmsCompanyAddress, id int) (update model.OmsCompanyAddress, err error) {
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

// 删除OmsCompanyAddress
func (d *dao) DeleteOmsCompanyAddress(ctx context.Context, e *model.OmsCompanyAddress, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.OmsCompanyAddress{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteOmsCompanyAddress(ctx context.Context, e *model.OmsCompanyAddress, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.OmsCompanyAddress{}).Error; err != nil {
		return
	}
	Result = true
	return
}
