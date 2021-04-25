package dao

import (
	"edu/service/sys/internal/model"
)

func (d *dao) GetSysOperLog(e *model.SysOperLog) (model.SysOperLog, error) {
	var doc model.SysOperLog

	table := d.orm.Table(e.TableName())
	if e.OperIp != "" {
		table = table.Where("oper_ip = ?", e.OperIp)
	}
	if e.OperId != 0 {
		table = table.Where("oper_id = ?", e.OperId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetSysOperLogPage(e *model.SysOperLog, pageSize int, pageIndex int) (doc []model.SysOperLog, count int64, err error) {
	table := d.orm.Table(e.TableName()).Select("*")
	if e.OperIp != "" {
		table = table.Where("oper_ip = ?", e.OperIp)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}
	if e.OperName != "" {
		table = table.Where("oper_name = ?", e.OperName)
	}
	if e.BusinessType != "" {
		table = table.Where("business_type = ?", e.BusinessType)
	}
	if err := table.Order("oper_id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (d *dao) CreateSysOperLog(e *model.SysOperLog) (model.SysOperLog, error) {
	var doc model.SysOperLog
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

func (d *dao) UpdateSysOperLog(e *model.SysOperLog, id int) (update model.SysOperLog, err error) {
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

func (d *dao) BatchDeleteSysOperLog(e *model.SysOperLog, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where(" oper_id in (?)", id).Delete(&model.SysOperLog{}).Error; err != nil {
		return
	}
	Result = true
	return
}
