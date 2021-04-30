package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateMemberAssets 创建MemberAssets
func (d *dao) CreateMemberAssets(ctx context.Context, e *model.MemberAssets) (model.MemberAssets, error) {
	var doc model.MemberAssets
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取MemberAssets
func (d *dao) GetMemberAssets(ctx context.Context, e *model.MemberAssets) (model.MemberAssets, error) {
	var doc model.MemberAssets
	table := d.orm.Table(e.TableName())

	if e.MemberId != 0 {
		table = table.Where("member_id = ?", e.MemberId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取MemberAssets带分页
func (d *dao) GetMemberAssetsPage(ctx context.Context, e *model.MemberAssets, pageSize int, pageIndex int) (docs []model.MemberAssets, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 获取MemberAssets列表
func (d *dao) GetMemberAssetsList(ctx context.Context, e *model.MemberAssets) (docs []model.MemberAssets, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新MemberAssets
func (d *dao) UpdateMemberAssets(ctx context.Context, e *model.MemberAssets, id uint64) (update model.MemberAssets, err error) {
	if err = d.orm.Table(e.TableName()).Where("member_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除MemberAssets
func (d *dao) DeleteMemberAssets(ctx context.Context, e *model.MemberAssets, id uint64) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("member_id = ?", id).Delete(&model.MemberAssets{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteMemberAssets(ctx context.Context, e *model.MemberAssets, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("member_id in (?)", id).Delete(&model.MemberAssets{}).Error; err != nil {
		return
	}
	Result = true
	return
}
