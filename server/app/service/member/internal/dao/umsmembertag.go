package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberTag 创建UmsMemberTag
func (d *dao) CreateUmsMemberTag(ctx context.Context, e *model.UmsMemberTag) (model.UmsMemberTag, error) {
	var doc model.UmsMemberTag
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberTag
func (d *dao) GetUmsMemberTag(ctx context.Context, e *model.UmsMemberTag) (model.UmsMemberTag, error) {
	var doc model.UmsMemberTag
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberTag带分页
func (d *dao) GetUmsMemberTagPage(ctx context.Context, e *model.UmsMemberTag, pageSize int, pageIndex int) (docs []model.UmsMemberTag, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberTag
func (d *dao) UpdateUmsMemberTag(ctx context.Context, e *model.UmsMemberTag, id int) (update model.UmsMemberTag, err error) {
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

// 删除UmsMemberTag
func (d *dao) DeleteUmsMemberTag(ctx context.Context, e *model.UmsMemberTag, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberTag{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberTag(ctx context.Context, e *model.UmsMemberTag, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberTag{}).Error; err != nil {
		return
	}
	Result = true
	return
}
