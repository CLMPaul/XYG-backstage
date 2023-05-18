package applet

import (
	"time"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
)

var ActivityDao activityDao

type activityDao struct{}

func (activityDao) GetEventsByUser(req proto.UserActivityRequest) (list []models.Event, err error) {
	query := db.DB.Model(&models.Event{}).
		//Select("events.event_id", "events_name", "events_details", "events_cover", "events_poster", "events_start_date", "events_end_date").
		Joins("INNER JOIN event_members ON events.event_id = event_members.event_id").
		Where("event_members.user_id = ?", req.UserId)

	if req.Keyword != "" {
		query.Where("events_name LIKE ? OR connector_name LIKE ? OR events_details LIKE ?",
			"%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	if req.IsOnline != 0 {
		query = query.Where("is_online = ?", req.IsOnline)
	}

	if req.IsSchool != 0 {
		query = query.Where("is_school = ?", req.IsSchool)
	}

	// 1-未开始 2-进行中 3-已结束
	switch req.IsBegin {
	case 3:
		now := time.Now().Unix()
		query = query.Where("events_end_date < ?", now)
	case 2:
		now := time.Now().Unix()
		query = query.Where("events_start_date < ?  and  events_end_date > ? ", now, now)
	case 1:
		now := time.Now().Unix()
		query = query.Where("events_start_date > ?", now)
	}

	offset := (req.CurrentPage - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if err = query.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
