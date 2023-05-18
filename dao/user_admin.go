package dao

import (
	"xueyigou_demo/db"
	"xueyigou_demo/models"

	"gorm.io/gorm/clause"
)

func AdminRegister(user *models.SuperUser) error {
	err := db.DB.Create(user).Error
	return err
}

func GetAdminIdByPhone(phone string) (int64, int, error) {
	var su models.SuperUser
	err := db.DB.Select("user_id, is_pass").Where("telephone = ?", phone).First(&su).Error
	return su.UserId, su.IsPass, err
}

func GetAdminByPhone(phone string) (*models.SuperUser, error) {
	var su models.SuperUser
	err := db.DB.Where("telephone = ?", phone).First(&su).Error
	return &su, err
}

func UpdateAdminPassword(user *models.SuperUser, password string) error {
	err := db.DB.Model(user).Update("password", password).Error
	return err
}

func GetPendingAccount() ([]models.SuperUser, error) {
	var su []models.SuperUser
	err := db.DB.Where("is_pass = ?", 0).Find(&su).Error
	return su, err
}

func AuditAccount(phone string, is_pass int) error {
	var err error
	if is_pass == 1 {
		err = db.DB.Model(&models.SuperUser{}).Where("telephone = ?", phone).Select("is_pass").Update("is_pass", is_pass).Error
	} else {
		err = db.DB.Where("telephone = ?", phone).Delete(&models.SuperUser{}).Error
	}

	return err
}

func GetAdminByName(name string) (*models.SuperUser, error) {
	var su models.SuperUser
	err := db.DB.Where("name = ?", name).First(&su).Error
	return &su, err
}

func DeletePeopleMid(peopleMidId int64) error {
	var peopleMid models.PeopleMid
	var p models.PeopleMid
	if err := db.DB.Where("people_id = ? ", peopleMidId).Preload(clause.Associations).Find(&p).Error; err != nil {
		return err
	}
	db.DB.Model(&p).Association("PeopleSubject").Delete(models.PeopleMidSubjectItem{PeopleId: peopleMidId})

	// TODO 删除subject表
	db.DB.Where("people_id = ?", peopleMidId).Unscoped().Delete(&models.PeopleMidSubjectItem{})
	//

	if err := db.DB.Raw("delete from people_mids where people_id = ?", peopleMidId).Find(&peopleMid).Error; err != nil {
		return err
	}

	return nil
}

func DeletePeople(peopleId int64) error {
	var people models.People
	var p models.PeopleMid
	if err := db.DB.Where("people_id = ? ", peopleId).Preload(clause.Associations).Find(&p).Error; err != nil {
		return err
	}
	db.DB.Model(&p).Association("PeopleSubject").Delete(models.PeopleSubjectItem{PeopleId: peopleId})

	// TODO 删除subject表
	db.DB.Where("people_id = ?", peopleId).Unscoped().Delete(&models.PeopleSubjectItem{})
	//

	if err := db.DB.Raw("delete from peoples where people_id = ?", peopleId).Find(&people).Error; err != nil {
		return err
	}

	return nil
}
