package dao

import (
	"context"
	"errors"
	"strconv"
	"xueyigou_demo/cache"
	"xueyigou_demo/db"
	"xueyigou_demo/global"
	"xueyigou_demo/models"

	"gorm.io/gorm"
)

// TODO 完成log报错，Redis过期删除，GetList可获得RedisorMysql两种数据
func Likes_get(Id int64, objecttype int64) int64 {
	switch objecttype {
	case 0:
		countStr, _ := cache.Get(cache.Comment_l1LikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	case 1:
		countStr, _ := cache.Get(cache.Comment_l2LikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	case 2:
		countStr, _ := cache.Get(cache.WorkLikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	case 3:
		countStr, _ := cache.Get(cache.TaskLikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	case 4:
		countStr, _ := cache.Get(cache.WelfareLikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	case 5:
		countStr, _ := cache.Get(cache.EventLikeKey(Id))
		count, _ := strconv.ParseInt(countStr, 10, 64)
		return count
	}
	return 0
}

func LikeInLocal(LikeChan chan models.LikeForm) {
	for {
		form := <-LikeChan
		if form.ActType {
			switch form.ObjectType {
			case 0:
				var comment models.FirstLevelComment
				if err := db.DB.Where("id = ?", form.ObjectId).First(&comment).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&comment).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("CommentL1", err).Info("LikeAddFailure")
					}
				}
			case 1:
				var comment models.SecondLevelComment
				if err := db.DB.Where("id = ?", form.ObjectId).First(&comment).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&comment).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("CommentL2LikeAddFailure", err)
					}
				}
			case 2:
				var good models.Work
				if err := db.DB.Where("id = ?", form.ObjectId).First(&good).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&good).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("GoodLikeAddFailure", err)
					}
				}
			case 3:
				var task models.Task
				if err := db.DB.Where("id = ?", form.ObjectId).First(&task).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&task).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("TaskLikeAddFailure", err)
					}
				}
			case 4:
				var welfare models.Welfare
				if err := db.DB.Where("id = ?", form.ObjectId).First(&welfare).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&welfare).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("WelfareLikeAddFailure", err)
					}
				}
			case 5:
				var event models.Event
				if err := db.DB.Where("event_id = ?", form.ObjectId).First(&event).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
					err = db.DB.Model(&event).Association("LikeUser").Append(&models.User{UserId: form.UserId})
					if err != nil {
						global.Log.WithField("WelfareLikeAddFailure", err)
					}
				}
			}
		} else {
			switch form.ObjectType {
			case 0:
				comment := FindL1CommentById(form.ObjectId)
				err := db.DB.Model(&comment).Association("LikeUser").Delete(&models.User{UserId: form.UserId})
				if err != nil {
					global.Log.WithField("CommentL1LikeDeleteFailure", err)
				}
			case 1:
				comment := FindL2CommentById(form.ObjectId)
				err := db.DB.Model(&comment).Association("LikeUser").Delete(&models.User{UserId: form.UserId})
				if err != nil {
					global.Log.WithField("CommentL2LikeDeleteFailure", err)
				}
			case 2:
				good, _ := models.WorkModel.FindByID(form.ObjectId)
				err := db.DB.Model(&good).Association("LikeUser").Delete(&models.User{UserId: form.UserId})
				if err != nil {
					global.Log.WithField("GoodLikeDeleteFailure", err)
				}
			case 3:
				task := GetTaskById(form.ObjectId)
				err := db.DB.Model(&task).Association("LikeUser").Delete(&models.User{UserId: form.UserId})
				if err != nil {
					global.Log.WithField("TaskLikeDeleteFailure", err)
				}
			case 4:
				welfare := GetWelfareById(form.ObjectId)
				err := db.DB.Model(&welfare).Association("LikeUser").Delete(&models.User{UserId: form.UserId})
				if err != nil {
					panic(err)
					global.Log.WithField("WelfareLikeDeleteFailure", err)
				}
			}
		}
	}
}
func AddCommentL1Like(commentId int64, userid int64) error {
	var comment models.FirstLevelComment
	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if result, _ := cache.GetClient().SIsMember(context.Background(), cache.Comment_l1LikeUserKey(commentId), strconv.Itoa(int(userid))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     userid,
				ObjectId:   commentId,
				ObjectType: 0,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), cache.Comment_l1LikeUserKey(commentId), strconv.Itoa(int(userid)))
			cache.GetClient().Incr(context.Background(), cache.Comment_l1LikeKey(commentId)) //增加点赞数
			return err
		}
	}
	return nil
}
func DeleteCommentL1Like(commentId int64, userid int64) error {
	if result, err := cache.GetClient().SIsMember(context.Background(), cache.Comment_l1LikeUserKey(commentId), strconv.Itoa(int(userid))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     userid,
			ObjectId:   commentId,
			ObjectType: 0,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), cache.Comment_l1LikeUserKey(commentId), strconv.Itoa(int(userid)))
		cache.GetClient().Decr(context.Background(), cache.Comment_l1LikeKey(commentId)) //减少点赞数
		return nil
	} else {
		return err
	}
}

func AddCommentL2Like(commentId int64, userid int64) error {
	var comment models.SecondLevelComment
	if err := db.DB.Where("id = ?", commentId).First(&comment).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if result, _ := cache.GetClient().SIsMember(context.Background(), cache.Comment_l2LikeUserKey(commentId), strconv.Itoa(int(userid))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     userid,
				ObjectId:   commentId,
				ObjectType: 1,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), cache.Comment_l2LikeUserKey(commentId), strconv.Itoa(int(userid)))
			cache.GetClient().Incr(context.Background(), cache.Comment_l2LikeKey(commentId)) //增加点赞数
			return err
		}
	}
	return nil
}
func DeleteCommentL2Like(commentId int64, userid int64) error {
	if result, err := cache.GetClient().SIsMember(context.Background(), cache.Comment_l2LikeUserKey(commentId), strconv.Itoa(int(userid))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     userid,
			ObjectId:   commentId,
			ObjectType: 1,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), cache.Comment_l2LikeUserKey(commentId), strconv.Itoa(int(userid)))
		cache.GetClient().Decr(context.Background(), cache.Comment_l2LikeKey(commentId)) //减少点赞数
		return nil
	} else {
		return err
	}
}
func AddWorkLike(GoodId int64, UserId int64) error {
	var good models.Work
	if err := db.DB.Where("id = ?", GoodId).First(&good).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if result, _ := cache.GetClient().SIsMember(context.Background(), cache.WorkLikeUserKey(GoodId), strconv.Itoa(int(UserId))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     UserId,
				ObjectId:   GoodId,
				ObjectType: 2,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), cache.WorkLikeUserKey(GoodId), strconv.Itoa(int(UserId)))
			cache.GetClient().Incr(context.Background(), cache.WorkLikeKey(GoodId)) //增加点赞数
			return err
		}
	}
	return nil
}
func DeleteWorkLike(GoodId int64, UserId int64) error {
	if result, err := cache.GetClient().SIsMember(context.Background(), cache.WorkLikeUserKey(GoodId), strconv.Itoa(int(UserId))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     UserId,
			ObjectId:   GoodId,
			ObjectType: 2,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), cache.WorkLikeUserKey(GoodId), strconv.Itoa(int(UserId)))
		cache.GetClient().Decr(context.Background(), cache.WorkLikeKey(GoodId)) //减少点赞数
		return nil
	} else {
		return err
	}
}

// 小程序作品点赞接口
func AppletAddWorkLike(workId int64, UserId int64) error {
	var good models.Work
	if err := db.DB.Where("id = ?", workId).First(&good).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&good).Association("LikeUser").Append(&models.User{UserId: UserId})
		if err != nil {
			global.Log.WithField("GoodLikeAddFailure", err)
			return err
		}
	}
	return nil
}
func AppletCancelWorkLike(workId int64, UserId int64) error {
	good, _ := models.WorkModel.FindByID(workId)
	err := db.DB.Model(&good).Association("LikeUser").Delete(&models.User{UserId: UserId})
	if err != nil {
		global.Log.WithField("GoodLikeDeleteFailure", err)
		return err
	}
	return nil
}
func AddTaskLike(TaskId int64, UserId int64) error {
	var task models.Task
	if err := db.DB.Where("id = ?", TaskId).First(&task).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if result, _ := cache.GetClient().SIsMember(context.Background(), cache.TaskLikeUserKey(TaskId), strconv.Itoa(int(UserId))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     UserId,
				ObjectId:   TaskId,
				ObjectType: 3,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), cache.TaskLikeUserKey(TaskId), strconv.Itoa(int(UserId)))
			cache.GetClient().Incr(context.Background(), cache.TaskLikeKey(TaskId)) //增加点赞数
			return err
		}
	}
	return nil
}
func DeleteTaskLike(TaskId int64, UserId int64) error {
	key := cache.TaskLikeUserKey(TaskId)
	if result, err := cache.GetClient().SIsMember(context.Background(), key, strconv.Itoa(int(UserId))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     UserId,
			ObjectId:   TaskId,
			ObjectType: 3,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), key, strconv.Itoa(int(UserId)))
		cache.GetClient().Decr(context.Background(), cache.TaskLikeKey(TaskId)) //减少点赞数
		return nil
	} else {
		return err
	}
}
func AddWelfareLike(WelfareId int64, UserId int64) error {
	var welfare models.Welfare
	if err := db.DB.Where("id = ?", WelfareId).First(&welfare).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		key := cache.WelfareLikeUserKey(WelfareId)
		if result, _ := cache.GetClient().SIsMember(context.Background(), key, strconv.Itoa(int(UserId))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     UserId,
				ObjectId:   WelfareId,
				ObjectType: 4,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), key, strconv.Itoa(int(UserId)))
			cache.GetClient().Incr(context.Background(), cache.WelfareLikeKey(WelfareId)) //增加点赞数
			return err
		}
	}
	return nil
}

func DeleteWelfareLike(welfareId int64, userId int64) error {
	key := cache.WelfareLikeUserKey(welfareId)
	if result, err := cache.GetClient().SIsMember(context.Background(), key, strconv.Itoa(int(userId))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     userId,
			ObjectId:   welfareId,
			ObjectType: 4,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), key, strconv.Itoa(int(userId)))
		cache.GetClient().Decr(context.Background(), cache.WelfareLikeKey(welfareId)) //减少点赞数
		return nil
	} else {
		return err
	}
}
func AddEventLike(EventId int64, UserId int64) error {
	var event models.Event
	if err := db.DB.Where("id = ?", EventId).First(&event).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		key := cache.EventLikeUserKey(EventId)
		if result, _ := cache.GetClient().SIsMember(context.Background(), key, strconv.Itoa(int(UserId))).Result(); !result {
			global.LikeChan <- models.LikeForm{
				UserId:     UserId,
				ObjectId:   EventId,
				ObjectType: 5,
				ActType:    true,
			}
			cache.GetClient().SAdd(context.Background(), key, strconv.Itoa(int(UserId)))
			cache.GetClient().Incr(context.Background(), cache.EventLikeKey(EventId)) //增加点赞数
			return err
		}
	}
	return nil
}
func AddEventsLike(EventId int64, UserId int64) error {
	var user models.User
	event := GetEventById(EventId)
	if err := db.DB.Where("user_id = ?", UserId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("LikeEvent").Append(&event); err != nil {
			panic(err)
		}
	}
	return nil
}
func DeleteEventLike(eventId int64, userId int64) error {
	key := cache.EventLikeUserKey(eventId)
	if result, err := cache.GetClient().SIsMember(context.Background(), key, strconv.Itoa(int(userId))).Result(); result && err == nil {
		global.LikeChan <- models.LikeForm{
			UserId:     userId,
			ObjectId:   eventId,
			ObjectType: 5,
			ActType:    false,
		}
		cache.GetClient().SRem(context.Background(), key, strconv.Itoa(int(userId)))
		cache.GetClient().Decr(context.Background(), cache.EventLikeKey(eventId)) //减少点赞数
		return nil
	} else {
		return err
	}
}
func DeleteEventsLike(EventId int64, UserId int64) error {
	var user models.User
	event := GetEventById(EventId)
	if err := db.DB.Where("user_id = ?", UserId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("LikeEvent").Delete(&event); err != nil {
			panic(err)
		}
	}
	return nil
}
func GetCommentL1Like(userId int64) []int64 {
	var likeList []int64
	if err := db.DB.Raw("select first_level_comment_id from like_comment_l1 where user_user_id = ?", userId).Find(&likeList).Error; err != nil {
		panic(err)
	}
	return likeList
}
func GetCommentL2Like(userId int64) []int64 {
	var likeList []int64
	if err := db.DB.Raw("select second_level_comment_id from like_comment_l2 where user_user_id = ?", userId).Find(&likeList).Error; err != nil {
		panic(err)
	}
	return likeList
}
func GetWorkLike(userId int64) []int64 {
	var likeList []int64
	if err := db.DB.Raw("select work_id from like_work where user_user_id = ?", userId).Find(&likeList).Error; err != nil {
		panic(err)
	}
	return likeList
}

func GetTaskLike(userId int64) []int64 {
	var likeList []int64
	if err := db.DB.Raw("select task_id from like_task where user_user_id = ?", userId).Find(&likeList).Error; err != nil {
		panic(err)
	}
	//db.DB.Model(&models.Task{}).Association("LikeUser").Find(&likeList)
	return likeList
}

func GetWelfareLike(userId int64) []int64 {
	var likeList []int64
	if err := db.DB.Raw("select welfare_id from like_welfare where user_user_id = ?", userId).Find(&likeList).Error; err != nil {
		panic(err)
	}
	return likeList
}
func GetEventLike(userId int64) []int64 {
	var user models.User
	var events []models.Event
	var list []int64
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("LikeEvent").Find(&events); err != nil {
			panic(err)
		}
	}
	for _, event := range events {
		list = append(list, event.EventId)
	}
	return list
}
func GetEventLikeCount(eventId int64) int64 {
	var event models.Event
	var count int64
	if err := db.DB.Where("event_id = ?", eventId).Find(&event).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		count = db.DB.Model(&event).Association("LikeUser").Count()
	}
	return count
}
func GetWelfareLikeCount(welfareId int64) int64 {
	var count int64
	if err := db.DB.Table("like_welfare").Where("welfare_id = ?", welfareId).Count(&count).Error; err != nil {
		panic(err)
	}
	return count
}
func AddWorkLikeCount(workId int64) {
	var work models.Work
	if err := db.DB.Raw("update works set work_like = work_like + 1 where id = ?", workId).Find(&work).Error; err != nil {
		panic(err)
	}
}
func DeleteWorkLikeCount(workId int64) {
	var work models.Work
	if err := db.DB.Raw("update works set work_like = work_like - 1 where id = ?", workId).Find(&work).Error; err != nil {
		panic(err)
	}
}
