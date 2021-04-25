package model

import (


)

type OmsCompanyAddress struct {

        AddressName string `json:"addressName" gorm:"type:varchar(200);"` // 地址名称
        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        City string `json:"city" gorm:"type:varchar(64);"` // 市
        DetailAddress string `json:"detailAddress" gorm:"type:varchar(200);"` // 详细地址
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        Name string `json:"name" gorm:"type:varchar(64);"` // 收发货人姓名
        Phone string `json:"phone" gorm:"type:varchar(64);"` // 收货人电话
        Province string `json:"province" gorm:"type:varchar(64);"` // 省/直辖市
        ReceiveStatus string `json:"receiveStatus" gorm:"type:int;"` // 是否默认收货地址：0->否；1->是
        Region string `json:"region" gorm:"type:varchar(64);"` // 区
        SendStatus string `json:"sendStatus" gorm:"type:int;"` // 默认发货地址：0->否；1->是
BaseModel
}

func (OmsCompanyAddress) TableName() string {
return "oms_company_address"
}

