package model

type PmsAlbumPic struct {
	AlbumId string `json:"albumId" gorm:"type:bigint;"`        //
	AppId   string `json:"appId" gorm:"type:bigint unsigned;"` //
	Id      int    `json:"id" gorm:"type:bigint;primary_key"`  //
	Pic     string `json:"pic" gorm:"type:varchar(1000);"`     //
	BaseModel
}

func (PmsAlbumPic) TableName() string {
	return "pms_album_pic"
}
