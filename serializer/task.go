package serializer

import (
	"strconv"
	"time"
	"xueyigou_demo/models"
)

type TaskPublishResponse struct {
	Response
	TaskId string `json:"task_id"`
}

type TaskInfoResponseForMe struct {
	Response
	TaskInfo TaskInfoItemForMe `json:"task_msg"`
}

type TaskListResponseForBoth struct {
	Response
	TotalCount int64
	TaskInfo   []TaskListInfoForBoth `json:"task_list"`
}

type PeopleInfoResponse struct {
	Response
	TotalCount int64        `json:"total_count"`
	PeopleList []PeopleInfo `json:"people_list"`
}

type TaskInfoResponseForVisitor struct {
	Response
	PicturesURLList []string               `json:"pictures_url_list"` // 获取该任务的图片集地址
	TaskInfo        TaskInfoItemForVisitor `json:"task_info"`         // 任务详细信息
	UserInfo        userInfoTask           `json:"user_info"`         // 发布该任务的用户信息
	UserOtherTask   []UserOtherTask        `json:"user_other_task"`   // 用户其他任务列表
}

type userInfoTask struct {
	UserContribution int64  `json:"user_contribution"` // 用户公益贡献值
	UserID           string `json:"user_id"`           // 发布该商品的用户id
	UserIdentity     string `json:"user_identity"`     // 用户身份，用户身份，如：普通用户
	UserMedals       int64  `json:"user_medals"`       // 用户获得学易购奖章数
	UserName         string `json:"user_name"`         // 企业名/用户名
	UserHeadPhoto    string `json:"user_head_photo"`   // 用户头像
}

type UserOtherTask struct {
	OtherTaskPictureList []string `json:"other_task_picture_list"` // 用户其他任务图片地址
	OtherTaskID          string   `json:"other_task_id"`           // 用户其他任务id
	OtherTaskName        string   `json:"other_task_name"`         // 用户其他任务名称
	// OtherTaskPrice       int64    `json:"other_task_price"`         // 用户其他任务价格
	TaskMax     int64  `json:"task_max"` // 最高价格
	TaskMin     int64  `json:"task_min"` // 最低价格
	TaskPicture string `json:"task_picture"`
}
type TaskInfoItemForVisitor struct {
	TaskBedonated  string `json:"task_bedonated"`  // 任务被捐赠对象
	TaskCollection int64  `json:"task_collection"` // 任务被收藏数
	TaskDetails    string `json:"task_details"`    // 任务简介
	TaskIntroduce  string `json:"task_introduce"`
	TaskDonation   string `json:"task_donation"` // 任务捐赠对象
	TaskForward    int64  `json:"task_forward"`  // 任务转发数
	TaskID         string `json:"task_id"`       // 当前任务的id
	TaskLike       int64  `json:"task_like"`     // 任务点赞量
	TaskName       string `json:"task_name"`     // 任务名
	TaskPercent    int64  `json:"task_percent"`  // 任务捐赠百分比
	// TaskPrice      int64  `json:"task_price"`      // 任务价格
	TaskMax     int64     `json:"task_max"`     // 最高价格
	TaskMin     int64     `json:"task_min"`     // 最低价格
	TaskSubject []string  `json:"task_subject"` // 任务所属科目
	TaskView    uint64    `json:"task_view"`    // 任务被浏览量
	TaskType    string    `json:"task_type"`
	TaskPicture string    `json:"task_picture"`
	TaskStatus  int       `json:"task_status"`
	CreatedTime time.Time `json:"created_time"` //创建时间
}

type TaskInfoItemForMe struct {
	TaskID       string `json:"task_id"`       // 任务id
	TaskMax      int64  `json:"task_max"`      // 最高价格
	TaskMessage  string `json:"task_message"`  // 任务留言
	TaskMin      int64  `json:"task_min"`      // 最低价格
	TaskName     string `json:"task_name"`     // 任务名
	TaskNumber   string `json:"task_number"`   // 任务编号
	TaskProgress string `json:"task_progress"` // 任务进度
	TaskRequire  string `json:"task_require"`  // 任务要求
	TaskType     string `json:"task_type"`
	TaskPicture  string `json:"task_picture"`
	TaskLike     int64  `json:"task_like"`
}

type TaskListInfoForBoth struct {
	PostUserId    int64     `json:"poster_id,omitempty"`
	PostUserName  string    `json:"poster_name,omitempty"`
	PostHeadPhoto string    `json:"poster_head_photo"`
	TaskIntroduce string    `json:"task_introduce"` // 任务简介/信息
	TaskID        string    `json:"task_id"`        // 任务id
	TaskName      string    `json:"task_name"`      // 任务名称
	TaskPicture   string    `json:"task_picture"`   // 任务封面图片url
	TaskMax       int64     `json:"task_max"`       // 最高价格
	TaskMin       int64     `json:"task_min"`       // 最低价格
	TaskSubject   []string  `json:"task_subject"`   // 任务所属科目
	TaskType      string    `json:"task_type"`
	TaskStatus    int       `json:"task_status"`
	TaskLike      int64     `json:"task_like"`
	CreatedTime   time.Time `json:"created_time"` //创建时间
	TaskDetails   string    `json:"task_details"`
}

type PeopleInfo struct {
	PeopleIntroduce  string   `json:"people_introduce"` // 人才具体自我介绍
	PeopleDetails    string   `json:"people_details"`
	PeopleIp         string   `json:"people_ip"`
	PeopleHeadPhoto  string   `json:"people_head_photo"` // 人才头像图片地址url
	PeopleID         string   `json:"people_id"`         // 人才的个人中心id
	PeopleMajorskill string   `json:"people_majorskill"` // 人才主要技能
	PeopleName       string   `json:"people_name"`       // 人才名称
	PeopleSubject    []string `json:"people_subject"`    // 人才具体擅长领域标签
	PeopleType       string   `json:"people_type"`       // 人才专业类型
	PeopleMax        int64    `json:"people_max"`        // 人才最高报价
	PeopleMin        int64    `json:"people_min"`        // 人才最低报价
}

func BuildTaskInfoResponseForMe(task models.Task) TaskInfoResponseForMe {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task info success",
	}
	task_info := TaskInfoItemForMe{
		TaskID:       strconv.FormatInt(task.ID, 10),
		TaskMax:      task.TaskMax,
		TaskMessage:  task.TaskMessage,
		TaskMin:      task.TaskMin,
		TaskName:     task.TaskName,
		TaskNumber:   task.TaskNumber,
		TaskProgress: task.TaskProgress,
		TaskRequire:  task.TaskRequire,
		TaskType:     task.TaskType,
		TaskPicture:  task.TaskCover,
		TaskLike:     task.TaskLike,
	}
	return TaskInfoResponseForMe{
		Response: response,
		TaskInfo: task_info,
	}
}

func BuildTaskListInfoResponseForBoth(tasks []models.Task) TaskListResponseForBoth {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task list success",
	}
	var tasks_info []TaskListInfoForBoth
	for _, task := range tasks {
		task_info := TaskListInfoForBoth{
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			TaskLike:      task.TaskLike,
			CreatedTime:   task.CreatedAt,
			TaskDetails:   task.TaskDetails,
		}
		for _, subject := range task.TaskSubject {
			task_info.TaskSubject = append(task_info.TaskSubject, subject.Item)
		}
		if task.PostUser.UserId != 0 {
			task_info.PostUserId = task.PostUserId
			task_info.PostUserName = task.PostUser.Name
			task_info.PostHeadPhoto = task.PostUser.HeaderPhoto
		}
		tasks_info = append(tasks_info, task_info)
	}

	task_list_response := TaskListResponseForBoth{
		Response: response,
		TaskInfo: tasks_info,
	}
	return task_list_response
}

func BuildTaskMidListResponse(tasks []models.TaskMid) TaskListResponseForBoth {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task list success",
	}
	var tasksInfo []TaskListInfoForBoth
	for _, task := range tasks {
		taskInfo := TaskListInfoForBoth{
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			TaskLike:      task.TaskLike,
			CreatedTime:   task.CreatedAt,
			TaskDetails:   task.TaskDetails,
		}
		for _, subject := range task.TaskSubject {
			taskInfo.TaskSubject = append(taskInfo.TaskSubject, subject.Item)
		}
		if task.PostUser.UserId != 0 {
			taskInfo.PostUserId = task.PostUserId
			taskInfo.PostUserName = task.PostUser.Name
			taskInfo.PostHeadPhoto = task.PostUser.HeaderPhoto
		}
		tasksInfo = append(tasksInfo, taskInfo)
	}

	task_list_response := TaskListResponseForBoth{
		Response: response,
		TaskInfo: tasksInfo,
	}
	return task_list_response
}

func BuildTaskSliceListResponse(tasks []models.Task, count int64) TaskListResponseForBoth {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task slice list success",
	}
	var tasks_info []TaskListInfoForBoth
	for _, task := range tasks {
		task_info := TaskListInfoForBoth{
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			TaskLike:      task.TaskLike,
			CreatedTime:   task.CreatedAt,
			TaskDetails:   task.TaskDetails,
		}
		for _, subject := range task.TaskSubject {
			task_info.TaskSubject = append(task_info.TaskSubject, subject.Item)
		}
		if task.PostUser.UserId != 0 {
			task_info.PostUserId = task.PostUserId
			task_info.PostUserName = task.PostUser.Name
			task_info.PostHeadPhoto = task.PostUser.HeaderPhoto
		}
		tasks_info = append(tasks_info, task_info)
	}

	task_list_response := TaskListResponseForBoth{
		Response:   response,
		TotalCount: count,
		TaskInfo:   tasks_info,
	}
	return task_list_response
}
func BuildTaskSliceListWithSearchResponse(tasks []models.Task, count int64) TaskListResponseForBoth {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task slice list with search success",
	}
	var tasks_info []TaskListInfoForBoth
	for _, task := range tasks {
		task_info := TaskListInfoForBoth{
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			TaskLike:      task.TaskLike,
			CreatedTime:   task.CreatedAt,
			TaskDetails:   task.TaskDetails,
		}
		for _, subject := range task.TaskSubject {
			task_info.TaskSubject = append(task_info.TaskSubject, subject.Item)
		}
		if task.PostUser.UserId != 0 {
			task_info.PostUserId = task.PostUserId
			task_info.PostUserName = task.PostUser.Name
			task_info.PostHeadPhoto = task.PostUser.HeaderPhoto
		}
		tasks_info = append(tasks_info, task_info)
	}

	task_list_response := TaskListResponseForBoth{
		Response:   response,
		TotalCount: count,
		TaskInfo:   tasks_info,
	}
	return task_list_response
}
func BuildTaskInfoResponseForVisitor(task models.Task, tasks []models.Task, taskView uint64) TaskInfoResponseForVisitor {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task info success",
	}
	user := task.PostUser
	taskInfo := TaskInfoItemForVisitor{
		TaskBedonated:  task.TaskBedonated,
		TaskCollection: task.TaskCollection,
		TaskDetails:    task.TaskDetails,
		TaskIntroduce:  task.TaskIntroduce,
		TaskDonation:   task.TaskDonation,
		TaskForward:    task.TaskForward,
		TaskID:         strconv.FormatInt(task.ID, 10),
		TaskLike:       task.TaskLike,
		TaskName:       task.TaskName,
		TaskPercent:    task.TaskPercent,
		TaskMax:        task.TaskMax,
		TaskMin:        task.TaskMin,
		TaskView:       taskView,
		TaskType:       task.TaskType,
		TaskPicture:    task.TaskCover,
		TaskStatus:     task.TaskStatus,
		CreatedTime:    task.CreatedAt,
	}
	for _, subject := range task.TaskSubject {
		taskInfo.TaskSubject = append(taskInfo.TaskSubject, subject.Item)
	}
	var userTasks []UserOtherTask
	for _, task := range tasks {
		item := UserOtherTask{
			OtherTaskID:   GenrateObjectId(models.TASK, task.ID),
			OtherTaskName: task.TaskName,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskPicture:   task.TaskCover,
		}
		for _, url := range task.PicturesUrlList {
			item.OtherTaskPictureList = append(item.OtherTaskPictureList, url.Url)
		}
		userTasks = append(userTasks, item)
	}
	userInfo := userInfoTask{
		UserContribution: int64(user.Contribution),
		UserID:           strconv.Itoa(int(user.UserId)),
		UserIdentity:     user.UserIdentity,
		UserMedals:       int64(user.Medals),
		UserName:         user.Name,
		UserHeadPhoto:    user.HeaderPhoto,
	}

	ret := TaskInfoResponseForVisitor{
		Response:      response,
		TaskInfo:      taskInfo,
		UserInfo:      userInfo,
		UserOtherTask: userTasks,
	}
	for _, url := range task.PicturesUrlList {
		ret.PicturesURLList = append(ret.PicturesURLList, url.Url)
	}
	return ret
}

func BuildTaskMidInfo(task models.TaskMid) TaskInfoResponseForVisitor {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task info success",
	}
	user := task.PostUser
	taskInfo := TaskInfoItemForVisitor{
		TaskBedonated:  task.TaskBedonated,
		TaskCollection: task.TaskCollection,
		TaskDetails:    task.TaskDetails,
		TaskIntroduce:  task.TaskIntroduce,
		TaskDonation:   task.TaskDonation,
		TaskForward:    task.TaskForward,
		TaskID:         strconv.FormatInt(task.ID, 10),
		TaskLike:       task.TaskLike,
		TaskName:       task.TaskName,
		TaskPercent:    task.TaskPercent,
		TaskMax:        task.TaskMax,
		TaskMin:        task.TaskMin,
		TaskType:       task.TaskType,
		TaskPicture:    task.TaskCover,
		TaskStatus:     task.TaskStatus,
		CreatedTime:    task.CreatedAt,
	}
	for _, subject := range task.TaskSubject {
		taskInfo.TaskSubject = append(taskInfo.TaskSubject, subject.Item)
	}
	userInfo := userInfoTask{
		UserContribution: int64(user.Contribution),
		UserID:           strconv.Itoa(int(user.UserId)),
		UserIdentity:     user.UserIdentity,
		UserMedals:       int64(user.Medals),
		UserName:         user.Name,
		UserHeadPhoto:    user.HeaderPhoto,
	}

	ret := TaskInfoResponseForVisitor{
		Response: response,
		TaskInfo: taskInfo,
		UserInfo: userInfo,
	}
	for _, url := range task.PicturesUrlList {
		ret.PicturesURLList = append(ret.PicturesURLList, url.Url)
	}
	return ret
}

func BuildPublishTaskResponse(task_id int64) TaskPublishResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "publish task",
	}
	return TaskPublishResponse{
		Response: response,
		TaskId:   strconv.Itoa(int(task_id)),
	}
}

type CollectionTask struct {
	PostUserId    int64    `json:"poster_id,omitempty"`
	PostUserName  string   `json:"poster_name,omitempty"`
	PostHeadPhoto string   `json:"poster_head_photo"`
	TaskIntroduce string   `json:"task_introduce"` // 任务简介/信息
	TaskID        string   `json:"task_id"`        // 任务id
	TaskName      string   `json:"task_name"`      // 任务名称
	TaskPicture   string   `json:"task_picture"`   // 任务封面图片url
	TaskMax       int64    `json:"task_max"`       // 最高价格
	TaskMin       int64    `json:"task_min"`       // 最低价格
	TaskSubject   []string `json:"task_subject"`   // 任务所属科目
	TaskType      string   `json:"task_type"`
	TaskStatus    int      `json:"task_status"` // 任务状态
}
type CollectionTaskListResponse struct {
	Response
	CollectionTaskList []CollectionTask `json:"collection_task_list"`
}

func BuildCollectionTaskListResponse(tasks []models.Task) CollectionTaskListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get collection_tasks_list success",
	}
	var collectionTaskList []CollectionTask
	for _, task := range tasks {
		collectionTask := CollectionTask{
			PostUserId:    task.PostUserId,
			PostUserName:  task.PostUser.Name,
			PostHeadPhoto: task.PostUser.HeaderPhoto,
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskStatus:    task.TaskStatus,
			TaskType:      task.TaskType,
		}
		for _, subject := range task.TaskSubject {
			collectionTask.TaskSubject = append(collectionTask.TaskSubject, subject.Item)
		}
		collectionTaskList = append(collectionTaskList, collectionTask)
	}
	collectionTaskListResponse := CollectionTaskListResponse{
		Response:           response,
		CollectionTaskList: collectionTaskList,
	}
	return collectionTaskListResponse
}
func BuildPeopleSliceListResponse(peoples []models.People, users []models.User, count int64) PeopleInfoResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get people slice list success",
	}
	var peoples_info []PeopleInfo
	for index, people := range peoples {
		people_info := PeopleInfo{
			PeopleIntroduce:  people.PeopleIntroduce,
			PeopleDetails:    people.PeopleDetails,
			PeopleIp:         people.PeopleIp,
			PeopleHeadPhoto:  users[index].HeaderPhoto,
			PeopleID:         strconv.FormatInt(people.PeopleId, 10),
			PeopleMajorskill: people.PeopleMajorskill,
			PeopleName:       users[index].Name,
			PeopleType:       people.PeopleType,
			PeopleMax:        people.PeopleMax,
			PeopleMin:        people.PeopleMin,
		}
		for _, subject := range people.PeopleSubject {
			people_info.PeopleSubject = append(people_info.PeopleSubject, subject.Item)
		}
		peoples_info = append(peoples_info, people_info)
	}

	peopleInfoResponse := PeopleInfoResponse{
		Response:   response,
		TotalCount: count,
		PeopleList: peoples_info,
	}
	return peopleInfoResponse
}
func BuildPeopleSliceListWithSearchResponse(peoples []models.People, users []models.User, count int64) PeopleInfoResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get people slice list with search success",
	}
	var peoples_info []PeopleInfo
	for index, people := range peoples {
		people_info := PeopleInfo{
			PeopleIntroduce:  people.PeopleIntroduce,
			PeopleDetails:    people.PeopleDetails,
			PeopleIp:         people.PeopleIp,
			PeopleHeadPhoto:  users[index].HeaderPhoto,
			PeopleID:         strconv.FormatInt(people.PeopleId, 10),
			PeopleMajorskill: people.PeopleMajorskill,
			PeopleName:       users[index].Name,
			PeopleType:       people.PeopleType,
			PeopleMax:        people.PeopleMax,
			PeopleMin:        people.PeopleMin,
		}
		for _, subject := range people.PeopleSubject {
			people_info.PeopleSubject = append(people_info.PeopleSubject, subject.Item)
		}
		peoples_info = append(peoples_info, people_info)
	}

	peopleInfoResponse := PeopleInfoResponse{
		Response:   response,
		TotalCount: count,
		PeopleList: peoples_info,
	}
	return peopleInfoResponse
}
func BuildPeopleInfoResponse(peoples []models.People, users []models.User) PeopleInfoResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get people list success",
	}
	var peoplesInfo []PeopleInfo
	for index, people := range peoples {
		peopleInfo := PeopleInfo{
			PeopleIntroduce:  people.PeopleIntroduce,
			PeopleDetails:    people.PeopleDetails,
			PeopleIp:         people.PeopleIp,
			PeopleHeadPhoto:  users[index].HeaderPhoto,
			PeopleID:         strconv.FormatInt(people.PeopleId, 10),
			PeopleMajorskill: people.PeopleMajorskill,
			PeopleName:       users[index].Name,
			PeopleType:       people.PeopleType,
			PeopleMax:        people.PeopleMax,
			PeopleMin:        people.PeopleMin,
		}
		for _, subject := range people.PeopleSubject {
			peopleInfo.PeopleSubject = append(peopleInfo.PeopleSubject, subject.Item)
		}
		peoplesInfo = append(peoplesInfo, peopleInfo)
	}

	peopleInfoResponse := PeopleInfoResponse{
		Response:   response,
		PeopleList: peoplesInfo,
	}
	return peopleInfoResponse
}

func BuildPeopleMidListResponse(peoples []models.PeopleMid, users []models.User) PeopleInfoResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get people list success",
	}
	var peoplesInfo []PeopleInfo
	for index, people := range peoples {
		peopleInfo := PeopleInfo{
			PeopleIntroduce:  people.PeopleIntroduce,
			PeopleDetails:    people.PeopleDetails,
			PeopleIp:         people.PeopleIp,
			PeopleHeadPhoto:  users[index].HeaderPhoto,
			PeopleID:         strconv.FormatInt(people.PeopleId, 10),
			PeopleMajorskill: people.PeopleMajorskill,
			PeopleName:       users[index].Name,
			PeopleType:       people.PeopleType,
			PeopleMax:        people.PeopleMax,
			PeopleMin:        people.PeopleMin,
		}
		for _, subject := range people.PeopleSubject {
			peopleInfo.PeopleSubject = append(peopleInfo.PeopleSubject, subject.Item)
		}
		peoplesInfo = append(peoplesInfo, peopleInfo)
	}

	peopleInfoResponse := PeopleInfoResponse{
		Response:   response,
		PeopleList: peoplesInfo,
	}
	return peopleInfoResponse
}

type CandidateResponse struct {
	ResultStatus  int      `json:"result_status"`
	ResultMsg     string   `json:"result_msg"`
	CandidateList []string `json:"candidate_list"`
}

func BuildCandidateResponse(CandidateList []int64) CandidateResponse {
	var list []string
	for _, candidate := range CandidateList {
		id := strconv.FormatInt(candidate, 10)
		list = append(list, id)
	}
	response := CandidateResponse{
		ResultStatus:  0,
		ResultMsg:     "get candidate list success",
		CandidateList: list,
	}
	return response
}
