package model

//sys_role_dept
type SysDeptMenu struct {
	Id       int32  `gorm:""`
	DeptId   int32  `gorm:""`
	MenuId   int32  `gorm:""`
	CreateBy string `gorm:"size:128;"`
	UpdateBy string `gorm:"size:128;"`
	BaseModel
}

func (SysDeptMenu) TableName() string {
	return "sys_dept_menu"
}
