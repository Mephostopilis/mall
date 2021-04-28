package model

type OmsOrderSetting struct {
	Id                  int    `json:"id" gorm:"type:bigint;primary_key"`    //
	FlashOrderOvertime  string `json:"flashOrderOvertime" gorm:"type:int;"`  // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime string `json:"normalOrderOvertime" gorm:"type:int;"` // 正常订单超时时间(分)
	ConfirmOvertime     string `json:"confirmOvertime" gorm:"type:int;"`     // 发货后自动确认收货时间（天）
	FinishOvertime      string `json:"finishOvertime" gorm:"type:int;"`      // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     string `json:"commentOvertime" gorm:"type:int;"`     // 订单完成后自动好评时间（天）
	BaseModel
}

func (OmsOrderSetting) TableName() string {
	return "oms_order_setting"
}
