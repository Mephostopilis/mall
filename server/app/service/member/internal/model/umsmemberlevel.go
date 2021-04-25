package model

import (
	"time"
)

type UmsMemberLevel struct {
	Access           string    `json:"access" gorm:"type:varchar(255);"`           // code
	AccessCreatedAt  time.Time `json:"accessCreatedAt" gorm:"type:timestamp;"`     // code创建时
	AccessExpiresIn  string    `json:"accessExpiresIn" gorm:"type:varchar(255);"`  // code创建时
	ClientId         string    `json:"clientId" gorm:"type:varchar(255);"`         // 客户端id
	Code             string    `json:"code" gorm:"type:varchar(255);"`             // code
	CodeCreatedAt    time.Time `json:"codeCreatedAt" gorm:"type:timestamp;"`       // code创建时
	CodeExpiresIn    string    `json:"codeExpiresIn" gorm:"type:varchar(255);"`    // code创建时
	CreateBy         string    `json:"createBy" gorm:"type:bigint;"`               //
	Id               int       `json:"id" gorm:"type:bigint;primary_key"`          // id
	RedirectUri      string    `json:"redirectUri" gorm:"type:varchar(255);"`      // redirect
	Refresh          string    `json:"refresh" gorm:"type:varchar(255);"`          // code
	RefreshCreatedAt time.Time `json:"refreshCreatedAt" gorm:"type:timestamp;"`    // code创建时
	RefreshExpiresIn string    `json:"refreshExpiresIn" gorm:"type:varchar(255);"` // code创建时
	Scope            string    `json:"scope" gorm:"type:varchar(255);"`            // 范围
	UpdateBy         string    `json:"updateBy" gorm:"type:bigint;"`               //
	UserId           string    `json:"userId" gorm:"type:varchar(255);"`           // 用户id
	BaseModel
}

func (UmsMemberLevel) TableName() string {
	return "ums_member_level"
}
