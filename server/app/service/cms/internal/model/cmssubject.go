package model

type CmsSubject struct {
	Id           int32  `json:"id" gorm:"type:int;primary_key"`         // id
	AppId        string `json:"appId" gorm:"type:bigint unsigned;"`     //
	CategoryId   string `json:"categoryId" gorm:"type:int;"`            // 客户端id
	Pic          string `json:"pic" gorm:"type:varchar(255);"`          // pic
	ProductCount string `json:"productCount" gorm:"type:varchar(255);"` // 产品数量
	Title        string `json:"title" gorm:"type:varchar(255);"`        // 标题
	CreateBy     string `json:"createBy" gorm:"type:bigint;"`           //
	UpdateBy     string `json:"updateBy" gorm:"type:bigint;"`           //
	BaseModel
}

func (CmsSubject) TableName() string {
	return "cms_subject"
}
