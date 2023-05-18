package dao

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
)

func AddEvent(event models.Event) error {
	err := db.DB.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteEvent(eventid int64) error {
	var event models.Event
	var p models.Event
	if err := db.DB.Where("event_id = ? ", eventid).Preload(clause.Associations).Find(&p).Error; err != nil {
		return err
	}
	err := db.DB.Model(&p).Association("EventsPictureList").Delete(models.Eventpicurl{EventId: eventid})
	if err != nil {
		return err
	}

	db.DB.Where("event_id = ?", eventid).Unscoped().Delete(&models.Eventpicurl{})

	if err := db.DB.Raw("delete from events where event_id = ?", eventid).Find(&event).Error; err != nil {
		return err
	}

	return nil
}

func GetEvent(eventid int64) (event models.Event, err error) {
	if err := db.DB.Preload(clause.Associations).Where(" event_id = ?", eventid).Find(&event).Error; err != nil {
		panic(err)
	}
	return event, nil
}
func GetEventList() (eventlist []models.Event, err error) {
	if err := db.DB.Find(&eventlist).Error; err != nil {
		panic(err)
	}
	return eventlist, nil
}
func GetEventById(eventId int64) models.Event {
	var event models.Event
	if err := db.DB.Model(&event).Where("event_id = ?", eventId).Find(&event).Error; err != nil {
		panic(err)
	}
	return event
}
func AddEventCollection(eventId int64, userId int64) error {
	var user models.User
	event := GetEventById(eventId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionEvent").Append(&event)
		return err
	}
	return nil
}
func AddEventCollectionCount(eventId int64) error {
	var event models.Event
	if err := db.DB.Raw("update events set event_collection = event_collection + 1 where event_id = ?", eventId).Find(&event).Error; err != nil {
		return err
	}
	return nil
}
func CancelEventCollection(eventId int64, userId int64) error {
	var user models.User
	event := GetEventById(eventId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionEvent").Delete(&event)
		return err
	}
	return nil
}
func CancelEventCollectionCount(eventId int64) error {
	var event models.Event
	if err := db.DB.Raw("update events set event_collection = event_collection - 1 where event_id = ?", eventId).Find(&event).Error; err != nil {
		return err
	}
	return nil
}
func GetCollectionEventsList(userId int64) []models.Event {
	var list []models.Event
	var user models.User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("CollectionEvent").Find(&list); err != nil {
			panic(err)
		}
	}
	return list
}

func GetEventsListBySearch(isOnline int, isSchool int, isBegin int, keyword string, currentPage int, pageSize int) []models.Event {
	var list []models.Event
	var event models.Event
	query := db.DB.Model(&event)
	query = query.Where("events_name LIKE ? OR connector_name LIKE ? OR events_details LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if isOnline != 0 {
		query = query.Where("is_online = ?", isOnline)
	}
	if isSchool != 0 {
		query = query.Where("is_school = ?", isSchool)
	}

	switch isBegin {
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

	err := query.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil {
		panic(err)
	}

	return list
}

func GetEventMembers(event_id int64) []models.EventMember {
	var event models.Event
	var members []models.EventMember
	if err := db.DB.First(&event, event_id).Error; err != nil {
		return nil
	}
	if err := db.DB.Model(&event).Association("EventMembers").Find(&members); err != nil {
		return members
	}
	return members
}

func PostActivitySign(eventid int64, id string, phone string, name string) error {
	var event models.Event
	if err := db.DB.Where("event_id = ?", eventid).First(&event).Error; err != nil {
		return err
	}
	eventMember := models.EventMember{
		UserID:   id,
		Phone:    phone,
		UserName: name,
	}
	event.EventMembers = append(event.EventMembers, eventMember)
	if err := db.DB.Save(&event).Error; err != nil {
		return err
	}
	return nil
}
