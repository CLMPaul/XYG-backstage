package models

type DefaultPicture struct {
	ObjectType string                 `gorm:"primaryKey;size:255"`
	PicMap     []DefaultPictureId2Url `gorm:"foreignKey:ObjectType;references:ObjectType"`
}

type DefaultPictureId2Url struct {
	Id         int64
	Url        string
	ObjectType string
}
