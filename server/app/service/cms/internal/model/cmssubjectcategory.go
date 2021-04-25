package model

import (


)

type CmsSubjectCategory struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        Icon string `json:"icon" gorm:"type:varchar(500);"` // 分类图标
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        Name string `json:"name" gorm:"type:varchar(100);"` // 
        ShowStatus string `json:"showStatus" gorm:"type:int;"` // 
        Sort string `json:"sort" gorm:"type:int;"` // 
        SubjectCount string `json:"subjectCount" gorm:"type:int;"` // 专题数量
BaseModel
}

func (CmsSubjectCategory) TableName() string {
return "cms_subject_category"
}

