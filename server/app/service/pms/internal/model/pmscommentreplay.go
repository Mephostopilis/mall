package model

type PmsCommentReplay struct {
	AppId          string `json:"appId" gorm:"type:bigint unsigned;"`       //
	CommentId      string `json:"commentId" gorm:"type:bigint;"`            //
	Content        string `json:"content" gorm:"type:varchar(1000);"`       //
	CreateTime     string `json:"createTime" gorm:"type:datetime;"`         //
	Id             int    `json:"id" gorm:"type:bigint;primary_key"`        //
	MemberIcon     string `json:"memberIcon" gorm:"type:varchar(255);"`     //
	MemberNickName string `json:"memberNickName" gorm:"type:varchar(255);"` //
	Type           string `json:"type" gorm:"type:int;"`                    // 评论人员类型；0->会员；1->管理员
	BaseModel
}

func (PmsCommentReplay) TableName() string {
	return "pms_comment_replay"
}
