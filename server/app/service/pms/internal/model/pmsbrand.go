package model

type PmsBrand struct {
	AppId               string `json:"appId" gorm:"type:bigint unsigned;"`   //
	BigPic              string `json:"bigPic" gorm:"type:varchar(255);"`     // 专区大图
	BrandStory          string `json:"brandStory" gorm:"type:text;"`         // 品牌故事
	FactoryStatus       string `json:"factoryStatus" gorm:"type:int;"`       // 是否为品牌制造商：0->不是；1->是
	FirstLetter         string `json:"firstLetter" gorm:"type:varchar(8);"`  // 首字母
	Id                  int    `json:"id" gorm:"type:bigint;primary_key"`    //
	Logo                string `json:"logo" gorm:"type:varchar(255);"`       // 品牌logo
	Name                string `json:"name" gorm:"type:varchar(64);"`        //
	ProductCommentCount string `json:"productCommentCount" gorm:"type:int;"` // 产品评论数量
	ProductCount        string `json:"productCount" gorm:"type:int;"`        // 产品数量
	ShowStatus          string `json:"showStatus" gorm:"type:int;"`          //
	Sort                string `json:"sort" gorm:"type:int;"`                //
	BaseModel
}

func (PmsBrand) TableName() string {
	return "pms_brand"
}
