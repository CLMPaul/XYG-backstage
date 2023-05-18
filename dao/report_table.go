package dao

import (
	"gorm.io/gorm/clause"
	"xueyigou_demo/db"
	"xueyigou_demo/models"
)

func PostReport(report models.Report) error {
	if err := db.DB.Create(&report).Error; err != nil {
		return err
	}
	return nil
}

func GetReportPeopleListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.People, []models.User) {
	var report_list []models.Report
	var people_list []models.People
	var user_list []models.User
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", " reporter_id", "reports.id").Where("subject_item = ?", 2).Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("people_majorskill", "people_type", "people_introduce", "people_min", "people_max", "people_id", "people_ip", "people_details").Preload("PeopleSubject").Joins("INNER JOIN reports ON subject_id = people_id").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&people_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("name", "header_photo", "user_id").Joins("INNER JOIN reports ON subject_id = user_id").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&user_list).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", " reporter_id", "reports.id").Where("subject_item = ?", 2).Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("people_majorskill", "people_type", "people_introduce", "people_min", "people_max", "people_id", "people_ip", "people_details").Preload("PeopleSubject").Joins("INNER JOIN reports ON subject_id = people_id").Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&people_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("name", "header_photo", "user_id").Joins("INNER JOIN reports ON subject_id = user_id").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&user_list).Error; err != nil {
			panic(err)
		}
	}
	return report_list, people_list, user_list
}

func GetReportPeopleWithSearchListByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.People, []models.User) {
	var report_list []models.Report
	var people_list []models.People
	var user_list []models.User
	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 2).Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("people_majorskill", "people_type", "people_introduce", "people_min", "people_max", "people_id", "people_ip", "people_details").
			Preload("PeopleSubject").Joins("INNER JOIN reports ON subject_id = people_id").
			Where("people_details like ? or people_introduce like ? or people_id like ? or people_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&people_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("name", "header_photo", "user_id").
			Joins("INNER JOIN reports ON subject_id = user_id").
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&user_list).Error; err != nil {
			panic(err)
		}

	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", " reporter_id", "reports.id").
			Where("subject_item = ?", 2).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("people_majorskill", "people_type", "people_introduce", "people_min", "people_max", "people_id", "people_ip", "people_details").
			Preload("PeopleSubject").Joins("INNER JOIN reports ON subject_id = people_id").
			Where("people_details like ? or people_introduce like ? or people_id like ? or people_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("people_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&people_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("name", "header_photo", "user_id").
			Joins("INNER JOIN reports ON subject_id = user_id").
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&user_list).Error; err != nil {
			panic(err)
		}
	}
	return report_list, people_list, user_list
}

func GetReportTaskListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.Task) {
	var task_list []models.Task
	var report_list []models.Report

	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 2).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("tasks.id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Joins("INNER JOIN reports ON subject_id = tasks.id").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&task_list).Error; err != nil {
			panic(err)
		}

	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", " reporter_id", "reports.id").
			Where("subject_item = ?", 2).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Joins("INNER JOIN reports ON subject_id = id").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&task_list).Error; err != nil {
			panic(err)
		}
	}
	for i := range task_list {
		task_list[i].TaskLike = Likes_get(task_list[i].ID, 2)
	}
	return report_list, task_list
}

func GetReportTaskWithSearchListByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.Task) {
	var task_list []models.Task
	var report_list []models.Report

	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 2).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("tasks.id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Joins("INNER JOIN reports ON subject_id = tasks.id").
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&task_list).Error; err != nil {
			panic(err)
		}

	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", " reporter_id", "reports.id").
			Where("subject_item = ?", 2).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("id", "task_cover", "task_name", "task_min",
			"task_max", "task_introduce", "post_user_id", "task_type", "task_status").Preload(clause.Associations).
			Joins("INNER JOIN reports ON subject_id = id").
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&task_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Task{}).
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Error; err != nil {
			panic(err)
		}
	}
	for i := range task_list {
		task_list[i].TaskLike = Likes_get(task_list[i].ID, 2)
	}
	return report_list, task_list
}

func GetReportWorkListByTime(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.Work) {
	var work_list []models.Work
	var report_list []models.Report

	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 0).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("work_introduce", "works.id", "work_name", "work_picture", "work_max",
			"work_min", "post_user_id", "work_status", "work_type").Preload("WorkSubject").Preload("PostUser").
			Preload("PicturesUrlList").
			Joins("INNER JOIN reports ON subject_id = works.id").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&work_list).Error; err != nil {
			panic(err)
		}

	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 0).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Select("work_collect", "work_introduce", "works.id",
			"work_like", "work_name", "work_picture", "work_max",
			"work_min", "work_view", "post_user_id", "work_status", "work_type").Preload("WorkSubject").Preload("PostUser").
			Preload("PicturesUrlList").Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Joins("INNER JOIN reports ON subject_id = works.id").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&work_list).Error; err != nil {
			panic(err)
		}
	}

	return report_list, work_list
}

func GetReportWorkWithSearchListByTime(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) ([]models.Report, []models.Work) {
	var work_list []models.Work
	var report_list []models.Report

	if searchChoiceFirst == "全部" {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 0).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("work_introduce", "works.id", "work_name", "work_picture", "work_max",
			"work_min", "post_user_id", "work_status", "work_type").
			Joins("INNER JOIN reports ON subject_id = works.id").
			Where("works.id like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Find(&work_list).Error; err != nil {
			panic(err)
		}

	} else {
		if err := db.DB.Select("created_at", " report_reason", " report_details", " subject_id", "reporter_id", "reports.id").
			Where("subject_item = ?", 0).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).Order("created_at desc").
			Find(&report_list).Error; err != nil {
			panic(err)
		}

		if err := db.DB.Select("work_introduce", "works.id", "work_name", "work_picture", "work_max",
			"work_min", "post_user_id", "work_status", "work_type").
			Joins("INNER JOIN reports ON subject_id = works.id").
			Where("works.id like ? or work_introduce like ? or work_name like ? or work_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("work_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").
			Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&work_list).Error; err != nil {
			panic(err)
		}
		if err := db.DB.Model(&models.Work{}).
			Where("task_message like ? or task_introduce like ? or task_name like ? or task_type like ?", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%", "%"+keywords+"%").
			Where("task_type like ?", "%"+searchChoiceFirst+"/"+searchChoiceSecond+"%").Error; err != nil {
			panic(err)
		}
	}

	return report_list, work_list
}

func DeleteAllReport(item_type int64, item_id int64) error {
	var report models.Report
	if err := db.DB.Raw("delete from reports where subject_item = ? and subject_id=?", item_type, item_id).Find(&report).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSingleReport(report_id int64) error {
	var report models.Report
	if err := db.DB.Raw("delete from reports where  id =?", report_id).Find(&report).Error; err != nil {
		return err
	}
	return nil
}
