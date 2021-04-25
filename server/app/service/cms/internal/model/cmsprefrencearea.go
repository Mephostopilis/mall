package model

import (


)

type CmsPrefrenceArea struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        Name string `json:"name" gorm:"type:varchar(255);"` // 
        Pic string `json:"pic" gorm:"type:varbinary(500);"` // 展示图片
        ShowStatus string `json:"showStatus" gorm:"type:int;"` // 
        Sort string `json:"sort" gorm:"type:int;"` // 
        SubTitle string `json:"subTitle" gorm:"type:varchar(255);"` // 
BaseModel
}

func (CmsPrefrenceArea) TableName() string {
return "cms_prefrence_area"
}

