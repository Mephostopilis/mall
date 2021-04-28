package model

type SmsCoupon struct {
	Id       uint64  `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	AppId    uint64  `json:"appId" gorm:"type:bigint unsigned;"`         //
	Name     string  `json:"name" gorm:"type:varchar(255);"`             // name
	Amount   float32 `json:"amount" gorm:"type:decimal(10,0);"`          // 价格
	Count    int32   `json:"count" gorm:"type:int;"`                     // 数量
	Platform int32   `json:"platform" gorm:"type:int;"`                  // 用户id
	Type     int32   `json:"type" gorm:"type:int;"`                      // 类型
	CreateBy uint64  `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy uint64  `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (SmsCoupon) TableName() string {
	return "sms_coupon"
}
