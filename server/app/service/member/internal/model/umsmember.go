package model

type UmsMember struct {
	CreateBy  string `json:"createBy" gorm:"type:bigint;"`        //
	Domain    string `json:"domain" gorm:"type:varchar(255);"`    // domain
	GrantType string `json:"grantType" gorm:"type:varchar(255);"` // 授权类型
	Id        int    `json:"id" gorm:"type:bigint;primary_key"`   // id
	Scope     string `json:"scope" gorm:"type:varchar(255);"`     // 权限范围
	Secret    string `json:"secret" gorm:"type:varchar(255);"`    // 密钥
	UpdateBy  string `json:"updateBy" gorm:"type:bigint;"`        //
	UserId    string `json:"userId" gorm:"type:varchar(255);"`    // 用户id
	BaseModel
}

func (UmsMember) TableName() string {
	return "ums_member"
}
