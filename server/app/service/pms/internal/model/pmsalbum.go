package model

type PmsAlbum struct {
	Id          uint64 `json:"id" gorm:"type:bigint;primary_key"`      //
	AppId       uint64 `json:"appId" gorm:"type:bigint unsigned;"`     //
	Name        string `json:"name" gorm:"type:varchar(64);"`          //
	CoverPic    string `json:"coverPic" gorm:"type:varchar(1000);"`    //
	PicCount    int32  `json:"picCount" gorm:"type:int;"`              //
	Description string `json:"description" gorm:"type:varchar(1000);"` //
	Sort        int32  `json:"sort" gorm:"type:int;"`                  //
	CreateBy    uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy    uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsAlbum) TableName() string {
	return "pms_album"
}
