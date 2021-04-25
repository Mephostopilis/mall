package model

type PmsComment struct {
	AppId            string `json:"appId" gorm:"type:bigint unsigned;"`         //
	CollectCouont    string `json:"collectCouont" gorm:"type:int;"`             //
	Content          string `json:"content" gorm:"type:text;"`                  //
	CreateTime       string `json:"createTime" gorm:"type:datetime;"`           //
	Id               int    `json:"id" gorm:"type:bigint;primary_key"`          //
	MemberIcon       string `json:"memberIcon" gorm:"type:varchar(255);"`       // 评论用户头像
	MemberIp         string `json:"memberIp" gorm:"type:varchar(64);"`          // 评价的ip
	MemberNickName   string `json:"memberNickName" gorm:"type:varchar(255);"`   //
	Pics             string `json:"pics" gorm:"type:varchar(1000);"`            // 上传图片地址，以逗号隔开
	ProductAttribute string `json:"productAttribute" gorm:"type:varchar(255);"` // 购买时的商品属性
	ProductId        string `json:"productId" gorm:"type:bigint;"`              //
	ProductName      string `json:"productName" gorm:"type:varchar(255);"`      //
	ReadCount        string `json:"readCount" gorm:"type:int;"`                 //
	ReplayCount      string `json:"replayCount" gorm:"type:int;"`               //
	ShowStatus       string `json:"showStatus" gorm:"type:int;"`                //
	Star             string `json:"star" gorm:"type:int;"`                      // 评价星数：0->5
	BaseModel
}

func (PmsComment) TableName() string {
	return "pms_comment"
}
