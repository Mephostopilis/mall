package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsCommentReplay 创建PmsCommentReplay
func (d *dao) CreatePmsCommentReplay(ctx context.Context, e *model.PmsCommentReplay) (model.PmsCommentReplay, error) {
	var doc model.PmsCommentReplay
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsCommentReplay
func (d *dao) GetPmsCommentReplay(ctx context.Context, e *model.PmsCommentReplay) (model.PmsCommentReplay, error) {
	var doc model.PmsCommentReplay
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsCommentReplay带分页
func (d *dao) GetPmsCommentReplayPage(ctx context.Context, e *model.PmsCommentReplay, pageSize int, pageIndex int) (docs []model.PmsCommentReplay, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsCommentReplay
func (d *dao) UpdatePmsCommentReplay(ctx context.Context, e *model.PmsCommentReplay, id int) (update model.PmsCommentReplay, err error) {
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

// 删除PmsCommentReplay
func (d *dao) DeletePmsCommentReplay(ctx context.Context, e *model.PmsCommentReplay, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsCommentReplay{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsCommentReplay(ctx context.Context, e *model.PmsCommentReplay, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsCommentReplay{}).Error; err != nil {
		return
	}
	Result = true
	return
}
