package models

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	ID            int64 `json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	SubjectItem   int            `json:"subject_item"`
	SubjectId     int64          `json:"subject_id"`
	ReportReason  int            `json:"report_reason"`
	ReportDetails string         `json:"report_details"`
	ReporterId    int64          `json:"reporter_id"` //举报用户id
}
type UserFeedbackPicturesURL struct {
	Url        string
	Id         int64
	FeedbackId int64
}

type UserFeedback struct {
	ID               int64 `gorm:"primaryKey"`
	UserFeedbackID   int64
	FeedbackContent  string                    `json:"feedback_content"` // 反馈内容
	FeedbackPictures []UserFeedbackPicturesURL `json:"feedback_pictures" gorm:"foreignKey:FeedbackId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserFeedbackForm struct {
	ID               int64 `gorm:"primaryKey"`
	UserFeedbackID   int64
	FeedbackContent  string   `json:"feedback_content"`  // 反馈内容
	FeedbackPictures []string `json:"feedback_pictures"` // 反馈的图片url数组

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
