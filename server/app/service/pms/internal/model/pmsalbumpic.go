package model

type PmsAlbumPic struct {
	Id       int    `json:"id" gorm:"type:bigint;primary_key"`  //
	AlbumId  string `json:"albumId" gorm:"type:bigint;"`        //
	AppId    string `json:"appId" gorm:"type:bigint unsigned;"` //
	Pic      string `json:"pic" gorm:"type:varchar(1000);"`     //
	CreateBy uint64 `gorm:"column:create_by;type:bigint;" json:"createBy"`
	UpdateBy uint64 `gorm:"column:update_By;type:bigint;" json:"updateBy"`
	BaseModel
}

func (PmsAlbumPic) TableName() string {
	return "pms_album_pic"
}
