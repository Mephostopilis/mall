package model

type UmsMemberStatisticsInfo struct {
	AttendCount         string `json:"attendCount" gorm:"type:int;"`             // 关注数量
	CollectCommentCount string `json:"collectCommentCount" gorm:"type:int;"`     //
	CollectProductCount string `json:"collectProductCount" gorm:"type:int;"`     //
	CollectSubjectCount string `json:"collectSubjectCount" gorm:"type:int;"`     //
	CollectTopicCount   string `json:"collectTopicCount" gorm:"type:int;"`       //
	CommentCount        string `json:"commentCount" gorm:"type:int;"`            // 评价数
	ConsumeAmount       string `json:"consumeAmount" gorm:"type:decimal(10,2);"` // 累计消费金额
	CouponCount         string `json:"couponCount" gorm:"type:int;"`             // 优惠券数量
	FansCount           string `json:"fansCount" gorm:"type:int;"`               // 粉丝数量
	Id                  int    `json:"id" gorm:"type:bigint;primary_key"`        //
	InviteFriendCount   string `json:"inviteFriendCount" gorm:"type:int;"`       //
	LoginCount          string `json:"loginCount" gorm:"type:int;"`              // 登录次数
	MemberId            string `json:"memberId" gorm:"type:bigint;"`             //
	OrderCount          string `json:"orderCount" gorm:"type:int;"`              // 订单数量
	RecentOrderTime     string `json:"recentOrderTime" gorm:"type:datetime;"`    // 最后一次下订单时间
	ReturnOrderCount    string `json:"returnOrderCount" gorm:"type:int;"`        // 退货数量
	BaseModel
}

func (UmsMemberStatisticsInfo) TableName() string {
	return "ums_member_statistics_info"
}
