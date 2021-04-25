package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsHelpCategory 创建CmsHelpCategory
func (d *dao) CreateCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory) (model.CmsHelpCategory, error) {
	var doc model.CmsHelpCategory
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsHelpCategory
func (d *dao) GetCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory) (model.CmsHelpCategory, error) {
	var doc model.CmsHelpCategory
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsHelpCategory带分页
func (d *dao) GetCmsHelpCategoryPage(ctx context.Context, e *model.CmsHelpCategory, pageSize int, pageIndex int) (docs []model.CmsHelpCategory, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsHelpCategory
func (d *dao) UpdateCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory, id int) (update model.CmsHelpCategory, err error) {
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

// 删除CmsHelpCategory
func (d *dao) DeleteCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsHelpCategory{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsHelpCategory(ctx context.Context, e *model.CmsHelpCategory, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsHelpCategory{}).Error; err != nil {
		return
	}
	Result = true
	return
}
