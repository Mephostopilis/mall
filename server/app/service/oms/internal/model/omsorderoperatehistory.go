package model

type OmsOrderOperateHistory struct {
	Id          int    `json:"id" gorm:"type:bigint;primary_key"`    //
	OrderId     string `json:"orderId" gorm:"type:bigint;"`          // 订单id
	OperateMan  string `json:"operateMan" gorm:"type:varchar(100);"` // 操作人：用户；系统；后台管理员
	CreateTime  string `json:"createTime" gorm:"type:datetime;"`     // 操作时间
	OrderStatus string `json:"orderStatus" gorm:"type:int;"`         // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string `json:"note" gorm:"type:varchar(500);"`       // 备注
	BaseModel
}

func (OmsOrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}
