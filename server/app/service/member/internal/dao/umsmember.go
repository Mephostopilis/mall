package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMember 创建UmsMember
func (d *dao) CreateUmsMember(ctx context.Context, e *model.UmsMember) (model.UmsMember, error) {
	var doc model.UmsMember
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMember
func (d *dao) GetUmsMember(ctx context.Context, e *model.UmsMember) (model.UmsMember, error) {
	var doc model.UmsMember
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMember带分页
func (d *dao) GetUmsMemberPage(ctx context.Context, e *model.UmsMember, pageSize int, pageIndex int) (docs []model.UmsMember, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMember
func (d *dao) UpdateUmsMember(ctx context.Context, e *model.UmsMember, id int) (update model.UmsMember, err error) {
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

// 删除UmsMember
func (d *dao) DeleteUmsMember(ctx context.Context, e *model.UmsMember, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMember{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMember(ctx context.Context, e *model.UmsMember, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMember{}).Error; err != nil {
		return
	}
	Result = true
	return
}
