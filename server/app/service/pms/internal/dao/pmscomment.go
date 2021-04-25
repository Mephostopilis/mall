package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsComment 创建PmsComment
func (d *dao) CreatePmsComment(ctx context.Context, e *model.PmsComment) (model.PmsComment, error) {
	var doc model.PmsComment
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsComment
func (d *dao) GetPmsComment(ctx context.Context, e *model.PmsComment) (model.PmsComment, error) {
	var doc model.PmsComment
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsComment带分页
func (d *dao) GetPmsCommentPage(ctx context.Context, e *model.PmsComment, pageSize int, pageIndex int) (docs []model.PmsComment, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsComment
func (d *dao) UpdatePmsComment(ctx context.Context, e *model.PmsComment, id int) (update model.PmsComment, err error) {
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

// 删除PmsComment
func (d *dao) DeletePmsComment(ctx context.Context, e *model.PmsComment, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsComment{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsComment(ctx context.Context, e *model.PmsComment, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsComment{}).Error; err != nil {
		return
	}
	Result = true
	return
}
