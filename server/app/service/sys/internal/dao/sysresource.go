package dao

import (
	"context"
	"edu/service/sys/internal/model"
)

//CreateSysResource 创建SysResource
func (d *dao) CreateSysResource(ctx context.Context, e *model.SysResource) (model.SysResource, error) {
	var doc model.SysResource
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SysResource
func (d *dao) GetSysResource(ctx context.Context, e *model.SysResource) (model.SysResource, error) {
	var doc model.SysResource
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SysResource带分页
func (d *dao) GetSysResourcePage(ctx context.Context, e *model.SysResource, pageSize int, pageIndex int) (docs []model.SysResource, total int64, err error) {
	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}
	return
}

// 更新SysResource
func (d *dao) UpdateSysResource(ctx context.Context, e *model.SysResource, id int) (update model.SysResource, err error) {
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

// 删除SysResource
func (d *dao) DeleteSysResource(ctx context.Context, e *model.SysResource, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SysResource{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSysResource(ctx context.Context, e *model.SysResource, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SysResource{}).Error; err != nil {
		return
	}
	Result = true
	return
}
