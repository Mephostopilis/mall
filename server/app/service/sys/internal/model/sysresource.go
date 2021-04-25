package model

type SysResource struct {
	Id       int32  `json:"id" gorm:"type:int;primary_key"`     //
	V1       string `json:"v1" gorm:"type:varchar(255);"`       //
	V2       string `json:"v2" gorm:"type:varchar(255);"`       //
	V3       string `json:"v3" gorm:"type:varchar(255);"`       //
	V4       string `json:"v4" gorm:"type:varchar(255);"`       //
	V5       string `json:"v5" gorm:"type:varchar(255);"`       //
	CreateBy string `json:"createBy" gorm:"type:varchar(128);"` //
	UpdateBy string `json:"updateBy" gorm:"type:varchar(128);"` //
	BaseModel
}

func (SysResource) TableName() string {
	return "sys_resource"
}
