package model

type UmsMemberTask struct {
	Growth       string `json:"growth" gorm:"type:int;"`           // 赠送成长值
	Id           int    `json:"id" gorm:"type:bigint;primary_key"` //
	Intergration string `json:"intergration" gorm:"type:int;"`     // 赠送积分
	Name         string `json:"name" gorm:"type:varchar(100);"`    //
	Type         string `json:"type" gorm:"type:int;"`             // 任务类型：0->新手任务；1->日常任务
	BaseModel
}

func (UmsMemberTask) TableName() string {
	return "ums_member_task"
}
