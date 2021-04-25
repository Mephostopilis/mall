package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberRuleSetting 创建UmsMemberRuleSetting
func (d *dao) CreateUmsMemberRuleSetting(ctx context.Context, e *model.UmsMemberRuleSetting) (model.UmsMemberRuleSetting, error) {
	var doc model.UmsMemberRuleSetting
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberRuleSetting
func (d *dao) GetUmsMemberRuleSetting(ctx context.Context, e *model.UmsMemberRuleSetting) (model.UmsMemberRuleSetting, error) {
	var doc model.UmsMemberRuleSetting
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberRuleSetting带分页
func (d *dao) GetUmsMemberRuleSettingPage(ctx context.Context, e *model.UmsMemberRuleSetting, pageSize int, pageIndex int) (docs []model.UmsMemberRuleSetting, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberRuleSetting
func (d *dao) UpdateUmsMemberRuleSetting(ctx context.Context, e *model.UmsMemberRuleSetting, id int) (update model.UmsMemberRuleSetting, err error) {
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

// 删除UmsMemberRuleSetting
func (d *dao) DeleteUmsMemberRuleSetting(ctx context.Context, e *model.UmsMemberRuleSetting, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberRuleSetting{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberRuleSetting(ctx context.Context, e *model.UmsMemberRuleSetting, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberRuleSetting{}).Error; err != nil {
		return
	}
	Result = true
	return
}
