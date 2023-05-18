package models

import (
	"time"
)

type SystemMessage struct {
	ID                int64 `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	MessageTitle      string `json:"message_title"`
	MessageDetails    string `json:"message_details"`
	MessageDate       string `json:"message_date"` //发布时间
	MessageStatus     int    `json:"message_status"`
	TargetType        int    `json:"target_type"`
	TargetId          int64  `json:"target_id"`
	PosterId          int64  `json:"poster_id"`
	PostName          string `json:"post_name"`
	TargetStatus      int    `json:"target_status"`
	UserId            int64  `json:"user_id"`
	UserName          string `json:"user_name"`
	UserTelephone     string `json:"user_telephone"`
	MessageReadStatus int    `json:"message_read_status" gorm:"default:0;"` //判断是否已读，0：未读；1：已读
}

type OfficialMessage struct {
	ID             int64 `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	MessageTitle   string `json:"message_title"`
	MessageDetails string `json:"message_details"`
	MessageDate    string `json:"message_date"`                     //发布时间
	PosterName     string `json:"poster_name"`                      //发布者名称
	UserList       []User `gorm:"many2many:official_message_user""` //向所有用户推送官方消息
}

type InteractiveMessage struct {
	ID                int64 `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	MessageDate       string `json:"message_date"`                          // 发送消息的日期
	MessageDetails    string `json:"message_details"`                       // 消息的具体内容
	ObjectID          int64  `json:"object_id"`                             // 该对象的id
	ObjectTypeID      int64  `json:"object_type_id"`                        // 用户发送互动消息的对象种类id，1-商品 ，2-任务，3-公益活动，4-一级评论，5-二级评论
	PosterHeadPhoto   string `json:"poster_head_photo"`                     // 发送消息的用户头像
	PosterID          int64  `json:"poster_id"`                             // 发送消息的用户id
	PosterName        string `json:"poster_name"`                           // 发送消息的用户名称
	UserID            string `json:"user_id"`                               // 发送互动消息的目标用户id
	MessageReadStatus int    `json:"message_read_status" gorm:"default:0;"` //判断是否已读，0：未读；1：已读
}

type MessageLength struct {
	SystemMessageLength      int `json:"systemmessage_length"`
	InteractiveMessageLength int `json:"eachmessage_length"`
}
