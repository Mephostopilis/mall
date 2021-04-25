package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberLevel 创建UmsMemberLevel
func (d *dao) CreateUmsMemberLevel(ctx context.Context, e *model.UmsMemberLevel) (model.UmsMemberLevel, error) {
	var doc model.UmsMemberLevel
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberLevel
func (d *dao) GetUmsMemberLevel(ctx context.Context, e *model.UmsMemberLevel) (model.UmsMemberLevel, error) {
	var doc model.UmsMemberLevel
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberLevel带分页
func (d *dao) GetUmsMemberLevelPage(ctx context.Context, e *model.UmsMemberLevel, pageSize int, pageIndex int) (docs []model.UmsMemberLevel, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberLevel
func (d *dao) UpdateUmsMemberLevel(ctx context.Context, e *model.UmsMemberLevel, id int) (update model.UmsMemberLevel, err error) {
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

// 删除UmsMemberLevel
func (d *dao) DeleteUmsMemberLevel(ctx context.Context, e *model.UmsMemberLevel, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberLevel{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberLevel(ctx context.Context, e *model.UmsMemberLevel, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberLevel{}).Error; err != nil {
		return
	}
	Result = true
	return
}
