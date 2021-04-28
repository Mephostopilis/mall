package model

type CmsHelpCategory struct {
	Id         uint64 `json:"id" gorm:"type:bigint;primary_key"`   // id
	AppId      uint64 `json:"appId" gorm:"type:bigint unsigned;"`  //
	HelpCount  uint32 `json:"helpCount" gorm:"type:int unsigned;"` //
	Icon       string `json:"icon" gorm:"type:varchar(255);"`      //
	Name       string `json:"name" gorm:"type:varchar(255);"`      //
	ShowStatus int32  `json:"showStatus" gorm:"type:int;"`         //
	Sort       int32  `json:"sort" gorm:"type:int;"`               //
	CreateBy   uint64 `json:"createBy" gorm:"type:bigint;"`        //
	UpdateBy   uint64 `json:"updateBy" gorm:"type:bigint;"`        //
	BaseModel
}

func (CmsHelpCategory) TableName() string {
	return "cms_help_category"
}
