package models

import "time"

type SuperUser struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      int64  `json:"user_id" gorm:"primaryKey"`
	Name        string `json:"user_name,omitempty"` //用户名
	Telephone   string `json:"telephone,omitempty"`
	HeaderPhoto string `json:"header_photo,omitempty"`
	IsPass      int
	IsSuper     int
	Password    string
}
