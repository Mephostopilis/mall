package dao

import (
	"context"
	"edu/service/member/internal/model"
)

//CreateUmsMemberStatisticsInfo 创建UmsMemberStatisticsInfo
func (d *dao) CreateUmsMemberStatisticsInfo(ctx context.Context, e *model.UmsMemberStatisticsInfo) (model.UmsMemberStatisticsInfo, error) {
	var doc model.UmsMemberStatisticsInfo
	table := d.orm.WithContext(ctx).Table(e.TableName())
	if err := table.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UmsMemberStatisticsInfo
func (d *dao) GetUmsMemberStatisticsInfo(ctx context.Context, e *model.UmsMemberStatisticsInfo) (model.UmsMemberStatisticsInfo, error) {
	var doc model.UmsMemberStatisticsInfo
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UmsMemberStatisticsInfo带分页
func (d *dao) GetUmsMemberStatisticsInfoPage(ctx context.Context, e *model.UmsMemberStatisticsInfo, pageSize int, pageIndex int) (docs []model.UmsMemberStatisticsInfo, total int64, err error) {

	table := d.orm.Table(e.TableName())

	if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
		return
	}

	return
}

// 更新UmsMemberStatisticsInfo
func (d *dao) UpdateUmsMemberStatisticsInfo(ctx context.Context, e *model.UmsMemberStatisticsInfo, id int) (update model.UmsMemberStatisticsInfo, err error) {
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

// 删除UmsMemberStatisticsInfo
func (d *dao) DeleteUmsMemberStatisticsInfo(ctx context.Context, e *model.UmsMemberStatisticsInfo, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UmsMemberStatisticsInfo{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUmsMemberStatisticsInfo(ctx context.Context, e *model.UmsMemberStatisticsInfo, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UmsMemberStatisticsInfo{}).Error; err != nil {
		return
	}
	Result = true
	return
}
