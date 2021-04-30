package model

type SmsCouponProductCategoryRelation struct {
	Id                  uint64 `json:"id" gorm:"type:bigint;primary_key"`             //
	AppId               uint64 `json:"appId" gorm:"type:bigint unsigned;"`            //
	CouponId            uint64 `json:"couponId" gorm:"type:bigint;"`                  //
	ParentCategoryName  string `json:"parentCategoryName" gorm:"type:varchar(200);"`  // 父分类名称
	ProductCategoryId   uint64 `json:"productCategoryId" gorm:"type:bigint;"`         //
	ProductCategoryName string `json:"productCategoryName" gorm:"type:varchar(200);"` // 产品分类名称
	BaseModel
}

func (SmsCouponProductCategoryRelation) TableName() string {
	return "sms_coupon_product_category_relation"
}
