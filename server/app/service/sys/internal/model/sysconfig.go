package model

import (
	_ "time"
)

type SysConfig struct {
	ConfigId    int32  `json:"configId" gorm:"primary_key;auto_increment;"` //编码
	ConfigName  string `json:"configName" gorm:"size:128;"`                 //参数名称
	ConfigKey   string `json:"configKey" gorm:"size:128;"`                  //参数键名
	ConfigValue string `json:"configValue" gorm:"size:255;"`                //参数键值
	ConfigType  string `json:"configType" gorm:"size:64;"`                  //是否系统内置
	Remark      string `json:"remark" gorm:"size:128;"`                     //备注
	CreateBy    string `json:"createBy" gorm:"size:128;"`
	UpdateBy    string `json:"updateBy" gorm:"size:128;"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
}

func (SysConfig) TableName() string {
	return "sys_config"
}
