package model

import "time"

type SmsCouponHistory struct {
	Id             uint64    `json:"id" gorm:"type:bigint;primary_key"`       //
	AppId          uint64    `json:"appId" gorm:"type:bigint unsigned;"`      //
	CouponCode     string    `json:"couponCode" gorm:"type:varchar(64);"`     //
	CouponId       uint64    `json:"couponId" gorm:"type:bigint;"`            //
	MemberId       uint64    `json:"memberId" gorm:"type:bigint;"`            //
	GetType        int32     `json:"getType" gorm:"type:int;"`                // 获取类型：0->后台赠送；1->主动获取
	MemberNickname string    `json:"memberNickname" gorm:"type:varchar(64);"` // 领取人昵称
	OrderId        uint64    `json:"orderId" gorm:"type:bigint;"`             // 订单编号
	OrderSn        string    `json:"orderSn" gorm:"type:varchar(100);"`       // 订单号码
	UseStatus      int32     `json:"useStatus" gorm:"type:int;"`              // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        time.Time `json:"useTime" gorm:"type:datetime;"`           // 使用时间
	BaseModel
}

func (SmsCouponHistory) TableName() string {
	return "sms_coupon_history"
}
