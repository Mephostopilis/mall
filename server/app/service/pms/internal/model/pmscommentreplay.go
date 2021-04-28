package model

type PmsCommentReplay struct {
	Id             int    `json:"id" gorm:"type:bigint;primary_key"`        //
	AppId          string `json:"appId" gorm:"type:bigint unsigned;"`       //
	CommentId      string `json:"commentId" gorm:"type:bigint;"`            //
	Content        string `json:"content" gorm:"type:varchar(1000);"`       //
	MemberIcon     string `json:"memberIcon" gorm:"type:varchar(255);"`     //
	MemberNickName string `json:"memberNickName" gorm:"type:varchar(255);"` //
	Type           string `json:"type" gorm:"type:int;"`                    // 评论人员类型；0->会员；1->管理员
	CreateBy       uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy       uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsCommentReplay) TableName() string {
	return "pms_comment_replay"
}
