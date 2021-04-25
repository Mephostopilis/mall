package model

import (
	"time"

	"github.com/go-oauth2/oauth2/v4"
)

type SsoToken struct {
	Id               string    `json:"id" gorm:"type:varchar(255);primary_key"`    // id
	AppId            uint64    `json:"appId" gorm:"type:bigint(20);"`              // 客户端id
	ClientId         string    `json:"clientId" gorm:"type:varchar(255);"`         // 客户端id
	UserId           string    `json:"userId" gorm:"type:varchar(255);"`           // 用户id
	RedirectUri      string    `json:"redirectUri" gorm:"type:varchar(255);"`      // redirect
	Scope            string    `json:"scope" gorm:"type:varchar(255);"`            // 范围
	Code             string    `json:"code" gorm:"type:varchar(255);"`             // code
	CodeCreatedAt    time.Time `json:"codeCreatedAt" gorm:"type:timestamp;"`       // code创建时
	CodeExpiresIn    string    `json:"codeExpiresIn" gorm:"type:varchar(255);"`    // code创建时
	Access           string    `json:"access" gorm:"type:varchar(255);"`           // code
	AccessCreatedAt  time.Time `json:"accessCreatedAt" gorm:"type:timestamp;"`     // code创建时
	AccessExpiresIn  string    `json:"accessExpiresIn" gorm:"type:varchar(255);"`  // code创建时
	Refresh          string    `json:"refresh" gorm:"type:varchar(255);"`          // code
	RefreshCreatedAt time.Time `json:"refreshCreatedAt" gorm:"type:timestamp;"`    // code创建时
	RefreshExpiresIn string    `json:"refreshExpiresIn" gorm:"type:varchar(255);"` // code创建时
	CreateBy         string    `json:"createBy" gorm:"type:bigint(20);"`           //
	UpdateBy         string    `json:"updateBy" gorm:"type:bigint(20);"`           //
	BaseModel
}

func (SsoToken) TableName() string {
	return "ums_token"
}

func (m *SsoToken) New() oauth2.TokenInfo {
	return &SsoToken{}
}

func (m *SsoToken) GetClientID() string {
	return m.ClientId
}

func (m *SsoToken) SetClientID(clientId string) {
	m.ClientId = clientId
	return
}

func (m *SsoToken) GetUserID() string {
	return m.UserId
}

func (m *SsoToken) SetUserID(userId string) {
	m.UserId = userId
	return
}

func (m *SsoToken) GetRedirectURI() string {
	return m.RedirectUri
}

func (m *SsoToken) SetRedirectURI(redirectUri string) {
	m.RedirectUri = redirectUri
	return
}
func (m *SsoToken) GetScope() string {
	return m.Scope
}

func (m *SsoToken) SetScope(scope string) {
	m.Scope = scope
}

func (m *SsoToken) GetCode() string {
	return m.Code
}

func (m *SsoToken) SetCode(code string) {
	m.Code = code
	return
}

func (m *SsoToken) GetCodeCreateAt() time.Time {
	return m.CodeCreatedAt
}

func (m *SsoToken) SetCodeCreateAt(ti time.Time) {
	m.CodeCreatedAt = ti
	return
}

func (m *SsoToken) GetCodeExpiresIn() time.Duration {
	du, err := time.ParseDuration(m.CodeExpiresIn)
	if err != nil {
		return time.Duration(0)
	}
	return du
}

func (m *SsoToken) SetCodeExpiresIn(ti time.Duration) {
	m.CodeExpiresIn = ti.String()
	return
}

func (m *SsoToken) GetAccess() string {
	return m.Access
}

func (m *SsoToken) SetAccess(access string) {
	m.Access = access
	return
}

func (m *SsoToken) GetAccessCreateAt() time.Time {
	return m.AccessCreatedAt
}

func (m *SsoToken) SetAccessCreateAt(ti time.Time) {
	m.AccessCreatedAt = ti
	return
}

func (m *SsoToken) GetAccessExpiresIn() time.Duration {
	du, err := time.ParseDuration(m.AccessExpiresIn)
	if err != nil {
		return time.Duration(0)
	}
	return du
}

func (m *SsoToken) SetAccessExpiresIn(time.Duration) {
	return
}

func (m *SsoToken) GetRefresh() string {
	return m.Refresh
}

func (m *SsoToken) SetRefresh(refresh string) {
	m.Refresh = refresh
	return
}

func (m *SsoToken) GetRefreshCreateAt() time.Time {
	return m.RefreshCreatedAt
}

func (m *SsoToken) SetRefreshCreateAt(ti time.Time) {
	m.RefreshCreatedAt = ti
	return
}

func (m *SsoToken) GetRefreshExpiresIn() time.Duration {
	du, err := time.ParseDuration(m.RefreshExpiresIn)
	if err != nil {
		return time.Duration(0)
	}
	return du
}

func (m *SsoToken) SetRefreshExpiresIn(du time.Duration) {
	m.RefreshExpiresIn = du.String()
	return
}
