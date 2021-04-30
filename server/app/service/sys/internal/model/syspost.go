package model

type SysPost struct {
	PostId   int32  `gorm:"primary_key;AUTO_INCREMENT" json:"postId"` //岗位编号
	PostName string `gorm:"size:128;" json:"postName"`                //岗位名称
	PostCode string `gorm:"size:128;" json:"postCode"`                //岗位代码
	Sort     int32  `gorm:"" json:"sort"`                             //岗位排序
	Status   string `gorm:"size:4;" json:"status"`                    //状态
	Remark   string `gorm:"size:255;" json:"remark"`                  //描述
	CreateBy string `gorm:"size:128;" json:"createBy"`
	UpdateBy string `gorm:"size:128;" json:"updateBy"`
	BaseModel
}

func (SysPost) TableName() string {
	return "sys_post"
}
