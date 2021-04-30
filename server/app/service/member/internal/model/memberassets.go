package model

type MemberAssets struct {
	MemberId uint64  `json:"memberId" gorm:"type:bigint;primary_key"` // 会员ID
	Balance  float32 `json:"balance" gorm:"type:decimal(38,10);"`     // 余额
	Frozen   float32 `json:"frozen" gorm:"type:decimal(38,10);"`      // 冻结
	CreateBy uint64  `json:"createBy" gorm:"type:bigint;"`            //
	UpdateBy uint64  `json:"updateBy" gorm:"type:bigint;"`            //
	BaseModel
}

func (MemberAssets) TableName() string {
	return "member_assets"
}
