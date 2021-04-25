package model

type CmsHelpCategory struct {
	AppId      string `json:"appId" gorm:"type:bigint unsigned;"`  //
	CreateBy   string `json:"createBy" gorm:"type:bigint;"`        //
	HelpCount  string `json:"helpCount" gorm:"type:int unsigned;"` //
	Icon       string `json:"icon" gorm:"type:varchar(255);"`      //
	Id         int    `json:"id" gorm:"type:int;primary_key"`      // id
	Name       string `json:"name" gorm:"type:varchar(255);"`      //
	ShowStatus string `json:"showStatus" gorm:"type:int;"`         //
	Sort       string `json:"sort" gorm:"type:int;"`               //
	UpdateBy   string `json:"updateBy" gorm:"type:bigint;"`        //
	BaseModel
}

func (CmsHelpCategory) TableName() string {
	return "cms_help_category"
}
