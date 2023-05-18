package dao

//import (
//	"errors"
//	"xueyigou_demo/global"
//	"xueyigou_demo/model"
//
//	"gorm.io/gorm"
//)
//
//func GetWorkById(good_id int64) model.Work {
//	var good model.Work
//	if err := DB.Select("work_collect", "work_details",
//		"id", "work_like", "work_name", "work_picture",
//		"work_max", "work_min", "work_view",
//		"post_user_id", "work_type", "work_status", "work_introduce").Where("id = ?", good_id).Preload("PostUser").Preload("PicturesUrlList").Preload("WorkSubject").Find(&good).Error; err != nil {
//		panic(err)
//	}
//	good.WorkLike = Likes_get(good_id, 2)
//	return good
//}
//
//func GetWorkMidById(good_id int64) model.WorkMid {
//	var good model.WorkMid
//	if err := DB.Select("work_collect", "work_details",
//		"id", "work_like", "work_name", "work_picture",
//		"work_max", "work_min", "work_view",
//		"post_user_id", "work_type", "work_status", "work_introduce").Where("id = ?", good_id).Preload("PostUser").Preload("PicturesUrlList").Preload("WorkSubject").Find(&good).Error; err != nil {
//		panic(err)
//	}
//	good.WorkLike = Likes_get(good_id, 2)
//	return good
//}
//
//func GetUserAllWorks(user_id int64, work_id int64) []model.Work {
//	var works []model.Work
//	if err := DB.Select("id", "work_name", "work_max", "work_min", "work_picture").Preload("PicturesUrlList").Where("post_user_id = ?", user_id).Find(&works).Error; err != nil {
//		panic(err)
//	}
//	for index, work := range works {
//		if work.ID == work_id {
//			works = append(works[:index], works[index+1:]...)
//		}
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works
//}
//
//func GetWorkInfo(good_id int64) (model.Work, []model.Work) {
//	work := GetWorkById(good_id)
//	works := GetUserAllWorks(work.PostUserId, good_id)
//	return work, works
//}
//
//func GetWorkMidList() []model.WorkMid {
//	var work []model.WorkMid
//	if err := DB.Debug().Preload("WorkSubject").Find(&work).Error; err != nil {
//		panic(err)
//	}
//	return work
//}
//
//func GetWorkMidInfo(workId int64) model.WorkMid {
//	var work model.WorkMid
//	if err := DB.Debug().Where("id = ?", workId).Preload("WorkSubject").Preload("PostUser").Preload("PicturesUrlList").Find(&work).Error; err != nil {
//		panic(err)
//	}
//	return work
//}
//
//func GetWorkSliceListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]model.Work, int64) {
//	var works []model.Work
//	var count int64
//	if searchChoiceFirst == "全部" {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Count(&count).Error; err != nil {
//			panic(err)
//		}
//	} else {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works, count
//}
//func GetWorkSliceListByLike(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]model.Work, int64) {
//	var works []model.Work
//	var count int64
//	if searchChoiceFirst == "全部" {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("work_like desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Count(&count).Error; err != nil {
//			panic(err)
//		}
//	} else {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("work_like desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works, count
//}
//func GetWorkSliceListWithSearchByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]model.Work, int64) {
//	var works []model.Work
//	var count int64
//	if searchChoiceFirst == "全部" {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	} else {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works, count
//}
//func GetWorkSliceListWithSearchByLike(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]model.Work, int64) {
//	var works []model.Work
//	var count int64
//	if searchChoiceFirst == "全部" {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("work_like desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	} else {
//		if err := DB.Debug().Select("work_collect", "work_introduce", "id",
//			"work_like", "work_name", "work_picture", "work_max",
//			"work_min", "work_view", "post_user_id", "work_type", "work_status").Preload("WorkSubject").Preload("PostUser").
//			Preload("PicturesUrlList").Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("work_like desc").Find(&works).Error; err != nil {
//			panic(err)
//		}
//		if err := DB.Debug().Model(&model.Work{}).Where("work_details like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
//			Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
//			panic(err)
//		}
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works, count
//}
//
//func GetWorksIDListWithSearchByEvent(event_name string, currentPage int, pageSize int) ([]model.WorkSubjectItem, int64) {
//	var worklist []model.WorkSubjectItem
//	var count int64
//
//	if currentPage != 0 && pageSize != 0 {
//		if err := DB.Debug().Select("work_id", "item").Where("item = ?", event_name).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&worklist).Count(&count).Error; err != nil {
//			panic(err)
//		}
//	} else {
//		if err := DB.Debug().Select("work_id", "item").Where("item = ?", event_name).Find(&worklist).Count(&count).Error; err != nil {
//			panic(err)
//		}
//	}
//
//	return worklist, count
//}
//
//func AddWork(good model.WorkMid) {
//	if err := DB.Debug().Create(&good).Error; err != nil {
//		panic(err)
//	}
//}
//
//func AddWorkMid(work model.Work) {
//	if err := DB.Debug().Create(&work).Error; err != nil {
//		panic(err)
//	}
//}
//
//func GetWorkListForMe(user_id int64) []model.Work {
//	var works []model.Work
//	if err := DB.Debug().Select("work_collect", "work_details", "id",
//		"work_like", "work_name", "work_picture", "work_max", "work_min", "work_view", "work_type").Where("post_user_id = ?",
//		user_id).Preload("WorkSubject").Find(&works).Error; err != nil {
//		panic(err)
//	}
//	for i := range works {
//		works[i].WorkLike = Likes_get(works[i].ID, 2)
//	}
//	return works
//}
//
//func AddWorkPictureUrl(work_id int64, urls []model.WorkPicturesUrl) error {
//	var work model.Work
//	if err := DB.Debug().Where("id = ?", work_id).First(&work).Error; err == nil {
//		if err := DB.Debug().Model(&work).Association("PicturesUrlList").Append(urls); err != nil {
//			return err
//		}
//	} else {
//		return err
//	}
//	return nil
//}
//
//func AddWorkCollection(workId int64, userId int64) error {
//	var user model.User
//	work := GetWorkById(int64(workId))
//	if err := DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
//		err = DB.Debug().Model(&user).Association("CollectionWork").Append(&work)
//		return err
//	}
//	return nil
//}
//
//func CancelWorkCollection(workId int64, userId int64) error {
//	var user model.User
//	work := GetWorkById(int64(workId))
//	if err := DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
//		err = DB.Debug().Model(&user).Association("CollectionWork").Delete(&work)
//		return err
//	}
//	return nil
//}
//
//func GetCollectionWorkList(userId int64) []model.Work {
//	var works []model.Work
//	var user model.User
//	if err := DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
//		if err = DB.Debug().Model(&user).Association("CollectionWork").Find(&works); err != nil {
//			panic(err)
//		}
//	}
//	for index, work := range works {
//		var w model.Work
//		if err := DB.Where("id = ?", work.ID).Preload("WorkSubject").Find(&w).Error; err != nil {
//			panic(err)
//		}
//		works[index].WorkSubject = w.WorkSubject
//		works[index].PostUser = user
//	}
//	return works
//}
//
//// @return	err	error
//// @return	access_url []model.WorkPicturesUrl url供fs删除
//func DeleteWork(workId int64) ([]model.WorkPicturesUrl, error) {
//	var work model.Work
//	var picture model.WorkPicturesUrl
//	// TODO 删除关联 找出url供fs删除
//	var pictures []model.WorkPicturesUrl
//	// if err := DB.Where("work_id = ? ", workId).Find(&pictures).Error; err != nil {
//	// 	return nil, err
//	// }
//	var w model.Work
//	if err := DB.Select("work_picture, id").Where("id = ? ",
//		workId).Preload("PicturesUrlList").Preload("CollectUser").Preload("LikeUser").Preload("OrderUser").Find(&w).Error; err != nil {
//		return nil, err
//	}
//	pictures = w.PicturesUrlList
//	flag := false
//	for _, v := range global.WorkPhotourls {
//		if v == w.CoverPicture {
//			flag = true
//			break
//		}
//	}
//	if !flag {
//		pictures = append(pictures, model.WorkPicturesUrl{Url: w.CoverPicture})
//	}
//	DB.Debug().Model(&work).Association("PicturesUrlList").Delete(pictures)
//	DB.Debug().Model(&work).Association("WorkSubject").Delete(model.WorkSubjectItem{WorkId: workId})
//	err := DB.Select("LikeUser").Where("id = ?", workId).Delete(&w).Error
//	global.Log.WithError(err).Info("del task")
//	//
//
//	// TODO 删除subject表
//	DB.Where("work_id = ?", workId).Unscoped().Delete(&model.WorkSubjectItem{})
//	//
//
//	for _, user := range w.CollectUser {
//		if err := CancelWorkCollection(workId, user.UserId); err != nil {
//			return nil, err
//		}
//	}
//	for _, user := range w.LikeUser {
//		if err := DeleteWorkLike(workId, user.UserId); err != nil {
//			return nil, err
//		}
//	}
//	for _, user := range w.OrderUser {
//		if err := DelWorkItem(user.UserId, workId); err != nil {
//			return nil, err
//		}
//	}
//	if err := DB.Debug().Raw("delete from work_pictures_urls where work_id = ?", workId).Find(&picture).Error; err == nil {
//		if err = DB.Debug().Raw("delete from works where id = ?", workId).Find(&work).Error; err != nil {
//			return nil, err
//		}
//	} else {
//		return nil, err
//	}
//	return pictures, nil
//}
//
//func DeleteWorkMid(workId int64) ([]model.WorkMidPicturesUrl, error) {
//	var work model.WorkMid
//	var picture model.WorkMidPicturesUrl
//	// TODO 删除关联 找出url供fs删除
//	var pictures []model.WorkMidPicturesUrl
//	// if err := DB.Where("work_id = ? ", workId).Find(&pictures).Error; err != nil {
//	// 	return nil, err
//	// }
//	var w model.WorkMid
//	if err := DB.Select("work_picture, id").Where("id = ? ",
//		workId).Preload("PicturesUrlList").Find(&w).Error; err != nil {
//		return nil, err
//	}
//	pictures = w.PicturesUrlList
//	flag := false
//	for _, v := range global.WorkPhotourls {
//		if v == w.CoverPicture {
//			flag = true
//			break
//		}
//	}
//	if !flag {
//		pictures = append(pictures, model.WorkMidPicturesUrl{Url: w.CoverPicture})
//	}
//	DB.Debug().Model(&work).Association("PicturesUrlList").Delete(pictures)
//	DB.Debug().Model(&work).Association("WorkSubject").Delete(model.WorkMidSubjectItem{WorkId: workId})
//	//
//
//	// TODO 删除subject表
//	DB.Where("work_id = ?", workId).Unscoped().Delete(&model.WorkMidSubjectItem{})
//	//
//
//	if err := DB.Debug().Raw("delete from work_mid_pictures_urls where work_id = ?", workId).Find(&picture).Error; err == nil {
//		if err = DB.Debug().Raw("delete from work_mids where id = ?", workId).Find(&work).Error; err != nil {
//			return nil, err
//		}
//	} else {
//		return nil, err
//	}
//	return pictures, nil
//}
//
//func WorkChangeStatus(Id int64, status int) {
//	if err := DB.Debug().Table("works").Where("id = ?", Id).Update("work_status", status).Error; err != nil {
//		panic(err)
//	}
//}
//
//func PostOrderWork(userId int64, workId int64) error {
//	user, err := GetUserById(userId)
//	if err != nil {
//		return err
//	}
//	var works []model.Work
//	if err := DB.Model(&user).Association("OrderWork").Find(&works); err != nil {
//		return err
//	}
//	for _, work := range works {
//		if work.ID == workId {
//			return errors.New("have exist")
//		}
//	}
//	updata := map[string]interface{}{"work_status": 1}
//	if err := DB.Table("works").Where("id = ?", workId).Updates(updata).Error; err != nil {
//		return err
//	}
//	if err := DB.Model(&user).Association("OrderWork").Append(&model.Work{ID: workId}); err != nil {
//		return nil
//	}
//	return nil
//}
//
//func DelWorkItem(uId int64, oId int64) error {
//	user, err := GetUserById(uId)
//	if err != nil {
//		return err
//	}
//	updata := map[string]interface{}{"work_status": 0}
//	if err := DB.Table("works").Where("id = ?", oId).Updates(updata).Error; err != nil {
//		return err
//	}
//	if err := DB.Model(&user).Association("OrderWork").Delete(&model.Work{ID: oId}); err != nil {
//		return nil
//	}
//	return nil
//}
