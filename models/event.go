package models

import (
	"gorm.io/gorm"
	"time"
)

type Eventform struct {
	EventId            string   `json:"events_id,omitempty"`
	ConnectorID        string   `json:"connector_id"`        // 活动联系人id
	ConnectorName      string   `json:"connector_name"`      // 活动联系人姓名
	ConnectorQrCode    string   `json:"connector_qr_code"`   // 活动二维码图片url
	ConnectorTelephone string   `json:"connector_telephone"` // 活动联系人电话号码
	EventsBeforejoin   string   `json:"events_beforejoin"`   // 活动参赛须知
	EventsCover        string   `json:"events_cover"`        // 活动封面url
	EventsDetails      string   `json:"events_details"`      // 活动详情信息
	EventsEndDate      string   `json:"events_end_date"`     // 活动结束日期
	EventsName         string   `json:"events_name"`         // 活动名称
	EventsPictureList  []string `json:"events_picture_list"` // 活动图片列表
	EventsPoster       string   `json:"events_poster"`       // 活动举办方名称
	EventsQA           string   `json:"events_QA"`           // 活动问答
	EventsRewards      string   `json:"events_rewards"`      // 活动奖品
	EventsRules        string   `json:"events_rules"`        // 活动规则
	EventsStartDate    string   `json:"events_start_date"`   // 活动开始日期
	EventsWorkRequire  string   `json:"events_work_require"` // 活动参赛作品要求
	IsOnline           *int64   `json:"isOnline,omitempty"`  // 是否线上
	IsSchool           *int64   `json:"isSchool,omitempty"`  // 是否校园
	LikeNums           int64
	//LikeUser           []*User `gorm:"many2many:like_event"`
}
type Event struct {
	EventId            int64 `json:"EventId" gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	ConnectorID        string         `json:"connector_id"`        // 活动联系人id
	ConnectorName      string         `json:"connector_name"`      // 活动联系人姓名
	ConnectorQrCode    string         `json:"connector_qr_code"`   // 活动二维码图片url
	ConnectorTelephone string         `json:"connector_telephone"` // 活动联系人电话号码
	EventsBeforejoin   string         `json:"events_beforejoin"`   // 活动参赛须知
	EventsCover        string         `json:"events_cover"`        // 活动封面url
	EventsDetails      string         `json:"events_details"`      // 活动详情信息
	EventsEndDate      string         `json:"events_end_date"`     // 活动结束日期
	EventsName         string         `json:"events_name"`         // 活动名称
	EventsPictureList  []Eventpicurl  `json:"events_picture_list"` // 活动图片列表
	EventsPoster       string         `json:"events_poster"`       // 活动举办方名称
	EventsQA           string         `json:"events_QA"`           // 活动问答
	EventsRewards      string         `json:"events_rewards"`      // 活动奖品
	EventsRules        string         `json:"events_rules"`        // 活动规则
	EventsStartDate    string         `json:"events_start_date"`   // 活动开始日期
	EventsWorkRequire  string         `json:"events_work_require"` // 活动参赛作品要求
	IsOnline           *int64         `json:"isOnline,omitempty"`  // 是否线上
	IsSchool           *int64         `json:"isSchool,omitempty"`  // 是否校园

	EventCollection int64         `json:"event_collection"`
	LikeUser        []*User       `gorm:"many2many:like_event"`
	EventMembers    []EventMember `json:"event_members,omitempty" gorm:"foreignKey:EventID"`
}

type EventMember struct {
	UserName string `json:"username"`
	Avtor    string `json:"avtor"`
	Phone    string `json:"phone"`
	EventID  string `json:"-" gorm:"primaryKey"`
	UserID   string `json:"user_id" gorm:"primaryKey"`
}

type Eventpicurl struct {
	EventId int64
	Id      int64
	Url     string
}

type EventDeleteform struct {
	EventId string `json:"events_id"` // 需要删除的活动id
}

type EventGetForm struct {
	EventsID string `json:"events_id"`
}
