package model

type PmsProductCategoryAttributeRelation struct {
	AppId              string `json:"appId" gorm:"type:bigint unsigned;"`     //
	Id                 int    `json:"id" gorm:"type:bigint;primary_key"`      //
	ProductAttributeId string `json:"productAttributeId" gorm:"type:bigint;"` //
	ProductCategoryId  string `json:"productCategoryId" gorm:"type:bigint;"`  //
	BaseModel
}

func (PmsProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}
