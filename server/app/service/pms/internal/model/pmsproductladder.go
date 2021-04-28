package model

type PmsProductLadder struct {
	AppId     string `json:"appId" gorm:"type:bigint unsigned;"`  //
	Count     string `json:"count" gorm:"type:int;"`              // 满足的商品数量
	Discount  string `json:"discount" gorm:"type:decimal(10,2);"` // 折扣
	Id        int    `json:"id" gorm:"type:bigint;primary_key"`   //
	Price     string `json:"price" gorm:"type:decimal(10,2);"`    // 折后价格
	ProductId string `json:"productId" gorm:"type:bigint;"`       //
	CreateBy  uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy  uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsProductLadder) TableName() string {
	return "pms_product_ladder"
}
