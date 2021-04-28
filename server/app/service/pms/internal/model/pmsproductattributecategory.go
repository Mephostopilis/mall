package model

type PmsProductAttributeCategory struct {
	Id             uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"` //
	AppId          uint64 `json:"appId" gorm:"type:bigint unsigned;"`         //
	Name           string `json:"name" gorm:"type:varchar(64);"`              //
	AttributeCount int32  `json:"attributeCount" gorm:"type:int;"`            // 属性数量
	ParamCount     int32  `json:"paramCount" gorm:"type:int;"`                // 参数数量
	CreateBy       uint64 `json:"createBy" gorm:"type:bigint;"`               //
	UpdateBy       uint64 `json:"updateBy" gorm:"type:bigint;"`               //
	BaseModel
}

func (PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}
