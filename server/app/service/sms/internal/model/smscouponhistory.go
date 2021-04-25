package model

import (


)

type SmsCouponHistory struct {

        AppId string `json:"appId" gorm:"type:bigint unsigned;"` // 
        CouponCode string `json:"couponCode" gorm:"type:varchar(64);"` // 
        CouponId string `json:"couponId" gorm:"type:bigint;"` // 
        CreateTime string `json:"createTime" gorm:"type:datetime;"` // 
        GetType string `json:"getType" gorm:"type:int;"` // 获取类型：0->后台赠送；1->主动获取
        Id int `json:"id" gorm:"type:bigint;primary_key"` // 
        MemberId string `json:"memberId" gorm:"type:bigint;"` // 
        MemberNickname string `json:"memberNickname" gorm:"type:varchar(64);"` // 领取人昵称
        OrderId string `json:"orderId" gorm:"type:bigint;"` // 订单编号
        OrderSn string `json:"orderSn" gorm:"type:varchar(100);"` // 订单号码
        UseStatus string `json:"useStatus" gorm:"type:int;"` // 使用状态：0->未使用；1->已使用；2->已过期
        UseTime string `json:"useTime" gorm:"type:datetime;"` // 使用时间
BaseModel
}

func (SmsCouponHistory) TableName() string {
return "sms_coupon_history"
}

