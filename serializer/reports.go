package serializer

import (
	"strconv"
	"time"
	"xueyigou_demo/models"
)

type Report struct {
	ReportDetails string `json:"report_details"`
	ReportReason  string `json:"report_reason"`
	ReportTime    string `json:"report_time"`
	ReporterID    string `json:"reporter_id"`
	ReportID      string `json:"report_id"`
}

type ReportTaskList struct {
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
	CreatedTime   time.Time `json:"created_time"` //创建时间
	Report        Report    `json:"report"`
}
type ReportTaskResponse struct {
	Response
	TaskList                      []ReportTaskList `json:"tasklist"`
	The01Gm3Cncrkty3M5E7Aq1E3Wrm8 interface{}      `json:"01GM3CNCRKTY3M5E7AQ1E3WRM8"`
}

type ReportPeopleList struct {
	PeopleDetails    string   `json:"people_details"`    // 接单详细信息
	PeopleHeadPhoto  string   `json:"people_head_photo"` // 头像图片url
	PeopleID         string   `json:"people_id"`
	PeopleIntroduce  string   `json:"people_introduce"` // 接单简介
	PeopleIP         string   `json:"people_ip"`
	PeopleMajorskill string   `json:"people_majorskill"` // 主要技能
	PeopleMax        int64    `json:"people_max"`        // 最高价格
	PeopleMin        int64    `json:"people_min"`        // 最低价格
	PeopleName       string   `json:"people_name"`       // 名称
	PeopleSubject    []string `json:"people_subject"`    // 具体擅长领域标签
	PeopleType       string   `json:"people_type"`       // 专业类型
	Report           Report   `json:"report"`
}

type ReportPeopleResponse struct {
	Response
	PeopleList []ReportPeopleList `json:"people_list"` // 人才列表
}

type ReportWorkList struct {
	Report        Report   `json:"report"`
	UserInfo      Userinfo `json:"user_info"`
	WorkID        int64    `json:"work_id"`
	WorkIntroduce string   `json:"work_introduce"`
	WorkMax       int64    `json:"work_max"`
	WorkMin       int64    `json:"work_min"`
	WorkName      string   `json:"work_name"`
	WorkPicture   string   `json:"work_picture"`
	WorkStatus    int64    `json:"work_status"`
	WorkSubject   []string `json:"work_subject"`
}

type Userinfo struct {
	UserHeadPhoto string `json:"user_head_photo"`
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
}

type ReportWorkResponse struct {
	Response
	WorkList                      []ReportWorkList `json:"work_list"` // 人才列表
	The01Gm3Dry9Xvgasndmkbq9Jaaw9 interface{}      `json:"01GM3DRY9XVGASNDMKBQ9JAAW9"`
}

func BuildReportPeopleListRespond(Reports []models.Report, People []models.People, User []models.User) ReportPeopleResponse {
	response := Response{
		ResultStatus: 200,
		ResultMsg:    "Get reports of people list successfully",
	}
	ReportPeopleRep := ReportPeopleResponse{
		Response: response,
	}
	for i, people := range People {
		report_people_list := ReportPeopleList{
			PeopleDetails:    people.PeopleDetails,
			PeopleHeadPhoto:  User[i].HeaderPhoto,
			PeopleID:         strconv.FormatInt(people.PeopleId, 10),
			PeopleIntroduce:  people.PeopleIntroduce,
			PeopleIP:         people.PeopleIp,
			PeopleMajorskill: people.PeopleMajorskill,
			PeopleMax:        people.PeopleMax,
			PeopleMin:        people.PeopleMin,
			PeopleName:       User[i].Name,
			PeopleType:       people.PeopleType,
		}
		for _, subject := range people.PeopleSubject {
			report_people_list.PeopleSubject = append(report_people_list.PeopleSubject, subject.Item)
		}

		report_people_list.Report = Report{
			ReportDetails: Reports[i].ReportDetails,
			ReportReason:  strconv.Itoa((Reports[i].ReportReason)),
			ReportTime:    Reports[i].CreatedAt.Format("2006-01-02 15:04:05"),
			ReporterID:    strconv.FormatInt(Reports[i].ReporterId, 10),
			ReportID:      strconv.FormatInt(Reports[i].ID, 10),
		}
		ReportPeopleRep.PeopleList = append(ReportPeopleRep.PeopleList, report_people_list)

	}
	return ReportPeopleRep
}

func BuildReportTaskListRespond(Reports []models.Report, Task []models.Task) ReportTaskResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get task list success",
	}
	ReportTaskRep := ReportTaskResponse{
		Response: response,
	}

	for i, task := range Task {
		report_task := ReportTaskList{
			TaskIntroduce: task.TaskIntroduce,
			TaskID:        strconv.FormatInt(task.ID, 10),
			TaskName:      task.TaskName,
			TaskPicture:   task.TaskCover,
			TaskMax:       task.TaskMax,
			TaskMin:       task.TaskMin,
			TaskType:      task.TaskType,
			TaskStatus:    task.TaskStatus,
			CreatedTime:   task.CreatedAt,
		}
		for _, subject := range task.TaskSubject {
			report_task.TaskSubject = append(report_task.TaskSubject, subject.Item)
		}
		if task.PostUser.UserId != 0 {
			report_task.PostUserId = task.PostUserId
			report_task.PostUserName = task.PostUser.Name
			report_task.PostHeadPhoto = task.PostUser.HeaderPhoto
		}
		report_task.Report = Report{
			ReportDetails: Reports[i].ReportDetails,
			ReportReason:  strconv.Itoa((Reports[i].ReportReason)),
			ReportTime:    Reports[i].CreatedAt.Format("2006-01-02 15:04:05"),
			ReporterID:    strconv.FormatInt(Reports[i].ReporterId, 10),
			ReportID:      strconv.FormatInt(Reports[i].ID, 10),
		}
		ReportTaskRep.TaskList = append(ReportTaskRep.TaskList, report_task)
	}
	return ReportTaskRep
}

func BuildReportWorkListRespond(Reports []models.Report, Work []models.Work) ReportWorkResponse {
	response := Response{
		ResultStatus: 200,
		ResultMsg:    "Get works Info success",
	}
	ReportWorkRep := ReportWorkResponse{
		Response: response,
	}

	//for i, work := range Work {
	//report_work := ReportWorkList{
	//	Introduce: work.Introduce,
	//	Title:      work.Title,
	//	WorkMax:       work.WorkMax,
	//	WorkMin:       work.WorkMin,
	//	WorkID:        work.ID,
	//	WorkStatus:    int64(work.WorkStatus),
	//	CoverPicture:   work.CoverPicture,
	//}
	//for _, subject := range work.WorkSubject {
	//	report_work.WorkSubject = append(report_work.WorkSubject, subject.Item)
	//}

	//report_work.UserInfo = Userinfo{
	//	UserID:        strconv.Itoa(int(work.PostUserId)),
	//	UserName:      work.PostUser.Name,
	//	UserHeadPhoto: work.PostUser.HeaderPhoto,
	//}
	//report_work.Report = Report{
	//	ReportDetails: Reports[i].ReportDetails,
	//	ReportReason:  strconv.Itoa((Reports[i].ReportReason)),
	//	ReportTime:    Reports[i].CreatedAt.Format("2006-01-02 15:04:05"),
	//	ReporterID:    strconv.FormatInt(Reports[i].ReporterId, 10),
	//	ReportID:      strconv.FormatInt(Reports[i].ID, 10),
	//}
	//
	//ReportWorkRep.WorkList = append(ReportWorkRep.WorkList, report_work)
	//}

	return ReportWorkRep
}
