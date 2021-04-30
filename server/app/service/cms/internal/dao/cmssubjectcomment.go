package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsSubjectComment 创建CmsSubjectComment
func (d *dao) CreateCmsSubjectComment(ctx context.Context, e *model.CmsSubjectComment) (model.CmsSubjectComment, error) {
	var doc model.CmsSubjectComment
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsSubjectComment
func (d *dao) GetCmsSubjectComment(ctx context.Context, e *model.CmsSubjectComment) (model.CmsSubjectComment, error) {
	var doc model.CmsSubjectComment
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsSubjectComment带分页
func (d *dao) GetCmsSubjectCommentPage(ctx context.Context, e *model.CmsSubjectComment, pageSize int, pageIndex int) (docs []model.CmsSubjectComment, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsSubjectComment
func (d *dao) UpdateCmsSubjectComment(ctx context.Context, e *model.CmsSubjectComment, id uint64) (update model.CmsSubjectComment, err error) {
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

// 删除CmsSubjectComment
func (d *dao) DeleteCmsSubjectComment(ctx context.Context, e *model.CmsSubjectComment, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsSubjectComment{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsSubjectComment(ctx context.Context, e *model.CmsSubjectComment, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsSubjectComment{}).Error; err != nil {
		return
	}
	Result = true
	return
}
