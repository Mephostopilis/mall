package model

import (
	"time"
)

type UmsMemberAssets struct {
	Birthday     time.Time `json:"birthday" gorm:"type:timestamp;"`               //
	CreateBy     string    `json:"createBy" gorm:"type:bigint;"`                  //
	Email        string    `json:"email" gorm:"type:varchar(255);"`               //
	LastSignInAt time.Time `json:"lastSignInAt" gorm:"type:timestamp;"`           //
	Mobile       string    `json:"mobile" gorm:"type:varchar(255);"`              //
	Password     string    `json:"password" gorm:"type:varchar(255);"`            //
	Privilege    string    `json:"privilege" gorm:"type:int;"`                    //
	Uid          string    `json:"uid" gorm:"type:bigint unsigned;"`              //
	UpdateBy     string    `json:"updateBy" gorm:"type:bigint;"`                  //
	Username     string    `json:"username" gorm:"type:varchar(255);primary_key"` //
	BaseModel
}

func (UmsMemberAssets) TableName() string {
	return "ums_member_assets"
}
