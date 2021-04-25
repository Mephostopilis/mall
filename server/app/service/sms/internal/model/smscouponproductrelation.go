package model

import (


)

type SmsCouponProductRelation struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        CouponId string `json:"couponId" gorm:"type:bigint;"` // 
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        ProductId string `json:"productId" gorm:"type:bigint;"` // 
        ProductName string `json:"productName" gorm:"type:varchar(500);"` // 商品名称
        ProductSn string `json:"productSn" gorm:"type:varchar(200);"` // 商品编码
BaseModel
}

func (SmsCouponProductRelation) TableName() string {
return "sms_coupon_product_relation"
}

