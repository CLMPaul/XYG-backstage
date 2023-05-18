package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
)

func PostReport(report models.Report) interface{} {
	err := dao.PostReport(report)
	if err != nil {
		return serializer.BuildFailResponse("post fail")
	} else {
		return serializer.BuildSuccessResponse("post success")
	}
}

func GetReportPeopleList(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportPeopleResponse {
	var people_list []models.People
	var report_list []models.Report
	var user_list []models.User
	report_list, people_list, user_list = dao.GetReportPeopleListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	return serializer.BuildReportPeopleListRespond(report_list, people_list, user_list)
}

func GetReportPeopleListWithSearch(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportPeopleResponse {
	var people_list []models.People
	var report_list []models.Report
	var user_list []models.User

	report_list, people_list, user_list = dao.GetReportPeopleWithSearchListByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	return serializer.BuildReportPeopleListRespond(report_list, people_list, user_list)
}

func GetReportTaskList(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportTaskResponse {

	var task_list []models.Task
	var report_list []models.Report
	report_list, task_list = dao.GetReportTaskListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)

	return serializer.BuildReportTaskListRespond(report_list, task_list)
}

func GetReportTaskListWithSearch(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportTaskResponse {
	var task_list []models.Task
	var report_list []models.Report

	report_list, task_list = dao.GetReportTaskWithSearchListByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	return serializer.BuildReportTaskListRespond(report_list, task_list)
}

func GetReportWorkList(currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportWorkResponse {

	var work_list []models.Work
	var report_list []models.Report
	report_list, work_list = dao.GetReportWorkListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)

	return serializer.BuildReportWorkListRespond(report_list, work_list)
}

func GetReportWorkListWithSearch(keywords string, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.ReportWorkResponse {
	var work_list []models.Work
	var report_list []models.Report

	report_list, work_list = dao.GetReportWorkWithSearchListByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	return serializer.BuildReportWorkListRespond(report_list, work_list)
}

func PostReportResult(item_type int64, item_id int64, result int64, report_id int64) interface{} {
	if result == 0 {
		if err := dao.DeleteAllReport(item_type, item_id); err != nil {
			return serializer.BuildFailResponse("Delete reports failed")
		}
		switch item_type {
		//case 0: //作品
		//	if urls, err := WorkService.Delete(item_id); err == nil {
		//		//TODO fs删除图片
		//		var urls_ []string
		//		for _, item := range urls {
		//			urls_ = append(urls_, item.Url)
		//		}
		//		if res, err1 := Deletefile(urls_); err1 != nil {
		//			return res
		//		}
		//		//
		//		return serializer.BuildSuccessResponse("Delete task success")
		//	} else {
		//		return serializer.BuildFailResponse("Delete task failed")
		//	}
		//	return serializer.BuildFailResponse("delete work")

		case 1: //任务
			if urls, err := dao.DeleteTask(item_id); err == nil {
				//TODO fs删除图片
				var urls_ []string
				for _, item := range urls {
					urls_ = append(urls_, item.Url)
				}
				if res, err1 := Deletefile(urls_); err1 != nil {
					return res
				}
				//
				return serializer.BuildSuccessResponse("Delete task success")
			} else {
				return serializer.BuildFailResponse("Delete task failed")
			}
			return serializer.BuildFailResponse("delete task")
		case 2: //人才
			if err := dao.DeletePeople(item_id); err == nil {
				return serializer.BuildSuccessResponse("Delete people success")
			} else {
				return serializer.BuildFailResponse("Delete people failed")
			}

		default:
			return serializer.BuildSuccessResponse("Object type is not defined")
		}

	}

	if err := dao.DeleteSingleReport(report_id); err != nil {
		return serializer.BuildFailResponse("Delete reports failed")
	}
	return serializer.BuildSuccessResponse("remain")
}
