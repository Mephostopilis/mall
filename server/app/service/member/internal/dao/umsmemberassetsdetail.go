package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberAssetsDetail 创建UmsMemberAssetsDetail
func (d *dao) CreateUmsMemberAssetsDetail(ctx context.Context, e *model.UmsMemberAssetsDetail) (model.UmsMemberAssetsDetail, error) {
	var doc model.UmsMemberAssetsDetail
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberAssetsDetail
func (d *dao) GetUmsMemberAssetsDetail(ctx context.Context, e *model.UmsMemberAssetsDetail) (model.UmsMemberAssetsDetail, error) {
	var doc model.UmsMemberAssetsDetail
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberAssetsDetail带分页
func (d *dao) GetUmsMemberAssetsDetailPage(ctx context.Context, e *model.UmsMemberAssetsDetail, pageSize int, pageIndex int) (docs []model.UmsMemberAssetsDetail, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberAssetsDetail
func (d *dao) UpdateUmsMemberAssetsDetail(ctx context.Context, e *model.UmsMemberAssetsDetail, id int) (update model.UmsMemberAssetsDetail, err error) {
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

// 删除UmsMemberAssetsDetail
func (d *dao) DeleteUmsMemberAssetsDetail(ctx context.Context, e *model.UmsMemberAssetsDetail, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberAssetsDetail{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberAssetsDetail(ctx context.Context, e *model.UmsMemberAssetsDetail, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberAssetsDetail{}).Error; err != nil {
		return
	}
	Result = true
	return
}
