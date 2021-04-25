package model

type UmsMemberMemberTagRelation struct {
	Id       int    `json:"id" gorm:"type:bigint;primary_key"` //
	MemberId string `json:"memberId" gorm:"type:bigint;"`      //
	TagId    string `json:"tagId" gorm:"type:bigint;"`         //
	BaseModel
}

func (UmsMemberMemberTagRelation) TableName() string {
	return "ums_member_member_tag_relation"
}
