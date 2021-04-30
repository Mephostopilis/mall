package model

type MemberAssetsDetail struct {
	Id       uint64  `json:"id" gorm:"type:bigint unsigned;primary_key"` // 主键
	MemberId uint64  `json:"memberId" gorm:"type:bigint unsigned;"`      // 会员ID
	Amount   float32 `json:"amount" gorm:"type:decimal(38,10);"`         // 变动金额
	Balance  float32 `json:"balance" gorm:"type:decimal(38,10);"`        // 变动后金额
	BillNo   string  `json:"billNo" gorm:"type:varchar(50);"`            // 流水号
	Time     string  `json:"time" gorm:"type:varchar(12);"`              // 日期
	Type     int32   `json:"type" gorm:"type:int;"`                      // 类型（1充值，2商品购买）
	CreateBy uint64  `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy uint64  `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (MemberAssetsDetail) TableName() string {
	return "member_assets_detail"
}
