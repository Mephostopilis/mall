package model

type UmsMemberProductCategoryRelation struct {
	Id                int    `json:"id" gorm:"type:bigint;primary_key"`     //
	MemberId          string `json:"memberId" gorm:"type:bigint;"`          //
	ProductCategoryId string `json:"productCategoryId" gorm:"type:bigint;"` //
	BaseModel
}

func (UmsMemberProductCategoryRelation) TableName() string {
	return "ums_member_product_category_relation"
}
