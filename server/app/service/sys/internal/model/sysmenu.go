package model

type Menu struct {
	MenuId     int32  `json:"menuId" gorm:"primary_key;AUTO_INCREMENT"`
	MenuName   string `json:"menuName" gorm:"size:128;"`
	Title      string `json:"title" gorm:"size:128;"`
	Icon       string `json:"icon" gorm:"size:128;"`
	Path       string `json:"path" gorm:"size:128;"`
	Paths      string `json:"paths" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Action     string `json:"action" gorm:"size:16;"`
	Permission string `json:"permission" gorm:"size:255;"`
	ParentId   int32  `json:"parentId" gorm:"size:11;"`
	NoCache    bool   `json:"noCache" gorm:"size:8;"`
	Breadcrumb string `json:"breadcrumb" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int32  `json:"sort" gorm:"size:4;"`
	Visible    bool   `json:"visible" gorm:"size:1;"`
	IsFrame    bool   `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	CreateBy   string `json:"createBy" gorm:"size:128;"`
	UpdateBy   string `json:"updateBy" gorm:"size:128;"`
	BaseModel
}

func (Menu) TableName() string {
	return "sys_menu"
}

type MenuLable struct {
	Id       int32       `json:"id" gorm:"-"`
	Label    string      `json:"label" gorm:"-"`
	Children []MenuLable `json:"children" gorm:"-"`
}

// type MS []Menu
