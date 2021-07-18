package model

type CmsHelp struct {
	Id         uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	AppId      uint64 `json:"appId" gorm:"type:bigint unsigned;"`         //
	CategoryId uint64 `json:"categoryId" gorm:"type:bigint;"`             // 分类
	Content    string `json:"content" gorm:"type:text;"`                  // 内容
	Icon       string `json:"icon" gorm:"type:varchar(255);"`             // icon
	ReadCount  int32  `json:"readCount" gorm:"type:int;"`                 // 读取数
	ShowStatus int32  `json:"showStatus" gorm:"type:int;"`                // 显示状态
	Title      string `json:"title" gorm:"type:varchar(255);"`            // title
	CreateBy   uint64 `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy   uint64 `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (CmsHelp) TableName() string {
	return "cms_help"
}
