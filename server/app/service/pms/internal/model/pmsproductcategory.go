package model

type PmsProductCategory struct {
	Id           uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` //
	AppId        uint64 `json:"appId" gorm:"type:bigint unsigned;"`         //
	ParentId     uint64 `json:"parentId" gorm:"type:bigint;"`               // 上机分类的编号：0表示一级分类
	Name         string `json:"name" gorm:"type:varchar(64);"`              //
	Level        int32  `json:"level" gorm:"type:int;"`                     // 分类级别：0->1级；1->2级
	ProductCount int32  `json:"productCount" gorm:"type:int;"`              //
	ProductUnit  string `json:"productUnit" gorm:"type:varchar(64);"`       //
	NavStatus    int32  `json:"navStatus" gorm:"type:int;"`                 // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int32  `json:"showStatus" gorm:"type:int;"`                // 显示状态：0->不显示；1->显示
	Sort         int32  `json:"sort" gorm:"type:int;"`                      //
	Icon         string `json:"icon" gorm:"type:varchar(255);"`             // 图标
	Keywords     string `json:"keywords" gorm:"type:varchar(255);"`         //
	Description  string `json:"description" gorm:"type:text;"`              // 描述
	CreateBy     uint64 `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy     uint64 `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (PmsProductCategory) TableName() string {
	return "pms_product_category"
}
