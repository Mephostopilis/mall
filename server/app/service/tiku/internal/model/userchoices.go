package model

type UserChoices struct {
	Id        uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	Answer    string `json:"answer" gorm:"type:varchar(255);"`           // 答案
	Approve   string `json:"approve" gorm:"type:text;"`                  // 批准
	Remark    string `json:"remark" gorm:"type:text;"`                   // 自己标记
	Tid       uint64 `json:"tid" gorm:"type:bigint unsigned;"`           // 题目id
	Uid       uint64 `json:"uid" gorm:"type:bigint unsigned;"`           // uid
	CreateBy  string `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy  string `json:"updateBy" gorm:"type:bigint;"`               //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (UserChoices) TableName() string {
	return "tiku_user_choices"
}
