package dao

import (
	"context"
	"edu/service/pms/internal/model"
)

//CreatePmsAlbumPic 创建PmsAlbumPic
func (d *dao) CreatePmsAlbumPic(ctx context.Context, e *model.PmsAlbumPic) (model.PmsAlbumPic, error) {
	var doc model.PmsAlbumPic
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取PmsAlbumPic
func (d *dao) GetPmsAlbumPic(ctx context.Context, e *model.PmsAlbumPic) (model.PmsAlbumPic, error) {
	var doc model.PmsAlbumPic
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取PmsAlbumPic带分页
func (d *dao) GetPmsAlbumPicPage(ctx context.Context, e *model.PmsAlbumPic, pageSize int, pageIndex int) (docs []model.PmsAlbumPic, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新PmsAlbumPic
func (d *dao) UpdatePmsAlbumPic(ctx context.Context, e *model.PmsAlbumPic, id int) (update model.PmsAlbumPic, err error) {
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

// 删除PmsAlbumPic
func (d *dao) DeletePmsAlbumPic(ctx context.Context, e *model.PmsAlbumPic, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.PmsAlbumPic{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeletePmsAlbumPic(ctx context.Context, e *model.PmsAlbumPic, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.PmsAlbumPic{}).Error; err != nil {
		return
	}
	Result = true
	return
}
