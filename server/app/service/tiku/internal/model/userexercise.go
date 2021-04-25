package model

type UserExercise struct {
	Answer    string `json:"answer" gorm:"type:varchar(255);"`           // 答案
	Approve   string `json:"approve" gorm:"type:text;"`                  // 批准
	CreateBy  string `json:"createBy" gorm:"type:bigint;"`               //
	Id        int    `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	Remark    string `json:"remark" gorm:"type:text;"`                   // 自己标记
	Tid       string `json:"tid" gorm:"type:bigint unsigned;"`           // 题目id
	Uid       string `json:"uid" gorm:"type:bigint unsigned;"`           // uid
	UpdateBy  string `json:"updateBy" gorm:"type:bigint;"`               //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (UserExercise) TableName() string {
	return "tiku_user_exercise"
}
