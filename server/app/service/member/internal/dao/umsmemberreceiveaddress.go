package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberReceiveAddress 创建UmsMemberReceiveAddress
func (d *dao) CreateUmsMemberReceiveAddress(ctx context.Context, e *model.UmsMemberReceiveAddress) (model.UmsMemberReceiveAddress, error) {
	var doc model.UmsMemberReceiveAddress
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberReceiveAddress
func (d *dao) GetUmsMemberReceiveAddress(ctx context.Context, e *model.UmsMemberReceiveAddress) (model.UmsMemberReceiveAddress, error) {
	var doc model.UmsMemberReceiveAddress
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberReceiveAddress带分页
func (d *dao) GetUmsMemberReceiveAddressPage(ctx context.Context, e *model.UmsMemberReceiveAddress, pageSize int, pageIndex int) (docs []model.UmsMemberReceiveAddress, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberReceiveAddress
func (d *dao) UpdateUmsMemberReceiveAddress(ctx context.Context, e *model.UmsMemberReceiveAddress, id int) (update model.UmsMemberReceiveAddress, err error) {
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

// 删除UmsMemberReceiveAddress
func (d *dao) DeleteUmsMemberReceiveAddress(ctx context.Context, e *model.UmsMemberReceiveAddress, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberReceiveAddress{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberReceiveAddress(ctx context.Context, e *model.UmsMemberReceiveAddress, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberReceiveAddress{}).Error; err != nil {
		return
	}
	Result = true
	return
}
