package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberAssets 创建UmsMemberAssets
func (d *dao) CreateUmsMemberAssets(ctx context.Context, e *model.UmsMemberAssets) (model.UmsMemberAssets, error) {
	var doc model.UmsMemberAssets
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberAssets
func (d *dao) GetUmsMemberAssets(ctx context.Context, e *model.UmsMemberAssets) (model.UmsMemberAssets, error) {
	var doc model.UmsMemberAssets
	table := d.orm.Table(e.TableName())

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberAssets带分页
func (d *dao) GetUmsMemberAssetsPage(ctx context.Context, e *model.UmsMemberAssets, pageSize int, pageIndex int) (docs []model.UmsMemberAssets, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberAssets
func (d *dao) UpdateUmsMemberAssets(ctx context.Context, e *model.UmsMemberAssets, id int) (update model.UmsMemberAssets, err error) {
	if err = d.orm.Table(e.TableName()).Where("username = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除UmsMemberAssets
func (d *dao) DeleteUmsMemberAssets(ctx context.Context, e *model.UmsMemberAssets, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("username = ?", id).Delete(&model.UmsMemberAssets{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberAssets(ctx context.Context, e *model.UmsMemberAssets, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("username in (?)", id).Delete(&model.UmsMemberAssets{}).Error; err != nil {
		return
	}
	Result = true
	return
}
