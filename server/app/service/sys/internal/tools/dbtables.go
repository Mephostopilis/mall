package tools

import (
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (d *dao) GetDBTablesPage(dbname string, e *model.DBTables, pageSize int, pageIndex int) (doc []model.DBTables, count int64, err error) {

	table := new(gorm.DB)
	table = d.orm.Table("tables")
	// table = table.Where("TABLE_NAME not in (select table_name from " + dbname + ".sys_tables) ")
	table = table.Where("table_schema= ? ", dbname)

	if e.TableName != "" {
		table = table.Where("TABLE_NAME = ?", e.TableName)
	}
	if err = table.Count(&count).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return
	}
	return
}

func (d *dao) GetDBTables(dbname string, e *model.DBTables) (doc model.DBTables, err error) {

	table := new(gorm.DB)
	table = d.orm.Table("tables")
	table = table.Where("table_schema= ? ", dbname)
	if e.TableName == "" {
		return doc, errors.Unknown("table name cannot be empty", "")
	}
	table = table.Where("TABLE_NAME = ?", e.TableName)

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
