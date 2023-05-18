package models

import (
	"gorm.io/gorm"
	"time"
)

type WorkMid struct {
	//	gorm.Model
	ID              int64 `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt       `gorm:"index"`
	WorkBedonated   string               `json:"work_bedonated"` // 商品被捐赠对象
	WorkCollect     int64                `json:"work_collect"`   // 商品被收藏数
	WorkDetails     string               `json:"work_details"`   // 商品简介
	WorkIntroduce   string               `json:"work_introduce"`
	WorkDonation    string               `json:"work_donation"` // 商品捐赠对象
	WorkLike        int64                `json:"work_like"`     // 商品点赞量
	WorkName        string               `json:"work_name"`     // 商品名
	WorkPercent     int64                `json:"work_percent"`  // 商品捐赠百分比
	WorkPicture     string               `json:"work_picture"`
	WorkMax         int64                `json:"work_max"`
	WorkMin         int64                `json:"work_min"`
	WorkType        string               `json:"work_type"`                                           // 商品所属专业类型
	WorkSubject     []WorkMidSubjectItem `json:"work_subject" gorm:"foreignKey:WorkId;references:ID"` // 商品所属类别
	WorkView        int64                `json:"work_view"`                                           // 商品被浏览量
	PostUserId      int64
	PostUser        User `gorm:"foreignKey:PostUserId;references:UserId"` // 发布人
	PicturesUrlFK   uint
	PicturesUrlList []WorkMidPicturesUrl `json:"pictures_url_list" gorm:"foreignKey:WorkId;references:ID"`
	// 图像url
	WorkStatus int `json:"work_status"` //商品状态
	//TODO : WorkForward uint
}

type WorkMidSubjectItem struct {
	Id     int64
	WorkId int64
	Item   string
}
type WorkMidPicturesUrl struct {
	Id     int64
	WorkId int64
	Url    string
}
