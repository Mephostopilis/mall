package model

type SysRole struct {
	RoleId   int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleName string `json:"roleName" gorm:"size:128;"`                // 角色名称
	Status   string `json:"status" gorm:"size:4;"`                    // 状态
	RoleKey  string `json:"roleKey" gorm:"size:128;"`                 //角色代码
	RoleSort int    `json:"roleSort" gorm:""`                         //角色排序
	Flag     string `json:"flag" gorm:"size:128;"`                    //
	Remark   string `json:"remark" gorm:"size:255;"`                  //备注
	Admin    bool   `json:"admin" gorm:"size:4;"`
	CreateBy string `json:"createBy" gorm:"size:128;"` //
	UpdateBy string `json:"updateBy" gorm:"size:128;"` //
	BaseModel
}

func (SysRole) TableName() string {
	return "ums_role"
}

type MenuIdList struct {
	MenuId int `json:"menuId"`
}

type DeptIdList struct {
	DeptId int `json:"DeptId"`
}
