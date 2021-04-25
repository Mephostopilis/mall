package dao

import (
	"edu/service/sso/internal/model"
)

func (d *dao) GetLoginLog(e *model.SysLoginLog) (model.SysLoginLog, error) {
	var doc model.SysLoginLog

	table := d.orm.Table(e.TableName())
	if e.Ipaddr != "" {
		table = table.Where("ipaddr = ?", e.Ipaddr)
	}
	if e.InfoId != 0 {
		table = table.Where("info_id = ?", e.InfoId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetLoginLogPage(e *model.SysLoginLog, pageSize int, pageIndex int) ([]model.SysLoginLog, int64, error) {
	var doc []model.SysLoginLog

	table := d.orm.Table(e.TableName())
	if e.Ipaddr != "" {
		table = table.Where("ipaddr = ?", e.Ipaddr)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}
	if e.Username != "" {
		table = table.Where("userName = ?", e.Username)
	}

	var count int64
	if err := table.Order("info_id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (d *dao) CreateLoginLog(e *model.SysLoginLog) (model.SysLoginLog, error) {
	var doc model.SysLoginLog
	e.CreateBy = "0"
	e.UpdateBy = "0"
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (d *dao) UpdateLoginLog(e *model.SysLoginLog, id int) (update model.SysLoginLog, err error) {

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

func (d *dao) BatchDeleteLoginLog(e *model.SysLoginLog, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("info_id in (?)", id).Delete(&model.SysLoginLog{}).Error; err != nil {
		return
	}
	Result = true
	return
}
