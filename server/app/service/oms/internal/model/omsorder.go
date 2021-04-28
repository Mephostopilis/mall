package model

type OmsOrder struct {
	Id                    int    `json:"id" gorm:"type:bigint;primary_key"`               // 订单id
	MemberId              string `json:"memberId" gorm:"type:bigint;"`                    //
	CouponId              string `json:"couponId" gorm:"type:bigint;"`                    //
	OrderSn               string `json:"orderSn" gorm:"type:varchar(64);"`                // 订单编号
	CreateTime            string `json:"createTime" gorm:"type:datetime;"`                // 提交时间
	MemberUsername        string `json:"memberUsername" gorm:"type:varchar(64);"`         // 用户帐号
	TotalAmount           string `json:"totalAmount" gorm:"type:decimal(10,2);"`          // 订单总金额
	PayAmount             string `json:"payAmount" gorm:"type:decimal(10,2);"`            // 应付金额（实际支付金额）
	FreightAmount         string `json:"freightAmount" gorm:"type:decimal(10,2);"`        // 运费金额
	PromotionAmount       string `json:"promotionAmount" gorm:"type:decimal(10,2);"`      // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     string `json:"integrationAmount" gorm:"type:decimal(10,2);"`    // 积分抵扣金额
	CouponAmount          string `json:"couponAmount" gorm:"type:decimal(10,2);"`         // 优惠券抵扣金额
	DiscountAmount        string `json:"discountAmount" gorm:"type:decimal(10,2);"`       // 管理员后台调整订单使用的折扣金额
	PayType               string `json:"payType" gorm:"type:int;"`                        // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            string `json:"sourceType" gorm:"type:int;"`                     // 订单来源：0->PC订单；1->app订单
	Status                string `json:"status" gorm:"type:int;"`                         // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             string `json:"orderType" gorm:"type:int;"`                      // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string `json:"deliveryCompany" gorm:"type:varchar(64);"`        // 物流公司(配送方式)
	DeliverySn            string `json:"deliverySn" gorm:"type:varchar(64);"`             // 物流单号
	AutoConfirmDay        string `json:"autoConfirmDay" gorm:"type:int;"`                 // 自动确认时间（天）
	Integration           string `json:"integration" gorm:"type:int;"`                    // 可以获得的积分
	Growth                string `json:"growth" gorm:"type:int;"`                         // 可以活动的成长值
	PromotionInfo         string `json:"promotionInfo" gorm:"type:varchar(100);"`         // 活动信息
	BillType              string `json:"billType" gorm:"type:int;"`                       // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string `json:"billHeader" gorm:"type:varchar(200);"`            // 发票抬头
	BillContent           string `json:"billContent" gorm:"type:varchar(200);"`           // 发票内容
	BillReceiverPhone     string `json:"billReceiverPhone" gorm:"type:varchar(32);"`      // 收票人电话
	BillReceiverEmail     string `json:"billReceiverEmail" gorm:"type:varchar(64);"`      // 收票人邮箱
	ReceiverName          string `json:"receiverName" gorm:"type:varchar(100);"`          // 收货人姓名
	ReceiverPhone         string `json:"receiverPhone" gorm:"type:varchar(32);"`          // 收货人电话
	ReceiverPostCode      string `json:"receiverPostCode" gorm:"type:varchar(32);"`       // 收货人邮编
	ReceiverProvince      string `json:"receiverProvince" gorm:"type:varchar(32);"`       // 省份/直辖市
	ReceiverCity          string `json:"receiverCity" gorm:"type:varchar(32);"`           // 城市
	ReceiverRegion        string `json:"receiverRegion" gorm:"type:varchar(32);"`         // 区
	ReceiverDetailAddress string `json:"receiverDetailAddress" gorm:"type:varchar(200);"` // 详细地址
	Note                  string `json:"note" gorm:"type:varchar(500);"`                  // 订单备注
	ConfirmStatus         string `json:"confirmStatus" gorm:"type:int;"`                  // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          string `json:"deleteStatus" gorm:"type:int;"`                   // 删除状态：0->未删除；1->已删除
	UseIntegration        string `json:"useIntegration" gorm:"type:int;"`                 // 下单时使用的积分
	PaymentTime           string `json:"paymentTime" gorm:"type:datetime;"`               // 支付时间
	DeliveryTime          string `json:"deliveryTime" gorm:"type:datetime;"`              // 发货时间
	ReceiveTime           string `json:"receiveTime" gorm:"type:datetime;"`               // 确认收货时间
	CommentTime           string `json:"commentTime" gorm:"type:datetime;"`               // 评价时间
	ModifyTime            string `json:"modifyTime" gorm:"type:datetime;"`                // 修改时间
	BaseModel
}

func (OmsOrder) TableName() string {
	return "oms_order"
}
