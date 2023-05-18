package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
)

func GetTaskInfoForMe(task_id int64) serializer.TaskInfoResponseForMe {
	task := dao.GetTaskInfoForMe(task_id)
	return serializer.BuildTaskInfoResponseForMe(task)
}

func GetTaskInfoForVisitor(task_id int64) interface{} {
	task, tasks, err := dao.GetTaskInfoForVisitor(task_id)
	if err != nil {
		return serializer.BuildFailResponse("task not exit")
	}
	taskView := dao.TaskView(task.ID)
	dao.AddTaskView(task.ID)
	return serializer.BuildTaskInfoResponseForVisitor(task, tasks, taskView)
}

func GetTaskMidInfo(taskId int64) interface{} {
	task, err := dao.GetTaskMidInfo(taskId)
	if err != nil {
		return serializer.BuildFailResponse("task not exit")
	}
	return serializer.BuildTaskMidInfo(task)
}

func GetTaskListForMe(user_id int64) interface{} {
	_, err := dao.GetUserById(user_id)
	if err == nil {
		tasks := dao.GetTaskList(user_id)
		return serializer.BuildTaskListInfoResponseForBoth(tasks)
	} else {
		return nil
	}

}

func GetTaskMidList() serializer.TaskListResponseForBoth {
	tasks := dao.GetTaskMidList()
	return serializer.BuildTaskMidListResponse(tasks)
}

func GetTaskSliceList(status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.TaskListResponseForBoth {
	var task_list []models.Task
	var count int64
	if status == 0 {
		task_list, count = dao.GetTaskSliceListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	//var taskViewList []uint64
	//for _, task := range task_list {
	//	taskViewList = append(taskViewList, dao.TaskView(task.ID))
	//}
	return serializer.BuildTaskSliceListResponse(task_list, count)
}
func GetTaskSliceListWithSearch(keywords string, status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.TaskListResponseForBoth {
	var task_list []models.Task
	var count int64
	if status == 0 {
		task_list, count = dao.GetTaskSliceListWithSearchByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	//var taskViewList []uint64
	//for _, task := range task_list {
	//	taskViewList = append(taskViewList, dao.TaskView(task.ID))
	//}
	return serializer.BuildTaskSliceListWithSearchResponse(task_list, count)
}
func PublishTask(task models.TaskMid, Claim *middleware.UserClaims) *serializer.TaskPublishResponse {
	user, err := dao.GetUserById(Claim.Id)
	if err == nil {
		task.PostUser = user
		task.PostUserId = user.UserId
		task.ID = global.Worker.GetId()
		dao.PublishTask(&task)
	} else {
		return nil
	}
	ret := serializer.BuildPublishTaskResponse(task.ID)
	return &ret

}

func PostTaskMid(taskId int64, ispassed int64) interface{} {
	if ispassed == 1 {
		taskMid := dao.GetTaskMidById(taskId)
		task := models.Task{
			TaskIntroduce: taskMid.TaskIntroduce,
			TaskMax:       taskMid.TaskMax,
			TaskMin:       taskMid.TaskMin,
			TaskDetails:   taskMid.TaskDetails,
			TaskName:      taskMid.TaskName,
			TaskType:      taskMid.TaskType,
			TaskStatus:    taskMid.TaskStatus,
			TaskCover:     taskMid.TaskCover,
			PostUser:      taskMid.PostUser,
			PostUserId:    taskMid.PostUserId,
			ID:            taskId,
		}

		for _, subject := range taskMid.TaskSubject {
			task.TaskSubject = append(task.TaskSubject, models.TaskSubjectItem{Item: subject.Item})
		}
		for _, url := range taskMid.PicturesUrlList {
			task.PicturesUrlList = append(task.PicturesUrlList, models.TaskPicturesUrl{Url: url.Url})
		}

		dao.PublishTaskMid(&task)
	}
	//图片就不用删了
	if _, err := dao.DeleteTaskMid(taskId); err == nil {
		//TODO fs删除图片
		// var urls_ []string
		// for _, item := range urls {
		// 	urls_ = append(urls_, item.Url)
		// }
		// if res, err1 := Deletefile(urls_); err1 != nil {
		// 	return res
		// }
		//
		return serializer.BuildSuccessResponse("Delete task success")
	} else {
		return serializer.BuildFailResponse("Delete task failed")
	}
	return serializer.BuildFailResponse("delete task")
	ret := serializer.BuildPublishTaskResponse(taskId)
	return &ret

}

func GetPeopleSliceListForVisitor(status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.PeopleInfoResponse {
	var people_list []models.People
	var count int64
	if status == 0 {
		people_list, count = dao.GetPeopleSliceListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	var users []models.User
	for _, people := range people_list {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	return serializer.BuildPeopleSliceListResponse(people_list, users, count)
}
func GetPeopleSliceListForMe(userId int64, status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.PeopleInfoResponse {
	var people_list []models.People
	var count int64
	if status == 0 {
		people_list, count = dao.GetPeopleSliceListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	var users []models.User
	for _, people := range people_list {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	for index, people := range people_list {
		if people.PeopleId == userId {
			user := users[index]
			users = append(users[:index], users[index+1:]...)
			users = append([]models.User{user}, users...)
			people_list = append(people_list[:index], people_list[index+1:]...)
			people_list = append([]models.People{people}, people_list...)
		}
	}
	return serializer.BuildPeopleSliceListResponse(people_list, users, count)
}
func GetPeopleSliceListWithSearchForVisitor(keywords string, status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.PeopleInfoResponse {
	var people_list []models.People
	var count int64
	if status == 0 {
		people_list, count = dao.GetPeopleSliceListWithSearchByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	var users []models.User
	for _, people := range people_list {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	return serializer.BuildPeopleSliceListWithSearchResponse(people_list, users, count)
}
func GetPeopleSliceListWithSearchForMe(userId int64, keywords string, status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.PeopleInfoResponse {
	var people_list []models.People
	var count int64
	if status == 0 {
		people_list, count = dao.GetPeopleSliceListWithSearchByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
	}
	var users []models.User
	for _, people := range people_list {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	for index, people := range people_list {
		if people.PeopleId == userId {
			user := users[index]
			users = append(users[:index], users[index+1:]...)
			users = append([]models.User{user}, users...)
			people_list = append(people_list[:index], people_list[index+1:]...)
			people_list = append([]models.People{people}, people_list...)
		}
	}
	return serializer.BuildPeopleSliceListWithSearchResponse(people_list, users, count)
}

func GetPeopleInfoForVisitor() serializer.PeopleInfoResponse {
	peoples := dao.GetAllPeopleList()
	var users []models.User
	for _, people := range peoples {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	return serializer.BuildPeopleInfoResponse(peoples, users)
}

func GetPeopleMidList() serializer.PeopleInfoResponse {
	peoples := dao.GetAllPeopleMidList()
	var users []models.User
	for _, people := range peoples {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	return serializer.BuildPeopleMidListResponse(peoples, users)
}

func GetPeopleInfoForMe(userId int64) serializer.PeopleInfoResponse {
	peoples := dao.GetAllPeopleList()
	var users []models.User
	for _, people := range peoples {
		user, _ := dao.GetUserById(people.PeopleId)
		users = append(users, user)
	}
	for index, people := range peoples {
		if people.PeopleId == userId {
			user := users[index]
			users = append(users[:index], users[index+1:]...)
			users = append([]models.User{user}, users...)
			peoples = append(peoples[:index], peoples[index+1:]...)
			peoples = append([]models.People{people}, peoples...)
		}
		for _, user := range users {
			println(user.Name)
		}

	}
	return serializer.BuildPeopleInfoResponse(peoples, users)
}

func ChangeStatus(Id int64, status int, item int) serializer.Response {
	if item == 1 {
		dao.TaskChangeStatus(Id, status)
	} else if item == 2 {
		dao.MessageChangeStatus(Id, status)
	} else if item == 3 {
		dao.PeopleChangeStatus(Id, status)
	}
	ret := serializer.BuildSuccessResponse("Change status successfully")
	return ret
}

func AddCandidate(taskId int64, userId int64) serializer.Response {
	if err := dao.AddCandidate(taskId, userId); err != nil {
		return serializer.BuildFailResponse("add candidate failed")
	}
	return serializer.BuildSuccessResponse("add candidate success")
}
func CancelCandidate(taskId int64, userId int64) serializer.Response {
	if err := dao.CancelCandidate(taskId, userId); err != nil {
		return serializer.BuildFailResponse("cancel candidate failed")
	}
	return serializer.BuildSuccessResponse("cancel candidate success")
}
func GetCandidateList(taskId int64) serializer.CandidateResponse {
	candidateList := dao.GetCandidateList(taskId)
	return serializer.BuildCandidateResponse(candidateList)
}
