package model

import (
	"time"
)

type SysLoginLog struct {
	InfoId        int       `json:"infoId" gorm:"primary_key;AUTO_INCREMENT"` //主键
	Username      string    `json:"username" gorm:"size:128;"`                //用户名
	Status        string    `json:"status" gorm:"size:4;"`                    //状态
	Ipaddr        string    `json:"ipaddr" gorm:"size:255;"`                  //ip地址
	LoginLocation string    `json:"loginLocation" gorm:"size:255;"`           //归属地
	Browser       string    `json:"browser" gorm:"size:255;"`                 //浏览器
	Os            string    `json:"os" gorm:"size:255;"`                      //系统
	Platform      string    `json:"platform" gorm:"size:255;"`                // 固件
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;"`         //登录时间
	CreateBy      string    `json:"createBy" gorm:"size:128;"`                //创建人
	UpdateBy      string    `json:"updateBy" gorm:"size:128;"`                //更新者
	DataScope     string    `json:"dataScope" gorm:"-"`                       //数据
	Params        string    `json:"params" gorm:"-"`                          //
	Remark        string    `json:"remark" gorm:"size:255;"`                  //备注
	Msg           string    `json:"msg" gorm:"size:255;"`
	BaseModel
}

func (SysLoginLog) TableName() string {
	return "ums_loginlog"
}
