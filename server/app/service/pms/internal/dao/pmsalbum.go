package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsAlbum 创建PmsAlbum
func (d *dao) CreatePmsAlbum(ctx context.Context, e *model.PmsAlbum) (model.PmsAlbum, error) {
	var doc model.PmsAlbum
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsAlbum
func (d *dao) GetPmsAlbum(ctx context.Context, e *model.PmsAlbum) (model.PmsAlbum, error) {
	var doc model.PmsAlbum
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsAlbum带分页
func (d *dao) GetPmsAlbumPage(ctx context.Context, e *model.PmsAlbum, pageSize int, pageIndex int) (docs []model.PmsAlbum, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsAlbum
func (d *dao) UpdatePmsAlbum(ctx context.Context, e *model.PmsAlbum, id int) (update model.PmsAlbum, err error) {
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

// 删除PmsAlbum
func (d *dao) DeletePmsAlbum(ctx context.Context, e *model.PmsAlbum, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsAlbum{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsAlbum(ctx context.Context, e *model.PmsAlbum, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsAlbum{}).Error; err != nil {
		return
	}
	Result = true
	return
}
