package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsPrefrenceAreaProductRelation 创建CmsPrefrenceAreaProductRelation
func (d *dao) CreateCmsPrefrenceAreaProductRelation(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation) (model.CmsPrefrenceAreaProductRelation, error) {
	var doc model.CmsPrefrenceAreaProductRelation
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsPrefrenceAreaProductRelation
func (d *dao) GetCmsPrefrenceAreaProductRelation(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation) (model.CmsPrefrenceAreaProductRelation, error) {
	var doc model.CmsPrefrenceAreaProductRelation
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsPrefrenceAreaProductRelation带分页
func (d *dao) GetCmsPrefrenceAreaProductRelationPage(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation, pageSize int, pageIndex int) (docs []model.CmsPrefrenceAreaProductRelation, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsPrefrenceAreaProductRelation
func (d *dao) UpdateCmsPrefrenceAreaProductRelation(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation, id int) (update model.CmsPrefrenceAreaProductRelation, err error) {
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

// 删除CmsPrefrenceAreaProductRelation
func (d *dao) DeleteCmsPrefrenceAreaProductRelation(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsPrefrenceAreaProductRelation{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsPrefrenceAreaProductRelation(ctx context.Context, e *model.CmsPrefrenceAreaProductRelation, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsPrefrenceAreaProductRelation{}).Error; err != nil {
		return
	}
	Result = true
	return
}
