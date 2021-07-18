package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsSubjectCategory 创建CmsSubjectCategory
func (d *dao) CreateCmsSubjectCategory(ctx context.Context, e *model.CmsSubjectCategory) (model.CmsSubjectCategory, error) {
	var doc model.CmsSubjectCategory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsSubjectCategory
func (d *dao) GetCmsSubjectCategory(ctx context.Context, e *model.CmsSubjectCategory) (model.CmsSubjectCategory, error) {
	var doc model.CmsSubjectCategory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsSubjectCategory带分页
func (d *dao) GetCmsSubjectCategoryPage(ctx context.Context, e *model.CmsSubjectCategory, pageSize int, pageIndex int) (docs []model.CmsSubjectCategory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsSubjectCategory
func (d *dao) UpdateCmsSubjectCategory(ctx context.Context, e *model.CmsSubjectCategory, id uint64) (update model.CmsSubjectCategory, err error) {
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

// 删除CmsSubjectCategory
func (d *dao) DeleteCmsSubjectCategory(ctx context.Context, e *model.CmsSubjectCategory, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsSubjectCategory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsSubjectCategory(ctx context.Context, e *model.CmsSubjectCategory, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsSubjectCategory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
