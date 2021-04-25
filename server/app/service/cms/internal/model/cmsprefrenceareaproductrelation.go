package model

import (


)

type CmsPrefrenceAreaProductRelation struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        PrefrenceAreaId string `json:"prefrenceAreaId" gorm:"type:bigint;"` // 
        ProductId string `json:"productId" gorm:"type:bigint;"` // 
BaseModel
}

func (CmsPrefrenceAreaProductRelation) TableName() string {
return "cms_prefrence_area_product_relation"
}

