package model

type PmsAlbum struct {
	AppId       string `json:"appId" gorm:"type:bigint unsigned;"`     //
	CoverPic    string `json:"coverPic" gorm:"type:varchar(1000);"`    //
	Description string `json:"description" gorm:"type:varchar(1000);"` //
	Id          int    `json:"id" gorm:"type:bigint;primary_key"`      //
	Name        string `json:"name" gorm:"type:varchar(64);"`          //
	PicCount    string `json:"picCount" gorm:"type:int;"`              //
	Sort        string `json:"sort" gorm:"type:int;"`                  //
	BaseModel
}

func (PmsAlbum) TableName() string {
	return "pms_album"
}
