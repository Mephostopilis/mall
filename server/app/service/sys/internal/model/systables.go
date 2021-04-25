package model

type SysTables struct {
	TableId             int32        `gorm:"primary_key;auto_increment" json:"tableId"`    //表编码
	TBName              string       `gorm:"column:table_name;size:255;" json:"tableName"` //表名称
	TableComment        string       `gorm:"size:255;" json:"tableComment"`                //表备注
	ClassName           string       `gorm:"size:255;" json:"className"`                   //类名
	TplCategory         string       `gorm:"size:255;" json:"tplCategory"`                 //模板分类
	PackageName         string       `gorm:"size:255;" json:"packageName"`                 //包名
	ModuleName          string       `gorm:"size:255;" json:"moduleName"`                  //模块名
	BusinessName        string       `gorm:"size:255;" json:"businessName"`                //业务模块
	FunctionName        string       `gorm:"size:255;" json:"functionName"`                //功能名称
	FunctionAuthor      string       `gorm:"size:255;" json:"functionAuthor"`              //功能作者
	PkColumn            string       `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string       `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string       `gorm:"size:255;" json:"pkJsonField"`
	Options             string       `gorm:"size:255;" json:"options"`
	TreeCode            string       `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string       `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string       `gorm:"size:255;" json:"treeName"`
	Tree                bool         `gorm:"size:1;" json:"tree"`
	Crud                bool         `gorm:"size:1;" json:"crud"`
	Remark              string       `gorm:"size:255;" json:"remark"`
	IsLogicalDelete     string       `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool         `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string       `gorm:"size:128;" json:"logicalDeleteColumn"`
	CreateBy            string       `gorm:"size:128;" json:"createBy"`
	UpdateBy            string       `gorm:"size:128;" json:"updateBy"`
	Columns             []SysColumns `gorm:"-" json:"columns"`
	Module              string       `gorm:"-" json:"module"`
	BaseModel
}

func (SysTables) TableName() string {
	return "sys_tables"
}

type Params struct {
	TreeCode       string `gorm:"-" json:"treeCode"`
	TreeParentCode string `gorm:"-" json:"treeParentCode"`
	TreeName       string `gorm:"-" json:"treeName"`
}
