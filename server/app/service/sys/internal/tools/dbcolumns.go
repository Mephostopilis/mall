package tools

import (
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

func (d *dao) GetDBColumnsPage(dbname string, e *model.DBColumns, pageSize int, pageIndex int) (doc []model.DBColumns, count int64, err error) {
	table := d.orm.Table("`COLUMNS`")
	table = table.Where("table_schema= ? ", dbname)
	if e.TableName != "" {
		err = errors.Unknown("table name cannot be empty！", "")
		return
	}
	table = table.Where("TABLE_NAME = ?", e.TableName)

	if err = table.Count(&count).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return
	}
	return

}

func (d *dao) GetDBColumnsList(dbname string, e *model.DBColumns) (doc []model.DBColumns, err error) {

	table := new(gorm.DB)
	if e.TableName == "" {
		return nil, errors.Unknown("table name cannot be empty！", "")
	}

	table = d.orm.Table("information_schema.columns")
	table = table.Where("table_schema= ? ", dbname)
	table = table.Where("TABLE_NAME = ?", e.TableName)

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
