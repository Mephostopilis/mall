package dao

import (
	"context"

	"edu/service/tiku/internal/model"
)

//CreateUserExercise 创建UserExercise
func (d *dao) CreateUserExercise(ctx context.Context, e *model.UserExercise) (model.UserExercise, error) {
	var doc model.UserExercise
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UserExercise
func (d *dao) GetUserExercise(ctx context.Context, e *model.UserExercise) (model.UserExercise, error) {
	var doc model.UserExercise
	table := d.orm.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取UserExercise带分页
func (d *dao) GetUserExercisePage(ctx context.Context, e *model.UserExercise, pageSize int, pageIndex int) ([]model.UserExercise, int64, error) {
	var doc []model.UserExercise

	table := d.orm.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	// dataPermission := new(model.DataPermission)
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

// 更新UserExercise
func (d *dao) UpdateUserExercise(ctx context.Context, e *model.UserExercise, id int) (update model.UserExercise, err error) {
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

// 删除UserExercise
func (d *dao) DeleteUserExercise(ctx context.Context, e *model.UserExercise, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id = ?", id).Delete(&model.UserExercise{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteUserExercise(ctx context.Context, e *model.UserExercise, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("id in (?)", id).Delete(&model.UserExercise{}).Error; err != nil {
		return
	}
	Result = true
	return
}
