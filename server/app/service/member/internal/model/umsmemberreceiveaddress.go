package model

import (


)

type UmsMemberReceiveAddress struct {

        City string `json:"city" gorm:"type:varchar(100);"` // 城市
        DefaultStatus string `json:"defaultStatus" gorm:"type:int;"` // 是否为默认
        DetailAddress string `json:"detailAddress" gorm:"type:varchar(128);"` // 详细地址(街道)
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        MemberId string `json:"memberId" gorm:"type:bigint;"` // 
        Name string `json:"name" gorm:"type:varchar(100);"` // 收货人名称
        PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(64);"` // 
        PostCode string `json:"postCode" gorm:"type:varchar(100);"` // 邮政编码
        Province string `json:"province" gorm:"type:varchar(100);"` // 省份/直辖市
        Region string `json:"region" gorm:"type:varchar(100);"` // 区
BaseModel
}

func (UmsMemberReceiveAddress) TableName() string {
return "ums_member_receive_address"
}

