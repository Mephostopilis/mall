package model

type CmsPrefrenceArea struct {
	Id         uint64 `json:"id" gorm:"type:bigint;primary_key"`  //
	AppId      uint64 `json:"appId" gorm:"type:bigint unsigned;"` //
	Name       string `json:"name" gorm:"type:varchar(255);"`     //
	SubTitle   string `json:"subTitle" gorm:"type:varchar(255);"` //
	Pic        string `json:"pic" gorm:"type:varbinary(500);"`    // 展示图片
	ShowStatus int32  `json:"showStatus" gorm:"type:int;"`        //
	Sort       int32  `json:"sort" gorm:"type:int;"`              //
	BaseModel
}

func (CmsPrefrenceArea) TableName() string {
	return "cms_prefrence_area"
}
