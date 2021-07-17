package dao

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
)

// Config 创建
func (d *dao) CreateSysConfig(ctx context.Context, e *model.SysConfig) (doc model.SysConfig, err error) {
	db := d.orm.WithContext(ctx).Table(e.TableName())
	i := int64(0)
	db.Where("config_name=? or config_key = ?", e.ConfigName, e.ConfigKey).Count(&i)
	if i > 0 {
		return doc, errors.NotFound("not found", "参数名称或者参数键名已经存在！")
	}
	result := db.Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return
}

// 获取 Config
func (d *dao) GetSysConfig(e *model.SysConfig) (model.SysConfig, error) {
	var doc model.SysConfig

	table := d.orm.Table(e.TableName())
	if e.ConfigId != 0 {
		table = table.Where("config_id = ?", e.ConfigId)
	}

	if e.ConfigKey != "" {
		table = table.Where("config_key = ?", e.ConfigKey)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetSysConfigPage(e *model.SysConfig, pageSize int, pageIndex int) ([]model.SysConfig, int64, error) {
	var doc []model.SysConfig

	table := d.orm.Table(e.TableName())

	if e.ConfigName != "" {
		table = table.Where("config_name = ?", e.ConfigName)
	}
	if e.ConfigKey != "" {
		table = table.Where("config_key = ?", e.ConfigKey)
	}
	if e.ConfigType != "" {
		table = table.Where("config_type = ?", e.ConfigType)
	}
	var count int64
	if err := table.Where("`deleted_at` IS NULL").Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	return doc, count, nil
}

func (d *dao) UpdateSysConfig(e *model.SysConfig, id int) (update model.SysConfig, err error) {
	if err = d.orm.Table(e.TableName()).Where("config_id = ?", id).First(&update).Error; err != nil {
		return
	}

	if e.ConfigName != "" && e.ConfigName != update.ConfigName {
		return update, pb.UnexpectedErr("configName is null")
	}

	if e.ConfigKey != "" && e.ConfigKey != update.ConfigKey {
		return update, pb.UnexpectedErr("configName is null")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) DeleteSysConfig(e *model.SysConfig, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("config_id = ?", id).Delete(&model.SysConfig{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) BatchDeleteSysConfig(e *model.SysConfig, ids []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("config_id in (?)", ids).Delete(&model.SysConfig{}).Error; err != nil {
		return
	}
	Result = true
	return
}
