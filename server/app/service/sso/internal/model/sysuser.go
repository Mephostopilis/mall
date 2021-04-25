package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserName struct {
	Username string `gorm:"size:64" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"size:128" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	UserId uint64 `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName string `gorm:"size:128" json:"nickName"` //昵称
	Phone    string `gorm:"size:11" json:"phone"`     //手机号
	RoleId   int    `gorm:"" json:"roleId"`           //角色编码
	Salt     string `gorm:"size:255" json:"salt"`     //盐
	Avatar   string `gorm:"size:255" json:"avatar"`   //头像
	Sex      string `gorm:"size:255" json:"sex"`      //性别
	Email    string `gorm:"size:128" json:"email"`    //邮箱
	DeptId   int    `gorm:"" json:"deptId"`           //部门编码
	PostId   int    `gorm:"" json:"postId"`           //职位编码
	Remark   string `gorm:"size:255" json:"remark"`   //备注
	Status   string `gorm:"size:4;" json:"status"`    //状态
	CreateBy string `gorm:"size:128" json:"createBy"` //
	UpdateBy string `gorm:"size:128" json:"updateBy"` //
	BaseModel
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

func (SysUser) TableName() string {
	return "ums_user"
}

//加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}
