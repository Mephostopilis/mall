package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateMember 创建Member
func (d *dao) CreateMember(ctx context.Context, e *model.Member) (model.Member, error) {
	var doc model.Member
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Member
func (d *dao) GetMember(ctx context.Context, e *model.Member) (model.Member, error) {
	var doc model.Member
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Member带分页
func (d *dao) GetMemberPage(ctx context.Context, e *model.Member, pageSize int, pageIndex int) (docs []model.Member, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取Member列表
func (d *dao) GetMemberList(ctx context.Context, e *model.Member) (docs []model.Member, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新Member
func (d *dao) UpdateMember(ctx context.Context, e *model.Member, id uint64) (update model.Member, err error) {
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

// 删除Member
func (d *dao) DeleteMember(ctx context.Context, e *model.Member, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.Member{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteMember(ctx context.Context, e *model.Member, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.Member{}).Error; err != nil {
		return
	}
	Result = true
	return
}
