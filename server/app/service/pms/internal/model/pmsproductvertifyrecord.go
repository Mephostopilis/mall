package model

type PmsProductVertifyRecord struct {
	AppId      string `json:"appId" gorm:"type:bigint unsigned;"`  //
	CreateTime string `json:"createTime" gorm:"type:datetime;"`    //
	Detail     string `json:"detail" gorm:"type:varchar(255);"`    // 反馈详情
	Id         int    `json:"id" gorm:"type:bigint;primary_key"`   //
	ProductId  string `json:"productId" gorm:"type:bigint;"`       //
	Status     string `json:"status" gorm:"type:int;"`             //
	VertifyMan string `json:"vertifyMan" gorm:"type:varchar(64);"` // 审核人
	BaseModel
}

func (PmsProductVertifyRecord) TableName() string {
	return "pms_product_vertify_record"
}
