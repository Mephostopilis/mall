package dao

import (
	"context"

	"edu/service/sso/internal/model"
)

//CreateSsoApp 创建SsoApp
func (d *dao) CreateSsoApp(ctx context.Context, e *model.SsoApp) (model.SsoApp, error) {
	var doc model.SsoApp
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SsoApp
func (d *dao) GetSsoApp(ctx context.Context, e *model.SsoApp) (model.SsoApp, error) {
	var doc model.SsoApp
	table := d.orm.Table(e.TableName())

	if e.Id != "" {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SsoApp带分页
func (d *dao) GetSsoAppPage(ctx context.Context, e *model.SsoApp, pageSize int, pageIndex int) ([]model.SsoApp, int64, error) {
	var doc []model.SsoApp

	table := d.orm.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	// dataPermission := new(admin.DataPermission)
	// dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	// table, err := d.GetDataScope(table, e.TableName(), dataPermission)
	// if err != nil {
	// 	return nil, 0, err
	// }
	var count int64
	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 更新SsoApp
func (d *dao) UpdateSsoApp(ctx context.Context, e *model.SsoApp, id string) (update model.SsoApp, err error) {
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

// 删除SsoApp
func (d *dao) DeleteSsoApp(ctx context.Context, e *model.SsoApp, id string) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SsoApp{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSsoApp(ctx context.Context, e *model.SsoApp, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SsoApp{}).Error; err != nil {
		return
	}
	Result = true
	return
}
