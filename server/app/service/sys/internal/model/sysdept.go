package model

import (
	_ "time"
)

type SysDept struct {
	DeptId   int       `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	ParentId int       `json:"parentId" gorm:""`                         //上级部门
	DeptPath string    `json:"deptPath" gorm:"size:255;"`                //
	DeptName string    `json:"deptName"  gorm:"size:128;"`               //部门名称
	Sort     int       `json:"sort" gorm:""`                             //排序
	Leader   string    `json:"leader" gorm:"size:128;"`                  //负责人
	Phone    string    `json:"phone" gorm:"size:11;"`                    //手机
	Email    string    `json:"email" gorm:"size:64;"`                    //邮箱
	Status   string    `json:"status" gorm:"size:4;"`                    //状态
	CreateBy string    `json:"createBy" gorm:"size:64;"`
	UpdateBy string    `json:"updateBy" gorm:"size:64;"`
	Children []SysDept `json:"children" gorm:"-"`
	BaseModel
}

func (SysDept) TableName() string {
	return "sys_dept"
}

type DeptLable struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []DeptLable `gorm:"-" json:"children"`
}
