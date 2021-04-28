package model

type PmsProductAttribute struct {
	Id                         uint64 `json:"id" gorm:"type:bigint unsigned;primary_key"`     //
	AppId                      uint64 `json:"appId" gorm:"type:bigint unsigned;"`             //
	ProductAttributeCategoryId uint64 `json:"productAttributeCategoryId" gorm:"type:bigint;"` //
	Name                       string `json:"name" gorm:"type:varchar(64);"`                  //
	SelectType                 int32  `json:"selectType" gorm:"type:int;"`                    // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int32  `json:"inputType" gorm:"type:int;"`                     // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string `json:"inputList" gorm:"type:varchar(255);"`            // 可选值列表，以逗号隔开
	Sort                       int32  `json:"sort" gorm:"type:int;"`                          // 排序字段：最高的可以单独上传图片
	FilterType                 int32  `json:"filterType" gorm:"type:int;"`                    // 分类筛选样式：1->普通；1->颜色
	SearchType                 int32  `json:"searchType" gorm:"type:int;"`                    // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int32  `json:"relatedStatus" gorm:"type:int;"`                 // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              int32  `json:"handAddStatus" gorm:"type:int;"`                 // 是否支持手动新增；0->不支持；1->支持
	Type                       int32  `json:"type" gorm:"type:int;"`                          // 属性的类型；0->规格；1->参数
	CreateBy                   uint64 `json:"createBy" gorm:"type:bigint;"`                   //
	UpdateBy                   uint64 `json:"updateBy" gorm:"type:bigint;"`                   //
	BaseModel
}

func (PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}
