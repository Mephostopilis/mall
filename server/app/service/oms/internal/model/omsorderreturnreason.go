package model

type OmsOrderReturnReason struct {
	Id         int    `json:"id" gorm:"type:bigint;primary_key"` //
	Name       string `json:"name" gorm:"type:varchar(100);"`    // 退货类型
	Sort       string `json:"sort" gorm:"type:int;"`             //
	Status     string `json:"status" gorm:"type:int;"`           // 状态：0->不启用；1->启用
	CreateTime string `json:"createTime" gorm:"type:datetime;"`  // 添加时间
	BaseModel
}

func (OmsOrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}
