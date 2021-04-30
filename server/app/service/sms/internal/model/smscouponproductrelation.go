package model

type SmsCouponProductRelation struct {
	Id          uint64 `json:"id" gorm:"type:bigint;primary_key"`     //
	AppId       uint64 `json:"appId" gorm:"type:bigint unsigned;"`    //
	CouponId    uint64 `json:"couponId" gorm:"type:bigint;"`          //
	ProductId   string `json:"productId" gorm:"type:bigint;"`         //
	ProductName string `json:"productName" gorm:"type:varchar(500);"` // 商品名称
	ProductSn   string `json:"productSn" gorm:"type:varchar(200);"`   // 商品编码
	BaseModel
}

func (SmsCouponProductRelation) TableName() string {
	return "sms_coupon_product_relation"
}
