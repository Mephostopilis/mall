package dao

import (
	"context"
	"edu/service/tiku/internal/model"
)

//CreateExercise 创建Exercise
func (d *dao) CreateExercise(ctx context.Context, e *model.Exercise) (model.Exercise, error) {
	var doc model.Exercise
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Exercise
func (d *dao) GetExercise(ctx context.Context, e *model.Exercise) (model.Exercise, error) {
	var doc model.Exercise
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取Exercise带分页
func (d *dao) GetExercisePage(ctx context.Context, e *model.Exercise, pageSize int, pageIndex int) ([]model.Exercise, int64, error) {
	var doc []model.Exercise
	table := d.orm.Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	// dataPermission := new(model.DataPermission)
	// dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	// table, err := d.GetDataScope(table, e.TableName(), dataPermission)
	// if err != nil {
	// 	return nil, 0, err
	// }
	if e.Level != 0 {
		table = table.Where("level = ?", e.Level)
	}
	if e.Ty != 0 {
		table = table.Where("ty = ?", e.Ty)
	}
	var count int64
	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 更新Exercise
func (d *dao) UpdateExercise(ctx context.Context, e *model.Exercise, id uint64) (update model.Exercise, err error) {
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

// 删除Exercise
func (d *dao) DeleteExercise(ctx context.Context, e *model.Exercise, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.Exercise{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteExercise(ctx context.Context, e *model.Exercise, id []uint64) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.Exercise{}).Error; err != nil {
		return
	}
	Result = true
	return
}
