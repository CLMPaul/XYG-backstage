package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

type taskform struct {
	TaskIntroduce string   `json:"task_introduce"` // 封面照片
	TaskDetails   string   `json:"task_details"`   // 任务详情介绍
	TaskMax       int64    `json:"task_max"`       // 最高价格
	TaskMin       int64    `json:"task_min"`       // 最低价格
	TaskName      string   `json:"task_name"`      // 任务名称
	TaskPicture   []string `json:"task_picture"`   // 任务相关照片
	TaskSubject   []string `json:"task_subject"`   // 学科类型
	TaskType      string   `json:"task_type"`      // 任务类别
	TaskStatus    int      `json:"task_status"`
}

// /xueyigou/mypublishment/task
func GetTaskInfoForMe(c *gin.Context) {
	o_id := c.Query("task_id")
	_, task_id, err := HandleObjectId(o_id)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	response := service.GetTaskInfoForMe(task_id)
	c.JSON(http.StatusOK, response)
}

// /xueyigou/mypublishment/taskdetails/
func GetTaskListForMe(c *gin.Context) {
	u_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	response := service.GetTaskListForMe(u_id)
	c.JSON(http.StatusOK, response)
}

// xueyigou/task/list/

func GetTaskSliceList(c *gin.Context) {
	status, _ := strconv.Atoi(c.Query("status"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchChoiceSecond := c.Query("searchChoiceSecond")
	keywords := c.Query("keywords")
	searchChoiceFirst := c.Query("searchChoiceFirst")
	if keywords == "" { //无搜索返回任务
		response := service.GetTaskSliceList(status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	} else { //返回搜索任务
		response := service.GetTaskSliceListWithSearch(keywords, status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	}
}
func GetTaskInfoForVisitor(c *gin.Context) {
	oId := c.Query("task_id")
	_, taskId, err := HandleObjectId(oId)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	response := service.GetTaskInfoForVisitor(taskId)
	c.JSON(http.StatusOK, response)
}

func PublishTask(c *gin.Context) {
	var task_form taskform
	if err := c.ShouldBind(&task_form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	if task_form.TaskMax < task_form.TaskMin {
		c.String(http.StatusOK, "Illegal data")
		return
	}
	task := models.TaskMid{
		TaskIntroduce: global.IllegalWords.Replace(task_form.TaskIntroduce, '*'),
		TaskMax:       task_form.TaskMax,
		TaskMin:       task_form.TaskMin,
		TaskDetails:   global.IllegalWords.Replace(task_form.TaskDetails, '*'),
		TaskName:      task_form.TaskName,
		TaskType:      task_form.TaskType,
		TaskStatus:    task_form.TaskStatus,
	}

	for _, subject := range task_form.TaskSubject {
		task.TaskSubject = append(task.TaskSubject, models.TaskMidSubjectItem{Item: string(subject)})
	}
	// if len(task_form.TaskPicture) == 0 {
	// 	for _, v := range global.TaskPhotourls {
	// 		task.TaskCover = v
	// 	}
	// }
	for index, url := range task_form.TaskPicture {
		if index == 0 {
			if url == "" {
				for _, v := range global.TaskPhotourls {
					task.TaskCover = v
					break
				}
			} else {
				task.TaskCover = url
			}
			continue
		}
		task.PicturesUrlList = append(task.PicturesUrlList, models.TaskMidPicturesUrl{Url: url})
	}
	// if task.TaskCover == "" {
	// 	task.TaskCover = global.TaskPhotourl
	// }
	response := service.PublishTask(task, claim)
	if response == nil {
		serializer.BuildUserDoesNotExitResponse("pubilsh task")
	}
	c.JSON(http.StatusOK, response)
}

func PostTaskMid(c *gin.Context) {
	taskId, _ := strconv.ParseInt(c.Query("task_id"), 10, 64)
	isPassed, _ := strconv.ParseInt(c.Query("ispassed"), 10, 64)

	response := service.PostTaskMid(taskId, isPassed)
	c.JSON(http.StatusOK, response)
}

func GetPeopleInfo(c *gin.Context) {
	userId := c.Query("user_id")

	if userId == "" {
		response := service.GetPeopleInfoForVisitor()
		c.JSON(http.StatusOK, response)
	} else {
		Id, _ := strconv.ParseInt(userId, 10, 64)
		response := service.GetPeopleInfoForMe(Id)
		c.JSON(http.StatusOK, response)
	}
}

func GetPeopleMidList(c *gin.Context) {
	response := service.GetPeopleMidList()
	c.JSON(http.StatusOK, response)
}

func GetPeopleSliceList(c *gin.Context) {
	userId := c.Query("user_id")
	status, _ := strconv.Atoi(c.Query("status"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchChoiceSecond := c.Query("searchChoiceSecond")
	keywords := c.Query("keywords")
	searchChoiceFirst := c.Query("searchChoiceFirst")
	if userId == "" {
		if keywords == "" { //无搜索返回任务
			response := service.GetPeopleSliceListForVisitor(status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
			c.JSON(http.StatusOK, response)
		} else { //返回搜索任务
			response := service.GetPeopleSliceListWithSearchForVisitor(keywords, status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
			c.JSON(http.StatusOK, response)
		}
	} else {
		id, _ := strconv.ParseInt(userId, 10, 64)
		if keywords == "" { //无搜索返回任务
			response := service.GetPeopleSliceListForMe(id, status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
			c.JSON(http.StatusOK, response)
		} else { //返回搜索任务
			response := service.GetPeopleSliceListWithSearchForMe(id, keywords, status, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
			c.JSON(http.StatusOK, response)
		}
	}
}

type StatusForm struct {
	ObjectId string `json:"object_id"`
	Status   int    `json:"status"`
	Item     int    `json:"item"`
}

func ChangeStatus(c *gin.Context) {
	var form StatusForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	if form.Item != 0 && form.Item != 1 && form.Item != 2 && form.Item != 3 {
		c.JSON(http.StatusOK, "Item is illegal")
		return
	}
	Id, _ := strconv.ParseInt(form.ObjectId, 10, 64)
	fmt.Println(form)
	response := service.ChangeStatus(Id, form.Status, form.Item)
	c.JSON(http.StatusOK, response)
}

func GetTaskMidList(c *gin.Context) {
	response := service.GetTaskMidList()
	c.JSON(http.StatusOK, response)
}

func GetTaskMidInfo(c *gin.Context) {
	taskId, _ := strconv.ParseInt(c.Query("task_id"), 10, 64)
	response := service.GetTaskMidInfo(taskId)
	c.JSON(http.StatusOK, response)
}

type CandidateForm struct {
	ActionType int    `json:"action_type"`
	TaskId     string `json:"task_id"`
	UserId     string `json:"user_id"`
}

func PostCandidate(c *gin.Context) {
	var form CandidateForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	taskId, _ := strconv.ParseInt(form.TaskId, 10, 64)
	userId, _ := strconv.ParseInt(form.UserId, 10, 64)
	if form.ActionType == 0 {
		response := service.AddCandidate(taskId, userId)
		c.JSON(http.StatusOK, response)
	} else if form.ActionType == 1 {
		response := service.CancelCandidate(taskId, userId)
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusOK, "action_type should be 0 or 1")
	}
}
func GetCandidate(c *gin.Context) {
	Id := c.Query("task_id")
	taskId, _ := strconv.ParseInt(Id, 10, 64)
	response := service.GetCandidateList(taskId)
	c.JSON(http.StatusOK, response)
}
