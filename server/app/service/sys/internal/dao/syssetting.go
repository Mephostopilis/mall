package dao

import (
	"context"

	"edu/service/sys/internal/model"
)

//查询
func (d *dao) GetSysSetting(ctx context.Context, s *model.SysSetting) (create model.SysSetting, err error) {
	db := d.orm.WithContext(ctx).Table(s.TableName())
	if err = db.First(&create).Error; err != nil {
		d.log.Errorf("err = %v", err)
		return
	}
	return
}

//修改
func (d *dao) UpdateSysSetting(ctx context.Context, s *model.SysSetting) (update model.SysSetting, err error) {
	db := d.orm.WithContext(ctx).Table(s.TableName())
	if err = db.Model(&update).Updates(&s).Error; err != nil {
		return
	}
	return
}
