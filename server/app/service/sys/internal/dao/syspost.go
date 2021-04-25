package dao

import (
	"edu/service/sys/internal/model"
)

func (d *dao) CreatePost(e *model.SysPost) (model.SysPost, error) {
	var doc model.SysPost
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (d *dao) GetPost(e *model.SysPost) (model.SysPost, error) {
	var doc model.SysPost

	table := d.orm.Table(e.TableName())
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetPostList(e *model.SysPost) ([]model.SysPost, error) {
	var doc []model.SysPost

	table := d.orm.Table(e.TableName())
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetPostPage(e *model.SysPost, pageSize int, pageIndex int) ([]model.SysPost, int64, error) {
	var doc []model.SysPost
	table := d.orm.Table(e.TableName())
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	var count int64
	if err := table.Order("sort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (d *dao) UpdatePost(e *model.SysPost, id int) (update model.SysPost, err error) {
	if err = d.orm.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) DeletePost(e *model.SysPost, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("post_id = ?", id).Delete(&model.SysPost{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) BatchDeletePost(e *model.SysPost, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("post_id in (?)", id).Delete(&model.SysPost{}).Error; err != nil {
		return
	}
	Result = true
	return
}
