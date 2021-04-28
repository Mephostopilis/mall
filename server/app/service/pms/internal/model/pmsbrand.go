package model

type PmsBrand struct {
	Id                  uint64 `json:"id" gorm:"type:bigint;primary_key"`    //
	AppId               uint64 `json:"appId" gorm:"type:bigint unsigned;"`   //
	BigPic              string `json:"bigPic" gorm:"type:varchar(255);"`     // 专区大图
	BrandStory          string `json:"brandStory" gorm:"type:text;"`         // 品牌故事
	FactoryStatus       int32  `json:"factoryStatus" gorm:"type:int;"`       // 是否为品牌制造商：0->不是；1->是
	FirstLetter         string `json:"firstLetter" gorm:"type:varchar(8);"`  // 首字母
	Logo                string `json:"logo" gorm:"type:varchar(255);"`       // 品牌logo
	Name                string `json:"name" gorm:"type:varchar(64);"`        //
	ProductCommentCount int32  `json:"productCommentCount" gorm:"type:int;"` // 产品评论数量
	ProductCount        int32  `json:"productCount" gorm:"type:int;"`        // 产品数量
	ShowStatus          int32  `json:"showStatus" gorm:"type:int;"`          //
	Sort                int32  `json:"sort" gorm:"type:int;"`                //
	CreateBy            uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy            uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsBrand) TableName() string {
	return "pms_brand"
}
