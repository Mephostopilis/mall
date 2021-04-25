package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsPrefrenceArea 创建CmsPrefrenceArea
func (d *dao) CreateCmsPrefrenceArea(ctx context.Context, e *model.CmsPrefrenceArea) (model.CmsPrefrenceArea, error) {
	var doc model.CmsPrefrenceArea
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsPrefrenceArea
func (d *dao) GetCmsPrefrenceArea(ctx context.Context, e *model.CmsPrefrenceArea) (model.CmsPrefrenceArea, error) {
	var doc model.CmsPrefrenceArea
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsPrefrenceArea带分页
func (d *dao) GetCmsPrefrenceAreaPage(ctx context.Context, e *model.CmsPrefrenceArea, pageSize int, pageIndex int) (docs []model.CmsPrefrenceArea, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsPrefrenceArea
func (d *dao) UpdateCmsPrefrenceArea(ctx context.Context, e *model.CmsPrefrenceArea, id int) (update model.CmsPrefrenceArea, err error) {
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

// 删除CmsPrefrenceArea
func (d *dao) DeleteCmsPrefrenceArea(ctx context.Context, e *model.CmsPrefrenceArea, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsPrefrenceArea{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsPrefrenceArea(ctx context.Context, e *model.CmsPrefrenceArea, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsPrefrenceArea{}).Error; err != nil {
		return
	}
	Result = true
	return
}
