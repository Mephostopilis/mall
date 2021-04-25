package model

type CmsHelp struct {
	Id         int32  `json:"id" gorm:"type:int;primary_key"`     // id
	AppId      string `json:"appId" gorm:"type:bigint unsigned;"` //
	CategoryId string `json:"categoryId" gorm:"type:bigint;"`     // 分类
	Content    string `json:"content" gorm:"type:text;"`          // 内容
	Icon       string `json:"icon" gorm:"type:varchar(255);"`     // icon
	ReadCount  string `json:"readCount" gorm:"type:int;"`         // 读取数
	ShowStatus string `json:"showStatus" gorm:"type:int;"`        // 显示状态
	Title      string `json:"title" gorm:"type:varchar(255);"`    // title
	CreateBy   string `json:"createBy" gorm:"type:bigint;"`       //
	UpdateBy   string `json:"updateBy" gorm:"type:bigint;"`       //
	BaseModel
}

func (CmsHelp) TableName() string {
	return "cms_help"
}
