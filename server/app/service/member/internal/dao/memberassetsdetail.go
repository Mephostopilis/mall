package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateMemberAssetsDetail 创建MemberAssetsDetail
func (d *dao) CreateMemberAssetsDetail(ctx context.Context, e *model.MemberAssetsDetail) (model.MemberAssetsDetail, error) {
	var doc model.MemberAssetsDetail
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取MemberAssetsDetail
func (d *dao) GetMemberAssetsDetail(ctx context.Context, e *model.MemberAssetsDetail) (model.MemberAssetsDetail, error) {
	var doc model.MemberAssetsDetail
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取MemberAssetsDetail带分页
func (d *dao) GetMemberAssetsDetailPage(ctx context.Context, e *model.MemberAssetsDetail, pageSize int, pageIndex int) (docs []model.MemberAssetsDetail, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取MemberAssetsDetail列表
func (d *dao) GetMemberAssetsDetailList(ctx context.Context, e *model.MemberAssetsDetail) (docs []model.MemberAssetsDetail, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新MemberAssetsDetail
func (d *dao) UpdateMemberAssetsDetail(ctx context.Context, e *model.MemberAssetsDetail, id uint64) (update model.MemberAssetsDetail, err error) {
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

// 删除MemberAssetsDetail
func (d *dao) DeleteMemberAssetsDetail(ctx context.Context, e *model.MemberAssetsDetail, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.MemberAssetsDetail{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteMemberAssetsDetail(ctx context.Context, e *model.MemberAssetsDetail, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.MemberAssetsDetail{}).Error; err != nil {
		return
	}
	Result = true
	return
}
