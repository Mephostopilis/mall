package model

type SsoApp struct {
	Id        string `json:"id" gorm:"type:varchar(255);primary_key"` // id
	Secret    string `json:"secret" gorm:"type:varchar(255);"`        // 密钥
	Domain    string `json:"domain" gorm:"type:varchar(255);"`        // domain
	UserId    string `json:"userId" gorm:"type:varchar(255);"`        // 用户id
	Scope     string `json:"scope" gorm:"type:varchar(255);"`         // 用户id
	GrantType string `json:"grant_type" gorm:"type:varchar(255);"`    // 用户id
	CreateBy  string `json:"createBy" gorm:"type:bigint(20);"`        //
	UpdateBy  string `json:"updateBy" gorm:"type:bigint(20);"`        //
	BaseModel
}

func (SsoApp) TableName() string {
	return "ums_app"
}

func (m *SsoApp) GetID() string {
	return m.Id
}

func (m *SsoApp) GetSecret() string {
	return m.Secret
}

func (m *SsoApp) GetDomain() string {
	return m.Domain
}

func (m *SsoApp) GetUserID() string {
	return m.UserId
}
