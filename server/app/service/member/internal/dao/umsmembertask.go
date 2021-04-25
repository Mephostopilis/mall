package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberTask 创建UmsMemberTask
func (d *dao) CreateUmsMemberTask(ctx context.Context, e *model.UmsMemberTask) (model.UmsMemberTask, error) {
	var doc model.UmsMemberTask
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberTask
func (d *dao) GetUmsMemberTask(ctx context.Context, e *model.UmsMemberTask) (model.UmsMemberTask, error) {
	var doc model.UmsMemberTask
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberTask带分页
func (d *dao) GetUmsMemberTaskPage(ctx context.Context, e *model.UmsMemberTask, pageSize int, pageIndex int) (docs []model.UmsMemberTask, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberTask
func (d *dao) UpdateUmsMemberTask(ctx context.Context, e *model.UmsMemberTask, id int) (update model.UmsMemberTask, err error) {
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

// 删除UmsMemberTask
func (d *dao) DeleteUmsMemberTask(ctx context.Context, e *model.UmsMemberTask, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberTask{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberTask(ctx context.Context, e *model.UmsMemberTask, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberTask{}).Error; err != nil {
		return
	}
	Result = true
	return
}
