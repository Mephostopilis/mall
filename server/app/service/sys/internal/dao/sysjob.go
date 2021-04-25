package dao

import (
	"edu/service/sys/internal/model"
)

// 创建SysJob
func (d *dao) CreateSysJob(e *model.SysJob) (model.SysJob, error) {
	var doc model.SysJob
	result := d.orm.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SysJob
func (d *dao) GetSysJob(e *model.SysJob) (model.SysJob, error) {
	var doc model.SysJob
	table := d.orm.Table(e.TableName())

	if e.JobId != 0 {
		table = table.Where("job_id = ?", e.JobId)
	}

	if e.JobName != "" {
		table = table.Where("job_name like ?", "%"+e.JobName+"%")
	}

	if e.JobGroup != "" {
		table = table.Where("job_group = ?", e.JobGroup)
	}

	if e.CronExpression != "" {
		table = table.Where("cron_expression = ?", e.CronExpression)
	}

	if e.InvokeTarget != "" {
		table = table.Where("invoke_target = ?", e.InvokeTarget)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SysJob带分页
func (d *dao) GetSysJobPage(e *model.SysJob, pageSize int, pageIndex int) ([]model.SysJob, int64, error) {
	var doc []model.SysJob

	table := d.orm.Table(e.TableName())

	if e.JobId != 0 {
		table = table.Where("job_id = ?", e.JobId)
	}

	if e.JobName != "" {
		table = table.Where("job_name like ?", "%"+e.JobName+"%")
	}

	if e.JobGroup != "" {
		table = table.Where("job_group = ?", e.JobGroup)
	}

	if e.CronExpression != "" {
		table = table.Where("cron_expression = ?", e.CronExpression)
	}

	if e.InvokeTarget != "" {
		table = table.Where("invoke_target = ?", e.InvokeTarget)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	var count int64
	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (d *dao) GetSysJobList(e *model.SysJob) ([]model.SysJob, error) {
	var doc []model.SysJob

	table := d.orm.Select("*").Table(e.TableName())
	table = table.Where("status = ?", 2)

	if err := table.Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

// 更新SysJob
func (d *dao) UpdateSysJob(e *model.SysJob, id int) (update model.SysJob, err error) {
	if err = d.orm.Table(e.TableName()).Where("job_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (d *dao) RemoveSysJobAllEntryID(e *model.SysJob) (update model.SysJob, err error) {

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Updates(map[string]interface{}{"entry_id": 0}).Error; err != nil {
		return
	}
	return
}

func (d *dao) RemoveSysJobEntryID(e *model.SysJob, entryID int) (update model.SysJob, err error) {

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = d.orm.Table(e.TableName()).Where("entry_id = ?", entryID).Updates(map[string]interface{}{"entry_id": 0}).Error; err != nil {
		return
	}
	return
}

// 删除SysJob
func (d *dao) DeleteSysJob(e *model.SysJob, id int) (success bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("job_id = ?", id).Delete(&model.SysJob{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (d *dao) BatchDeleteSysJob(e *model.SysJob, id []int) (Result bool, err error) {
	if err = d.orm.Table(e.TableName()).Where("job_id in (?)", id).Delete(&model.SysJob{}).Error; err != nil {
		return
	}
	Result = true
	return
}
