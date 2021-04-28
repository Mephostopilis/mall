package model

type PmsProductFullReduction struct {
	AppId       string `json:"appId" gorm:"type:bigint unsigned;"`     //
	FullPrice   string `json:"fullPrice" gorm:"type:decimal(10,2);"`   //
	Id          int    `json:"id" gorm:"type:bigint;primary_key"`      //
	ProductId   string `json:"productId" gorm:"type:bigint;"`          //
	ReducePrice string `json:"reducePrice" gorm:"type:decimal(10,2);"` //
	CreateBy    uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy    uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
