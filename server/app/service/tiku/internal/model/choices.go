package model

type Choices struct {
	Id       uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` // id
	Ty       int32  `json:"ty" gorm:"type:int;"`                        // 类型
	Level    int32  `json:"level" gorm:"type:int;"`                     // 等级
	Title    string `json:"title" gorm:"type:text;"`                    // 标题
	Pics     string `json:"pics" gorm:"type:text;"`                     // 图片
	Answer   string `json:"answer" gorm:"type:varchar(255);"`           // 答案
	CreateBy uint64 `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy uint64 `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (Choices) TableName() string {
	return "tiku_choices"
}
