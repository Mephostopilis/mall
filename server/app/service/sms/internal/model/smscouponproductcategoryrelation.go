package model

import (


)

type SmsCouponProductCategoryRelation struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        CouponId string `json:"couponId" gorm:"type:bigint;"` // 
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        ParentCategoryName string `json:"parentCategoryName" gorm:"type:varchar(200);"` // 父分类名称
        ProductCategoryId string `json:"productCategoryId" gorm:"type:bigint;"` // 
        ProductCategoryName string `json:"productCategoryName" gorm:"type:varchar(200);"` // 产品分类名称
BaseModel
}

func (SmsCouponProductCategoryRelation) TableName() string {
return "sms_coupon_product_category_relation"
}

