package dao

import (
	"errors"
	"strings"
	"xueyigou_demo/db"
	"xueyigou_demo/models"

	"gorm.io/gorm"

	// "errors"

	"xueyigou_demo/global"

	// "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 发布任务
func PublishTask(task *models.TaskMid) int64 {
	if err := db.DB.Omit("oredered_user_id").Create(&task).Error; err != nil {
		panic(err)
	}
	return int64(task.ID)
}

func PublishTaskMid(task *models.Task) int64 {
	if err := db.DB.Omit("oredered_user_id").Create(&task).Error; err != nil {
		panic(err)
	}
	return int64(task.ID)
}

// 通过userid获取任务列表
func GetTaskList(user_id int64) []models.Task {
	var tasks []models.Task
	if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
		"task_max", "task_details", "post_user_id", "task_type", "task_introduce").Where(
		"post_user_id = ?", user_id).Preload(clause.Associations).Find(&tasks).Error; err != nil {
		panic(err)
	}
	for i := range tasks {
		tasks[i].TaskLike = Likes_get(tasks[i].ID, 3)
	}
	return tasks
}

func GetTaskMidList() []models.TaskMid {
	var tasks []models.TaskMid
	if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
		"task_max", "task_introduce", "post_user_id", "task_type", "task_status", "task_details").Preload(clause.Associations).Find(&tasks).Error; err != nil {
		panic(err)
	}
	return tasks
}

func GetTaskSliceListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Task, int64) {
	var tasks []models.Task
	var count int64
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&tasks).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Task{}).Count(&count).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&tasks).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Task{}).Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
			panic(err)
		}
	}
	for i := range tasks {
		tasks[i].TaskLike = Likes_get(tasks[i].ID, 2)
	}
	return tasks, count
}
func GetTaskSliceListWithSearchByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Task, int64) {
	var tasks []models.Task
	var count int64
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&tasks).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Task{}).Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Count(&count).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&tasks).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Task{}).
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
			panic(err)
		}
	}
	for i := range tasks {
		tasks[i].TaskLike = Likes_get(tasks[i].ID, 2)
	}
	return tasks, count
}
func GetPeopleSliceListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.People, int64) {
	var peoples []models.People
	var count int64
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("people_id", "people_introduce", "people_majorskill",
			"people_type", "people_max", "people_min", "people_ip", "people_details").Preload(clause.Associations).
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&peoples).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.People{}).Count(&count).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.DB.Select("people_id", "people_introduce", "people_majorskill",
			"people_type", "people_max", "people_min", "people_ip", "people_details").Preload(clause.Associations).Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&peoples).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.People{}).Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
			panic(err)
		}
	}
	return peoples, count
}
func GetPeopleSliceListWithSearchByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.People, int64) {
	var peoples []models.People
	var count int64
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("people_id", "people_introduce", "people_majorskill",
			"people_type", "people_max", "people_min", "people_ip", "people_details").Preload(clause.Associations).
			Where("people_introduce like ? or people_majorskill like ? or people_type like ? or people_details like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&peoples).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.People{}).Where("people_introduce like ? or people_majorskill like ? or people_type like ? or people_details like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Count(&count).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.DB.Select("people_id", "people_introduce", "people_majorskill",
			"people_type", "people_max", "people_min", "people_ip", "people_details").Preload(clause.Associations).
			Where("people_introduce like ? or people_majorskill like ? or people_type like ? or people_details like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&peoples).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.People{}).
			Where("people_introduce like ? or people_majorskill like ? or people_type like ? or people_details like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Count(&count).Error; err != nil {
			panic(err)
		}
	}
	return peoples, count
}

// 通过userid获取任务列表 !!!自己看不需要user信息
func GetTaskListForMe(user_id int64) []models.Task {
	var tasks []models.Task

	if err := db.DB.Select("id", "task_cover", "task_name", "task_subject", "task_min",
		"task_max", "task_details", "post_user_id", "task_type").Where(
		"post_user_id = ?", user_id).Preload("PicturesUrlList").Find(&tasks).Error; err != nil {
		panic(err)
	}
	for i := range tasks {
		tasks[i].TaskLike = Likes_get(tasks[i].ID, 3)
	}
	return tasks
}

func GetTaskById(task_id int64) models.Task {
	var task models.Task
	if err := db.DB.Preload(clause.Associations).Where("id = ?", task_id).Find(&task).Error; err != nil {
		panic(err)
	}
	return task
}

func GetTaskMidById(taskId int64) models.TaskMid {
	var task models.TaskMid
	if err := db.DB.Where("id = ?", taskId).Preload(clause.Associations).Find(&task).Error; err != nil {
		panic(err)
	}
	return task
}

func GetTaskInfoForMe(task_id int64) models.Task {
	var task models.Task
	if err := db.DB.Select("id", "task_name", "task_require", "task_message", "task_number", "task_min",
		"task_max", "task_progress", "post_user_id", "task_type", "task_cover").Where("id = ?", task_id).Preload(clause.Associations).Find(&task).Error; err != nil {
		panic(err)
	}
	return task
}

func GetTaskInfoForVisitor(task_id int64) (models.Task, []models.Task, error) {
	var task models.Task
	if err := db.DB.Select("id", "task_name", "task_view", "task_max", "task_min", "task_percent",
		"task_donation", "task_bedonated", "task_details", "task_like", "task_forward", "task_collection",
		"post_user_id", "task_type", "task_cover", "task_status", "task_introduce").Preload(clause.Associations).Where(
		"id = ?", task_id).Find(&task).Error; err != nil {
		return task, nil, err
	}
	tasks := GetUserAllTasks(task.PostUserId, task_id)
	return task, tasks, nil
}

func GetTaskMidInfo(task_id int64) (models.TaskMid, error) {
	var task models.TaskMid
	if err := db.DB.Select("id", "task_name", "task_view", "task_max", "task_min", "task_percent",
		"task_donation", "task_bedonated", "task_details", "task_like", "task_forward", "task_collection",
		"post_user_id", "task_type", "task_cover", "task_status", "task_introduce").Preload(clause.Associations).Where(
		"id = ?", task_id).Find(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func AddTaskCollection(taskId int64, userId int64) error {
	var user models.User
	task := GetTaskById(taskId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionTask").Append(&task)
		return err
	}
	return nil
}

func CancelTaskCollection(taskId int64, userId int64) error {
	var user models.User
	task := GetTaskById(taskId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionTask").Delete(&task)
		return err
	}
	return nil
}

func GetCollectionTaskList(userId int64) []models.Task {
	var tasks []models.Task
	var user models.User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("CollectionTask").Find(&tasks); err != nil {
			panic(err)
		}
	}
	for index, work := range tasks {
		var t models.Task
		if err := db.DB.Where("id = ?", work.ID).Preload("TaskSubject").Find(&t).Error; err != nil {
			panic(err)
		}
		tasks[index].TaskSubject = t.TaskSubject
		tasks[index].PostUser = user
	}
	return tasks
}

func DeleteTask(taskId int64) ([]models.TaskPicturesUrl, error) {
	var task models.Task
	var picture models.TaskPicturesUrl
	// TODO 删除关联 找出url供fs删除
	var pictures []models.TaskPicturesUrl
	// if err := db.DB.Where("work_id = ? ", workId).Find(&pictures).Error; err != nil {
	// 	return nil, err
	// }
	var t models.Task
	if err := db.DB.Where("id = ? ",
		taskId).Preload("PicturesUrlList").Preload("CollectUser").Preload("LikeUser").Preload("OrderUser").Find(&t).Error; err != nil {
		return nil, err
	}
	global.Log.WithField("task", t.LikeUser).Info("del task")
	pictures = t.PicturesUrlList
	flag := false
	for _, v := range global.TaskPhotourls {
		if v == t.TaskCover {
			flag = true
			break
		}
	}
	if !flag {
		pictures = append(pictures, models.TaskPicturesUrl{Url: t.TaskCover})
	}
	db.DB.Model(&task).Association("PicturesUrlList").Delete(pictures)
	db.DB.Model(&task).Association("TaskSubject").Delete(models.TaskSubjectItem{TaskId: taskId})
	err := db.DB.Select("LikeUser").Where("id = ?", taskId).Delete(&t).Error
	global.Log.WithError(err).Info("del task")
	//
	for _, user := range t.CollectUser {
		if err := CancelTaskCollection(taskId, user.UserId); err != nil {
			return nil, err
		}
	}
	for _, user := range t.LikeUser {
		if err := DeleteTaskLike(taskId, user.UserId); err != nil {
			return nil, err
		}
	}
	for _, user := range t.OrderUser {
		if err := DelTaskItem(user.UserId, taskId); err != nil {
			return nil, err
		}
	}
	// TODO 删除subject表
	db.DB.Where("task_id = ?", taskId).Unscoped().Delete(&models.TaskSubjectItem{})
	//

	if err := db.DB.Raw("delete from task_pictures_urls where task_id = ?", taskId).Find(&picture).Error; err == nil {
		if err = db.DB.Raw("delete from tasks where id = ?", taskId).Find(&task).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	return pictures, nil
}

func DeleteTaskMid(taskId int64) ([]models.TaskMidPicturesUrl, error) {
	var task models.TaskMid
	var picture models.TaskMidPicturesUrl
	// TODO 删除关联 找出url供fs删除
	var pictures []models.TaskMidPicturesUrl
	// if err := db.DB.Where("work_id = ? ", workId).Find(&pictures).Error; err != nil {
	// 	return nil, err
	// }
	var t models.TaskMid
	if err := db.DB.Where("id = ? ",
		taskId).Preload("PicturesUrlList").Find(&t).Error; err != nil {
		return nil, err
	}
	pictures = t.PicturesUrlList
	flag := false
	for _, v := range global.TaskPhotourls {
		if v == t.TaskCover {
			flag = true
			break
		}
	}
	if !flag {
		pictures = append(pictures, models.TaskMidPicturesUrl{Url: t.TaskCover})
	}
	db.DB.Model(&task).Association("PicturesUrlList").Delete(pictures)
	db.DB.Model(&task).Association("TaskSubject").Delete(models.TaskMidSubjectItem{TaskId: taskId})

	// TODO 删除subject表
	db.DB.Where("task_id = ?", taskId).Unscoped().Delete(&models.TaskMidSubjectItem{})
	//

	if err := db.DB.Raw("delete from task_mid_pictures_urls where task_id = ?", taskId).Find(&picture).Error; err == nil {
		if err = db.DB.Raw("delete from task_mids where id = ?", taskId).Find(&task).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	return pictures, nil
}

func GetAllPeopleList() []models.People {
	var peoples []models.People
	if err := db.DB.Where("people_status = ?", 0).Select("people_id", "people_introduce", "people_majorskill",
		"people_type", "people_max", "people_min", "people_ip", "people_details").Find(&peoples).Error; err != nil {
		panic(err)
	}
	for index, people := range peoples {
		if err := db.DB.Select("item").Where("people_id = ?", people.PeopleId).Find(&peoples[index].PeopleSubject).Error; err != nil {
			panic(err)
		}
		peoples[index].PeopleIp = strings.Split(people.PeopleIp, "/")[1]
	}
	return peoples
}

func GetAllPeopleMidList() []models.PeopleMid {
	var peoples []models.PeopleMid
	if err := db.DB.Where("people_status = ?", 0).Select("people_id", "people_introduce", "people_majorskill",
		"people_type", "people_max", "people_min", "people_ip", "people_details").Find(&peoples).Error; err != nil {
		panic(err)
	}
	for index, people := range peoples {
		if err := db.DB.Select("item").Where("people_id = ?", people.PeopleId).Find(&peoples[index].PeopleSubject).Error; err != nil {
			panic(err)
		}
		peoples[index].PeopleIp = strings.Split(people.PeopleIp, "/")[1]
	}
	return peoples
}

func TaskChangeStatus(Id int64, status int) {
	if err := db.DB.Table("tasks").Where("id = ?", Id).Update("task_status", status).Error; err != nil {
		panic(err)
	}
}

func DelTaskItem(uId int64, oId int64) error {
	user, err := GetUserById(uId)
	if err != nil {
		return err
	}
	updata := map[string]interface{}{"task_status": 0}
	if err := db.DB.Table("tasks").Where("id = ?", oId).Updates(updata).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&user).Association("OrderTask").Delete(&models.Task{ID: oId}); err != nil {
		return nil
	}
	return nil
}

func AddCandidate(taskId int64, userId int64) error {
	candidate := models.Candidate{
		TaskId: taskId,
		UserId: userId,
	}
	if err := db.DB.Model(&candidate).Create(&candidate).Error; err != nil {
		return err
	}
	return nil
}
func CancelCandidate(taskId int64, userId int64) error {
	candidate := models.Candidate{
		TaskId: taskId,
		UserId: userId,
	}
	if err := db.DB.Model(&candidate).Where("task_id = ? and user_id = ?", taskId, userId).Delete(&candidate).Error; err != nil {
		return err
	}
	return nil
}
func GetCandidateList(taskId int64) []int64 {
	var list []int64
	if err := db.DB.Model(&models.Candidate{}).Select("user_id").Where("task_id = ?", taskId).Find(&list).Error; err != nil {
		panic(err)
	}
	return list
}
