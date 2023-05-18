package models

import (
	"gorm.io/gorm"
	"time"
)

type PeopleSubjectItem struct {
	Id       int64
	PeopleId int64
	Item     string
}
type People struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt      `gorm:"index"`
	PeopleId         int64               `json:"people_id" gorm:"primaryKey"`
	PeopleIntroduce  string              `json:"people_introduce"`
	PeopleIp         string              `json:"people_ip"`
	PeopleDetails    string              `json:"people_details"`    // 人才具体说明
	PeopleMajorskill string              `json:"people_majorskill"` // 人才主要技能
	PeopleSubject    []PeopleSubjectItem `json:"people_subject"`    // 人才具体擅长标签
	PeopleType       string              `json:"people_type"`       // 人才专业类型
	PeopleMax        int64               `json:"people_max"`        //人才能接受的最高报价
	PeopleMin        int64               `json:"people_min"`        //人才能接受的最低报价
	PeopleStatus     int                 `json:"people_status"`     // 人才状态，0-展示，1-隐藏
}
