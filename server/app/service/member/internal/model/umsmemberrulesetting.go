package model

type UmsMemberRuleSetting struct {
	ConsumePerPoint   string `json:"consumePerPoint" gorm:"type:decimal(10,2);"` // 每消费多少元获取1个点
	ContinueSignDay   string `json:"continueSignDay" gorm:"type:int;"`           // 连续签到天数
	ContinueSignPoint string `json:"continueSignPoint" gorm:"type:int;"`         // 连续签到赠送数量
	Id                int    `json:"id" gorm:"type:bigint;primary_key"`          //
	LowOrderAmount    string `json:"lowOrderAmount" gorm:"type:decimal(10,2);"`  // 最低获取点数的订单金额
	MaxPointPerOrder  string `json:"maxPointPerOrder" gorm:"type:int;"`          // 每笔订单最高获取点数
	Type              string `json:"type" gorm:"type:int;"`                      // 类型：0->积分规则；1->成长值规则
	BaseModel
}

func (UmsMemberRuleSetting) TableName() string {
	return "ums_member_rule_setting"
}
