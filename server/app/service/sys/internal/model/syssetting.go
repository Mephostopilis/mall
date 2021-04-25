package model

type SysSetting struct {
	SettingsId int    `json:"settings_id" gorm:"primary_key;AUTO_INCREMENT"`
	Name       string `json:"name" gorm:"type:varchar(256);"`
	Logo       string `json:"logo" gorm:"type:varchar(256);"`
	BaseModel
}

func (SysSetting) TableName() string {
	return "sys_setting"
}

type ResponseSystemConfig struct {
	Name string `json:"name" binding:"required"` // 名称
	Logo string `json:"logo" binding:"required"` // 头像
}
