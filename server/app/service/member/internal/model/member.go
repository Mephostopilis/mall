package model

type Member struct {
	Id                 uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` //
	MemberLevelId      uint64 `json:"memberLevelId" gorm:"type:bigint unsigned;"` //
	Growth             int32  `json:"growth" gorm:"type:int;"`                    // 成长值
	HistoryIntegration int32  `json:"historyIntegration" gorm:"type:int;"`        // 历史积分数量
	Integration        int32  `json:"integration" gorm:"type:int;"`               // 积分
	LuckeyCount        int32  `json:"luckeyCount" gorm:"type:int;"`               // 剩余抽奖次数
	SourceType         int32  `json:"sourceType" gorm:"type:int;"`                // 用户来源
	CreateBy           string `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy           string `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (Member) TableName() string {
	return "member"
}
