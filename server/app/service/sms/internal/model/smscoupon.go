package model

type SmsCoupon struct {
	Amount   string `json:"amount" gorm:"type:decimal(10,0);"`          // 价格
	AppId    string `json:"appId" gorm:"type:bigint unsigned;"`         //
	Count    string `json:"count" gorm:"type:int;"`                     // 数量
	CreateBy string `json:"createBy" gorm:"type:bigint;"`               //
	Id       int    `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	Name     string `json:"name" gorm:"type:varchar(255);"`             // name
	Platform string `json:"platform" gorm:"type:int;"`                  // 用户id
	Type     string `json:"type" gorm:"type:int;"`                      // 类型
	UpdateBy string `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (SmsCoupon) TableName() string {
	return "sms_coupon"
}
