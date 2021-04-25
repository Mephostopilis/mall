package model

type SysRoleMenu struct {
	RoleId   int32  `gorm:""`
	MenuId   int32  `gorm:""`
	RoleName string `gorm:"size:128)"`
	CreateBy string `gorm:"size:128)"`
	UpdateBy string `gorm:"size:128)"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

type MenuPath struct {
	Path string `json:"path"`
}
