package models

import (
	"gorm.io/gorm"
	"time"
)

type FirstLevelComment struct {
	ID           int64 `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	GoodId       int64
	ObjectType   ObjectType
	CommenterId  int64
	Commenter    User `gorm:"foreignKey:CommenterId;references:UserId"`
	Content      string
	PostDate     string
	LikeUser     []User `gorm:"many2many:like_commentL1"`
	CommentLikes int64
}

type SecondLevelComment struct {
	ID                  int64 `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	GoodId              int64
	ObjectType          ObjectType
	FirstLevelCommentId int64
	ReplierId           int64
	Replier             User `gorm:"foreignKey:ReplierId;references:UserId"`
	ReplytoId           int64
	ReplyTo             User `gorm:"foreignKey:ReplytoId;references:UserId"`
	ReplyToCommentId    string
	Content             string
	ReplyDate           string
	CommentLikes        int64
	LikeUser            []User `gorm:"many2many:like_commentL2"`
}
