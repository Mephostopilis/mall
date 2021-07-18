package model

type CmsSubjectCategory struct {
	Id           uint64 `json:"id" gorm:"type:bigint;primary_key"`  //
	AppId        string `json:"appId" gorm:"type:bigint unsigned;"` //
	Icon         string `json:"icon" gorm:"type:varchar(500);"`     // 分类图标
	Name         string `json:"name" gorm:"type:varchar(100);"`     //
	SubjectCount int32  `json:"subjectCount" gorm:"type:int;"`      // 专题数量
	ShowStatus   int32  `json:"showStatus" gorm:"type:int;"`        //
	Sort         int32  `json:"sort" gorm:"type:int;"`              //
	CreateBy     uint64 `json:"createBy" gorm:"type:bigint;"`       //
	UpdateBy     uint64 `json:"updateBy" gorm:"type:bigint;"`       //
	BaseModel
}

func (CmsSubjectCategory) TableName() string {
	return "cms_subject_category"
}
