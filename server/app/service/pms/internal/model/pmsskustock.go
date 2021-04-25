package model

type PmsSkuStock struct {
	AppId          string `json:"appId" gorm:"type:bigint unsigned;"`        //
	Id             int    `json:"id" gorm:"type:bigint;primary_key"`         //
	LockStock      string `json:"lockStock" gorm:"type:int;"`                // 锁定库存
	LowStock       string `json:"lowStock" gorm:"type:int;"`                 // 预警库存
	Pic            string `json:"pic" gorm:"type:varchar(255);"`             // 展示图片
	Price          string `json:"price" gorm:"type:decimal(10,2);"`          //
	ProductId      string `json:"productId" gorm:"type:bigint;"`             //
	PromotionPrice string `json:"promotionPrice" gorm:"type:decimal(10,2);"` // 单品促销价格
	Sale           string `json:"sale" gorm:"type:int;"`                     // 销量
	SkuCode        string `json:"skuCode" gorm:"type:varchar(64);"`          // sku编码
	SpData         string `json:"spData" gorm:"type:varchar(500);"`          // 商品销售属性，json格式
	Stock          string `json:"stock" gorm:"type:int;"`                    // 库存
	BaseModel
}

func (PmsSkuStock) TableName() string {
	return "pms_sku_stock"
}
