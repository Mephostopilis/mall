package model

import (
	"time"
)

//sys_operlog
type SysOperLog struct {
	OperId        int32     `json:"operId" gorm:"primary_key;AUTO_INCREMENT"` //日志编码
	Title         string    `json:"title" gorm:"size:255;"`                   //操作模块
	BusinessType  string    `json:"businessType" gorm:"size:128;"`            //操作类型
	BusinessTypes string    `json:"businessTypes" gorm:"size:128;"`           //操作类型
	Method        string    `json:"method" gorm:"size:128;"`                  //函数
	RequestMethod string    `json:"requestMethod" gorm:"size:128;"`           //请求方式
	OperatorType  string    `json:"operatorType" gorm:"size:128;"`            //操作类型
	OperName      string    `json:"operName" gorm:"size:128;"`                //操作者
	DeptName      string    `json:"deptName" gorm:"size:128;"`                //部门名称
	OperUrl       string    `json:"operUrl" gorm:"size:255;"`                 //访问地址
	OperIp        string    `json:"operIp" gorm:"size:128;"`                  //客户端ip
	OperLocation  string    `json:"operLocation" gorm:"size:128;"`            //访问位置
	OperParam     string    `json:"operParam" gorm:"size:255;"`               //请求参数
	Status        string    `json:"status" gorm:"size:4;"`                    //操作状态
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;"`          //操作时间
	JsonResult    string    `json:"jsonResult" gorm:"size:255;"`              //返回数据
	DataScope     string    `json:"dataScope" gorm:"-"`                       //数据
	Params        string    `json:"params" gorm:"-"`                          //参数
	Remark        string    `json:"remark" gorm:"size:255;"`                  //备注
	LatencyTime   string    `json:"latencyime" gorm:"size:128;"`              //耗时
	UserAgent     string    `json:"userAgent" gorm:"size:255;"`               //ua
	CreateBy      string    `json:"createBy" gorm:"size:128;"`                //创建人
	UpdateBy      string    `json:"updateBy" gorm:"size:128;"`                //更新者
	BaseModel
}

func (SysOperLog) TableName() string {
	return "sys_operlog"
}
