package dao

import (
	"context"

	"edu/service/sso/internal/model"
)

//CreateSsoToken 创建SsoToken
func (d *dao) CreateSsoToken(ctx context.Context, e *model.SsoToken) (model.SsoToken, error) {
	var doc model.SsoToken
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SsoToken
func (d *dao) GetSsoToken(ctx context.Context, e *model.SsoToken) (model.SsoToken, error) {
	var doc model.SsoToken
	table := d.orm.Table(e.TableName())
	if e.Id != "" {
		table = table.Where("id = ?", e.Id)
	}
	if e.Code != "" {
		table = table.Where("code = ?", e.Code)
	}
	if e.Access != "" {
		table = table.Where("access = ?", e.Access)
	}
	if e.Refresh != "" {
		table = table.Where("refresh = ?", e.Refresh)
	}
	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SsoToken带分页
func (d *dao) GetSsoTokenPage(ctx context.Context, e *model.SsoToken, pageSize int, pageIndex int) ([]model.SsoToken, int64, error) {
	var doc []model.SsoToken

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

// 更新SsoToken
func (d *dao) UpdateSsoToken(ctx context.Context, e *model.SsoToken, id string) (update model.SsoToken, err error) {
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

// 删除SsoToken
func (d *dao) DeleteSsoToken(ctx context.Context, e *model.SsoToken, id string) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.SsoToken{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (d *dao) DeleteSsoTokenX(ctx context.Context, e *model.SsoToken) (success bool, err error) {
	table := d.orm.Table(e.TableName())
	if e.Code != "" {
		table = table.Where("code = ?", e.Code)
	}
	if e.Access != "" {
		table = table.Where("access = ?", e.Access)
	}
	if e.Refresh != "" {
		table = table.Where("refresh = ?", e.Refresh)
	}
	if err = table.Delete(&model.SsoToken{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSsoToken(ctx context.Context, e *model.SsoToken, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.SsoToken{}).Error; err != nil {
		return
	}
	Result = true
	return
}
