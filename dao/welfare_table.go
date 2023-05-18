package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
)

func GetMyWelfareList(userId int64) []models.Welfare {
	var welfareId []int64
	var welfares []models.Welfare
	if err := db.DB.Table("welfare_members").Select("welfare_id").Where("user_id = ?", userId).Find(&welfareId).Error; err != nil {
		panic(err)
	}
	for _, id := range welfareId {
		var welfare models.Welfare
		if err := db.DB.Where("id = ?", id).Find(&welfare).Error; err != nil {
			panic(err)
		} else {
			welfares = append(welfares, welfare)
		}
	}
	return welfares
}

func GetWelfareInfoForCollection(userId int64) []models.Welfare {
	var user models.User
	var welfares []models.Welfare
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("CollectionWelfare").Find(&welfares); err != nil {
			panic(err)
		}
	}
	return welfares
}

//	func GetWelfareCollection(welfareId uint) int64 {
//		var collection int64
//		if err := db.DB.Raw("select count(welfare_id = ? or null) from collection_welfare", welfareId).Find(&collection).Error; err != nil {
//			panic(err)
//		}
//		return collection
//	}
func GetWelfareNewsList() []models.Welfare {
	var welfares []models.Welfare
	if err := db.DB.Order("created_at desc").Find(&welfares).Error; err != nil {
		panic(err)
	}
	return welfares
}
func GetWelfareById(welfareId int64) models.Welfare {
	var welfare models.Welfare
	if err := db.DB.Where("id = ?", welfareId).Preload("PicturesUrlList").Find(&welfare).Error; err != nil {
		panic(err)
	}
	return welfare
}

func AddWelfareCollection(welfareId int64, userId int64) error {
	var user models.User
	welfare := GetWelfareById(welfareId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionWelfare").Append(&welfare)
		return err
	}
	return nil
}

func AddWelfareCollectionCount(welfareId int64) error {
	var welfare models.Welfare
	if err := db.DB.Raw("update welfares set welfare_collection=welfare_collection+1 where id = ?", welfareId).Find(&welfare).Error; err != nil {
		return err
	}
	return nil
}
func AddTaskCollectionCount(taskId int64) error {
	var task models.Task
	if err := db.DB.Raw("update tasks set task_collection=task_collection+1 where id = ?", taskId).Find(&task).Error; err != nil {
		return err
	}
	return nil
}
func AddWorkCollectionCount(WorkId int64) error {
	var work models.Work
	if err := db.DB.Raw("update works set work_collect=work_collect+1 where id = ?", WorkId).Find(&work).Error; err != nil {
		return err
	}
	return nil
}
func CancelWelfareCount(welfareId int64) error {
	var welfare models.Welfare
	if err := db.DB.Raw("update welfares set welfare_collection=if(welfare_collection>0,welfare_collection-1,welfare_collection) where id = ?", welfareId).Find(&welfare).Error; err != nil {
		return err
	}
	return nil
}
func CancelTaskCollectionCount(taskId int64) error {
	var task models.Task
	if err := db.DB.Raw("update tasks set task_collection=if(task_collection>0,task_collection-1,task_collection) where id = ?", taskId).Find(&task).Error; err != nil {
		return err
	}
	return nil
}
func CancelWorkCollectionCount(WorkId int64) error {
	var work models.Work
	if err := db.DB.Raw("update works set work_collect=if(work_collect>0,work_collect-1,work_collect) where id = ?", WorkId).Find(&work).Error; err != nil {
		return err
	}
	return nil
}
func CancelWelfareCollection(welfareId int64, userId int64) error {
	var user models.User
	welfare := GetWelfareById(welfareId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionWelfare").Delete(&welfare)
		return err
	}
	return nil
}

func GetWelfareActivity(welfareId int64) models.Welfare {
	var welfare models.Welfare
	if err := db.DB.Where("id = ?", welfareId).Preload("PostUser").Find(&welfare).Error; err != nil {
		panic(err)
	}
	return welfare
}

func PostWelfareInfo(welfare models.Welfare) error {
	//if err := db.DB.Raw("insert into welfares (post_user_id,welfare_name,welfare_picture,welfare_details,welfare_info,welfare_join,welfare_status,welfare_address,start_date,end_date,donation_name,user_type,connector_name,connector_telephone,connector_qr_code_photo) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", welfare.PostUserId,
	//	welfare.WelfareName, welfare.WelfarePicture, welfare.WelfareDetails, welfare.WelfareInfo, welfare.WelfareJoin, welfare.WelfareStatus, welfare.WelfareAddress, welfare.StartDate, welfare.EndDate, welfare.DonationName, welfare.UserType, welfare.ConnectorName, welfare.ConnectorTelephone, welfare.ConnectorQrCodePhoto).Error; err != nil {
	//	return err
	//}
	//return nil
	if err := db.DB.Create(&welfare).Error; err != nil {
		fmt.Println(err)
	}
	return nil
}

func GetWelfareCollection(welfareId int64) int64 {
	var count int64
	if err := db.DB.Raw("select count(welfare_id = ? or null) from welfare_collection", welfareId).Find(&count).Error; err != nil {
		panic(err)
	}
	return count
}
func GetWelfarePicturesUrl(welfareId int64) []string {
	var pictures []string
	if err := db.DB.Raw("select url from welfare_picture_urls where welfare_id = ?", welfareId).Find(&pictures).Error; err != nil {
		panic(err)
	}
	return pictures
}
func GetWelfareHistory() (int64, int64) {
	var welfareCount int64
	var welfarePeople int64
	if err := db.DB.Table("welfare_members").Distinct("welfare_id").Count(&welfareCount).Error; err == nil {
		if err = db.DB.Table("welfare_members").Count(&welfarePeople).Error; err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return welfareCount, welfarePeople
}

func JoinWelfare(welfareMember models.WelfareMember) error {
	var member models.WelfareMember
	if err := db.DB.First(&member, models.WelfareMember{WelfareId: welfareMember.WelfareId, UserId: welfareMember.UserId}).Error; err == nil {
		return errors.New("have existed")
	}
	if err := db.DB.Create(&welfareMember).Error; err != nil {
		return err
	}
	return nil
}

func GetWelfarePeople(welfareId int64) []models.Indentity {
	var indentities []models.Indentity
	var members []models.WelfareMember
	if err := db.DB.Table("welfare_members").Where("welfare_id = ?", welfareId).Find(&members).Error; err != nil {
		panic(err)
	}
	for _, member := range members {
		var indentity models.Indentity
		if err := db.DB.Select("user_name", "user_telenum").Where("user_id = ?", member.UserId).
			First(&indentity).Error; err != nil {
			panic(err)
		}
		indentity.UserId = member.UserId
		indentities = append(indentities, indentity)
	}
	return indentities
}

func AddWelfareInfo(userId []int64, welfareId int64, welfareTime int) error {
	tx := db.DB.Begin()
	var rows int64
	result := tx.Table("welfare_members").Where("user_id IN ?", userId).Updates(map[string]interface{}{"status": 1})
	rows = result.RowsAffected + rows
	if result.Error != nil {
		panic(result.Error)
	}
	result = tx.Table("welfares").Where("id = ?", welfareId).Update("welfare_time", welfareTime)
	if result.Error != nil {
		panic(result.Error)
	}
	rows = result.RowsAffected + rows
	result = tx.Table("users").Where("user_id IN ?", userId).Updates(map[string]interface{}{"welfare_time": welfareTime,
		"contribution": gorm.Expr("contribution + ?", welfareTime*10), "medals": gorm.Expr("medals + ?", 1)})
	if result.Error != nil {
		panic(result.Error)
	}
	rows = result.RowsAffected + rows
	if rows != int64(2*len(userId)+1) {
		tx.Rollback()
		return errors.New("illegal")
	}
	tx.Commit()
	return nil
}
func ModifyWelfare(welfare models.Welfare, welfareId int64, urls []models.WelfarePictureUrl) error {
	if err := db.DB.Model(&welfare).Where("id = ?", welfareId).Updates(&welfare).Error; err != nil {
		panic(err)
	}
	var urlsList []models.WelfarePictureUrl
	if err := db.DB.Model(&models.WelfarePictureUrl{}).Where("welfare_id = ?", welfareId).Delete(&urlsList).Error; err == nil {
		if err = db.DB.Create(&urls).Error; err != nil {
			return err
		}
	} else {
		panic(err)
	}
	return nil
}
func DeleteWelfare(welfareId int64) error {
	var urls []models.WelfarePictureUrl
	var welfare models.Welfare
	if err := db.DB.Model(&models.WelfarePictureUrl{}).Where("welfare_id = ?", welfareId).Delete(&urls).Error; err == nil {
		if err = db.DB.Model(&models.Welfare{}).Where("id = ?", welfareId).Preload("PostUser").Preload("LikeUser").
			Preload("AppendUser").Delete(&welfare).Error; err != nil {
			panic(err)
			return err
		}
	} else {
		panic(err)
	}
	return nil
}
