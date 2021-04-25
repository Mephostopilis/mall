package model

type PmsProductFullReduction struct {
	AppId       string `json:"appId" gorm:"type:bigint unsigned;"`     //
	FullPrice   string `json:"fullPrice" gorm:"type:decimal(10,2);"`   //
	Id          int    `json:"id" gorm:"type:bigint;primary_key"`      //
	ProductId   string `json:"productId" gorm:"type:bigint;"`          //
	ReducePrice string `json:"reducePrice" gorm:"type:decimal(10,2);"` //
	BaseModel
}

func (PmsProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
