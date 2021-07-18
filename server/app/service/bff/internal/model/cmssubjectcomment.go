package model

type CmsSubjectComment struct {
	Id             uint64 `json:"id" gorm:"type:bigint;primary_key"`        //
	AppId          uint64 `json:"appId" gorm:"type:bigint unsigned;"`       //
	SubjectId      uint64 `json:"subjectId" gorm:"type:bigint;"`            //
	MemberIcon     string `json:"memberIcon" gorm:"type:varchar(255);"`     //
	MemberNickName string `json:"memberNickName" gorm:"type:varchar(255);"` //
	Content        string `json:"content" gorm:"type:varchar(1000);"`       //
	ShowStatus     int32  `json:"showStatus" gorm:"type:int;"`              //
	CreateBy       uint64 `json:"createBy" gorm:"type:bigint;"`             //
	UpdateBy       uint64 `json:"updateBy" gorm:"type:bigint;"`             //
	BaseModel
}

func (CmsSubjectComment) TableName() string {
	return "cms_subject_comment"
}
