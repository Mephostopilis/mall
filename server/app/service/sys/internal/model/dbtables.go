package model

import "time"

type DBTables struct {
	TableName      string    `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string    `gorm:"column:ENGINE" json:"engine"`
	TableRows      string    `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string    `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     time.Time `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string    `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}
