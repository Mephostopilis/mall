package model

//sys_role_dept
type SysPostMenu struct {
	Id       int32  `gorm:""`
	PostId   int32  `gorm:""`
	MenuId   int32  `gorm:""`
	CreateBy string `gorm:"size:128;"`
	UpdateBy string `gorm:"size:128;"`
	BaseModel
}

func (SysPostMenu) TableName() string {
	return "sys_post_menu"
}
