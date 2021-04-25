package model

import (


)

type CmsSubjectComment struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        Content string `json:"content" gorm:"type:varchar(1000);"` // 
        CreateTime string `json:"createTime" gorm:"type:datetime;"` // 
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        MemberIcon string `json:"memberIcon" gorm:"type:varchar(255);"` // 
        MemberNickName string `json:"memberNickName" gorm:"type:varchar(255);"` // 
        ShowStatus string `json:"showStatus" gorm:"type:int;"` // 
        SubjectId string `json:"subjectId" gorm:"type:bigint;"` // 
BaseModel
}

func (CmsSubjectComment) TableName() string {
return "cms_subject_comment"
}

