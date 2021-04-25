package model

type SysDictType struct {
	DictId   int    `gorm:"primary_key;AUTO_INCREMENT" json:"dictId"`
	DictName string `gorm:"size:128;" json:"dictName"` //字典名称
	DictType string `gorm:"size:128;" json:"dictType"` //字典类型
	Status   string `gorm:"size:4;" json:"status"`     //状态
	Remark   string `gorm:"size:255;" json:"remark"`   //备注
	CreateBy string `gorm:"size:11;" json:"createBy"`  //创建者
	UpdateBy string `gorm:"size:11;" json:"updateBy"`  //更新者
	Params   string `gorm:"-" json:"params"`           //
	BaseModel
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
