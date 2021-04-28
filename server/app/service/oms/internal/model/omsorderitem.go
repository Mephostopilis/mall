package model

type OmsOrderItem struct {
	Id                int    `json:"id" gorm:"type:bigint;primary_key"`            //
	OrderId           string `json:"orderId" gorm:"type:bigint;"`                  // 订单id
	OrderSn           string `json:"orderSn" gorm:"type:varchar(64);"`             // 订单编号
	ProductId         string `json:"productId" gorm:"type:bigint;"`                //
	ProductPic        string `json:"productPic" gorm:"type:varchar(500);"`         //
	ProductName       string `json:"productName" gorm:"type:varchar(200);"`        //
	ProductBrand      string `json:"productBrand" gorm:"type:varchar(200);"`       //
	ProductSn         string `json:"productSn" gorm:"type:varchar(64);"`           //
	ProductPrice      string `json:"productPrice" gorm:"type:decimal(10,2);"`      // 销售价格
	ProductQuantity   string `json:"productQuantity" gorm:"type:int;"`             // 购买数量
	ProductSkuId      string `json:"productSkuId" gorm:"type:bigint;"`             // 商品sku编号
	ProductSkuCode    string `json:"productSkuCode" gorm:"type:varchar(50);"`      // 商品sku条码
	ProductCategoryId string `json:"productCategoryId" gorm:"type:bigint;"`        // 商品分类id
	PromotionName     string `json:"promotionName" gorm:"type:varchar(200);"`      // 商品促销名称
	PromotionAmount   string `json:"promotionAmount" gorm:"type:decimal(10,2);"`   // 商品促销分解金额
	CouponAmount      string `json:"couponAmount" gorm:"type:decimal(10,2);"`      // 优惠券优惠分解金额
	IntegrationAmount string `json:"integrationAmount" gorm:"type:decimal(10,2);"` // 积分优惠分解金额
	RealAmount        string `json:"realAmount" gorm:"type:decimal(10,2);"`        // 该商品经过优惠后的分解金额
	GiftIntegration   string `json:"giftIntegration" gorm:"type:int;"`             //
	GiftGrowth        string `json:"giftGrowth" gorm:"type:int;"`                  //
	ProductAttr       string `json:"productAttr" gorm:"type:varchar(500);"`        // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	BaseModel
}

func (OmsOrderItem) TableName() string {
	return "oms_order_item"
}
