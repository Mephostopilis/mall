package dao

import (
	"context"
	"edu/service/cms/internal/model"
)

//CreateCmsSubject 创建CmsSubject
func (d *dao) CreateCmsSubject(ctx context.Context, e *model.CmsSubject) (model.CmsSubject, error) {
	var doc model.CmsSubject
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CmsSubject
func (d *dao) GetCmsSubject(ctx context.Context, e *model.CmsSubject) (model.CmsSubject, error) {
	var doc model.CmsSubject
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CmsSubject带分页
func (d *dao) GetCmsSubjectPage(ctx context.Context, e *model.CmsSubject, pageSize int, pageIndex int) (docs []model.CmsSubject, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新CmsSubject
func (d *dao) UpdateCmsSubject(ctx context.Context, e *model.CmsSubject, id uint64) (update model.CmsSubject, err error) {
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

// 删除CmsSubject
func (d *dao) DeleteCmsSubject(ctx context.Context, e *model.CmsSubject, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.CmsSubject{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteCmsSubject(ctx context.Context, e *model.CmsSubject, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.CmsSubject{}).Error; err != nil {
		return
	}
	Result = true
	return
}
