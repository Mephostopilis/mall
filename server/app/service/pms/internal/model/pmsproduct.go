package model

import "time"

type PmsProduct struct {
	Id                         uint64    `json:"id" gorm:"type:bigint unsigned;primary_key"`     //
	AppId                      uint64    `json:"appId" gorm:"type:bigint unsigned;"`             //
	BrandId                    uint64    `json:"brandId" gorm:"type:bigint;"`                    //
	ProductCategoryId          uint64    `json:"productCategoryId" gorm:"type:bigint;"`          //
	FeightTemplateId           uint64    `json:"feightTemplateId" gorm:"type:bigint;"`           //
	ProductAttributeCategoryId uint64    `json:"productAttributeCategoryId" gorm:"type:bigint;"` //
	Name                       string    `json:"name" gorm:"type:varchar(64);"`                  //
	Pic                        string    `json:"pic" gorm:"type:varchar(255);"`                  //
	ProductSn                  string    `json:"productSn" gorm:"type:varchar(64);"`             // 货号
	DeleteStatus               int32     `json:"deleteStatus" gorm:"type:int;"`                  // 删除状态：0->未删除；1->已删除
	PublishStatus              int32     `json:"publishStatus" gorm:"type:int;"`                 // 上架状态：0->下架；1->上架
	NewStatus                  int32     `json:"newStatus" gorm:"type:int;"`                     // 新品状态:0->不是新品；1->新品
	RecommandStatus            int32     `json:"recommandStatus" gorm:"type:int;"`               // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int32     `json:"verifyStatus" gorm:"type:int;"`                  // 审核状态：0->未审核；1->审核通过
	Sort                       int32     `json:"sort" gorm:"type:int;"`                          // 排序
	Sale                       int32     `json:"sale" gorm:"type:int;"`                          // 销量
	Price                      float32   `json:"price" gorm:"type:decimal(10,2);"`               //
	PromotionPrice             float32   `json:"promotionPrice" gorm:"type:decimal(10,2);"`      // 促销价格
	GiftGrowth                 int32     `json:"giftGrowth" gorm:"type:int;"`                    // 赠送的成长值
	GiftPoint                  int32     `json:"giftPoint" gorm:"type:int;"`                     // 赠送的积分
	UsePointLimit              int32     `json:"usePointLimit" gorm:"type:int;"`                 // 限制使用的积分数
	SubTitle                   string    `json:"subTitle" gorm:"type:varchar(255);"`             // 副标题
	Description                string    `json:"description" gorm:"type:text;"`                  // 商品描述
	OriginalPrice              float32   `json:"originalPrice" gorm:"type:decimal(10,2);"`       // 市场价
	Stock                      int32     `json:"stock" gorm:"type:int;"`                         // 库存
	LowStock                   int32     `json:"lowStock" gorm:"type:int;"`                      // 库存预警值
	Unit                       string    `json:"unit" gorm:"type:varchar(16);"`                  // 单位
	Weight                     float32   `json:"weight" gorm:"type:decimal(10,2);"`              // 商品重量，默认为克
	PreviewStatus              int32     `json:"previewStatus" gorm:"type:int;"`                 // 是否为预告商品：0->不是；1->是
	ServiceIds                 string    `json:"serviceIds" gorm:"type:varchar(64);"`            // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string    `json:"keywords" gorm:"type:varchar(255);"`             //
	Note                       string    `json:"note" gorm:"type:varchar(255);"`                 //
	AlbumPics                  string    `json:"albumPics" gorm:"type:varchar(255);"`            // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string    `json:"detailTitle" gorm:"type:varchar(255);"`          //
	DetailDesc                 string    `json:"detailDesc" gorm:"type:text;"`                   //
	DetailHtml                 string    `json:"detailHtml" gorm:"type:text;"`                   // 产品详情网页内容
	DetailMobileHtml           string    `json:"detailMobileHtml" gorm:"type:text;"`             // 移动端网页详情
	PromotionStartTime         time.Time `json:"promotionStartTime" gorm:"type:datetime;"`       // 促销开始时间
	PromotionEndTime           time.Time `json:"promotionEndTime" gorm:"type:datetime;"`         // 促销结束时间
	PromotionPerLimit          int32     `json:"promotionPerLimit" gorm:"type:int;"`             // 活动限购数量
	PromotionType              int32     `json:"promotionType" gorm:"type:int;"`                 // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string    `json:"brandName" gorm:"type:varchar(255);"`            // 品牌名称
	ProductCategoryName        string    `json:"productCategoryName" gorm:"type:varchar(255);"`  // 商品分类名称
	CreateBy                   string    `json:"createBy" gorm:"type:bigint;"`                   //
	UpdateBy                   string    `json:"updateBy" gorm:"type:bigint;"`                   //
	BaseModel
}

func (PmsProduct) TableName() string {
	return "pms_product"
}
