package model

type CmsSubject struct {
	Id           uint64 `json:"id" gorm:"type:bigint;primary_key"`      // id
	AppId        uint64 `json:"appappId" gorm:"type:bigint unsigned;"`  //
	CategoryId   uint64 `json:"categoryId" gorm:"type:bigint;"`         // 客户端id
	Title        string `json:"title" gorm:"type:varchar(255);"`        // 标题
	Pic          string `json:"pic" gorm:"type:varchar(255);"`          // pic
	ProductCount string `json:"productCount" gorm:"type:varchar(255);"` // 产品数量
	CreateBy     uint64 `json:"createBy" gorm:"type:bigint;"`           //
	UpdateBy     uint64 `json:"updateBy" gorm:"type:bigint;"`           //
	BaseModel
}

func (CmsSubject) TableName() string {
	return "cms_subject"
}
