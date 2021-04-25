package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsHelp 创建CmsHelp
func (d *dao) CreateCmsHelp(ctx context.Context, e *model.CmsHelp) (model.CmsHelp, error) {
	var doc model.CmsHelp
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsHelp
func (d *dao) GetCmsHelp(ctx context.Context, e *model.CmsHelp) (model.CmsHelp, error) {
	var doc model.CmsHelp
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsHelp带分页
func (d *dao) GetCmsHelpPage(ctx context.Context, e *model.CmsHelp, pageSize int, pageIndex int) (docs []model.CmsHelp, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsHelp
func (d *dao) UpdateCmsHelp(ctx context.Context, e *model.CmsHelp, id int) (update model.CmsHelp, err error) {
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

// 删除CmsHelp
func (d *dao) DeleteCmsHelp(ctx context.Context, e *model.CmsHelp, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsHelp{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsHelp(ctx context.Context, e *model.CmsHelp, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsHelp{}).Error; err != nil {
		return
	}
	Result = true
	return
}
