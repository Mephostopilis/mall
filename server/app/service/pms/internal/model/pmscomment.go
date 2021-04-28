package model

type PmsComment struct {
	Id               uint64 `json:"id" gorm:"type:bigint;primary_key"`          //
	AppId            string `json:"appId" gorm:"type:bigint unsigned;"`         //
	CollectCouont    int32  `json:"collectCouont" gorm:"type:int;"`             //
	Content          string `json:"content" gorm:"type:text;"`                  //
	MemberIcon       string `json:"memberIcon" gorm:"type:varchar(255);"`       // 评论用户头像
	MemberIp         string `json:"memberIp" gorm:"type:varchar(64);"`          // 评价的ip
	MemberNickName   string `json:"memberNickName" gorm:"type:varchar(255);"`   //
	Pics             string `json:"pics" gorm:"type:varchar(1000);"`            // 上传图片地址，以逗号隔开
	ProductAttribute string `json:"productAttribute" gorm:"type:varchar(255);"` // 购买时的商品属性
	ProductId        uint64 `json:"productId" gorm:"type:bigint;"`              //
	ProductName      string `json:"productName" gorm:"type:varchar(255);"`      //
	ReadCount        int32  `json:"readCount" gorm:"type:int;"`                 //
	ReplayCount      int32  `json:"replayCount" gorm:"type:int;"`               //
	ShowStatus       int32  `json:"showStatus" gorm:"type:int;"`                //
	Star             int32  `json:"star" gorm:"type:int;"`                      // 评价星数：0->5
	CreateBy         uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy         uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsComment) TableName() string {
	return "pms_comment"
}
