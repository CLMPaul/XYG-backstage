package dao

import (
	"errors"
	"gorm.io/gorm"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
)

func GetMessageList(userId int64) ([]models.SystemMessage, error) {
	var messages []models.SystemMessage
	if err := db.DB.Where("poster_id = ?", userId).Order("created_at desc").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func PostSystemMessage(systemMessage models.SystemMessage) error {
	if err := db.DB.Create(&systemMessage).Error; err != nil {
		return err
	}
	return nil
}
func DeleteSystemMessage(messageId int64) error {
	var message models.SystemMessage
	if err := db.DB.Raw("delete from system_messages where id = ?", messageId).Find(&message).Error; err != nil {
		return err
	}
	return nil
}

func PostOfficialMessage(officialMessage models.OfficialMessage) error {
	if users, err := GetAllUsers(); err != nil {
		return err
	} else {
		officialMessage.UserList = users
		if err = db.DB.Create(&officialMessage).Error; err != nil {
			return err
		}
	}
	return nil
}
func GetOfficialMessageList(userId int64) ([]models.OfficialMessage, error) {
	var messages []models.OfficialMessage
	var user models.User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("OfficialMessageList").Find(&messages)
	}
	var descMesssage []models.OfficialMessage
	lenth := len(messages)
	for i := 0; i < lenth; i++ {
		descMesssage = append(descMesssage, messages[lenth-1-i])
	}
	return descMesssage, nil
}
func DeleteOfficialMessage(messageId int64, userId int64) error {
	var message models.OfficialMessage
	if user, err := GetUserById(userId); err != nil {
		return err
	} else {
		if err = db.DB.Where("id = ?", messageId).Find(&message).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.DB.Model(&message).Association("UserList").Delete(&user)
			return err
		}
	}
	return nil
}
func Getinteractivemessagelist(user_id int64) ([]models.InteractiveMessage, error) {
	var messages []models.InteractiveMessage
	if err := db.DB.Where("user_id  = ?", user_id).Order("created_at desc").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func PostInteractiveMessage(message models.InteractiveMessage) error {
	if err := db.DB.Create(&message).Error; err != nil {
		return err
	}
	return nil
}

func DeleteInteractiveMessage(message_id int64) error {
	if err := db.DB.Where("id = ?", message_id).Delete(&models.InteractiveMessage{}).Error; err != nil {
		return err
	}
	return nil
}
func GetMessageLength(user_id int64) (*models.MessageLength, error) {
	Length := new(models.MessageLength)
	SystemMessage, err := GetMessageList(user_id)
	if err != nil {
		return nil, err
	}
	EachMessage, err := Getinteractivemessagelist(user_id)
	if err != nil {
		return nil, err
	}
	Length.InteractiveMessageLength = len(EachMessage)
	Length.SystemMessageLength = len(SystemMessage)
	return Length, err
}

func MessageChangeStatus(Id int64, status int) {
	if err := db.DB.Table("system_messages").Where("id = ?", Id).Update("message_status", status).Error; err != nil {
		panic(err)
	}
}

func SystemMessageRead(messageId int64) error {
	if err := db.DB.Model(&models.SystemMessage{}).Where("id = ?", messageId).Update("message_read_status", 1).Error; err != nil {
		return err
	}
	return nil
}
func InteractiveMessageRead(messageId int64) error {
	if err := db.DB.Model(&models.InteractiveMessage{}).Where("id = ?", messageId).Update("message_read_status", 1).Error; err != nil {
		return err
	}
	return nil
}
