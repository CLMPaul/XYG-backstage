package work

import (
	"strconv"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/proto"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

type workAddForm struct {
	WorkIntroduce     string   `json:"work_introduce"`
	WorkDetails       string   `json:"work_details"`       // 内容详情介绍
	WorkMax           int64    `json:"work_max"`           // 最高价格
	WorkMin           int64    `json:"work_min"`           // 最低价格
	WorkName          string   `json:"work_name"`          // 作品名称
	WorkPicture       []string `json:"work_picture"`       // 照片集
	WorkSubject       []string `json:"work_subject"`       // 学科类型
	WorkType          string   `json:"work_type"`          // 作品类别
	WorkDonation      string   `json:"work_donation"`      // 作品捐赠对象
	WorkDonatepercent int64    `json:"work_donatepercent"` // 作品捐赠比例
}

func getPaged(c *gin.Context) {
	var req proto.WorkRequest
	err := c.BindQuery(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}

	data, err := service.WorkService.GetPaged(req)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func getInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	data, err := service.WorkService.FindById(id, 0)
	if err != nil {
		web.HandleError(c, err)
		return
	}

	web.SuccessResponse(c, data)
}

func auditApproved(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err = service.WorkService.UpdateState(id, 1)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}

func auditApprovedBatch(c *gin.Context) {
	var req struct {
		Ids []string `json:"ids"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if len(req.Ids) == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	for _, v := range req.Ids {
		id, _ := strconv.Atoi(v)
		err = service.WorkService.UpdateState(id, 1)
		if err != nil {
			web.HandleError(c, err)
			return
		}
	}

	web.SuccessResponse(c, req.Ids)
}

func auditNotApproved(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err = service.WorkService.UpdateState(id, 2)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}

func auditNotApprovedBatch(c *gin.Context) {
	var req struct {
		Ids []string `json:"ids"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if len(req.Ids) == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	for _, v := range req.Ids {
		id, _ := strconv.Atoi(v)
		err = service.WorkService.UpdateState(id, 2)
		if err != nil {
			web.HandleError(c, err)
			return
		}
	}

	web.SuccessResponse(c, req.Ids)
}

func auditTakeDown(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err = service.WorkService.UpdateState(id, 9)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}

func auditTakeDownBatch(c *gin.Context) {
	var req struct {
		Ids []string `json:"ids"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if len(req.Ids) == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	for _, v := range req.Ids {
		id, _ := strconv.Atoi(v)
		err = service.WorkService.UpdateState(id, 9)
		if err != nil {
			web.HandleError(c, err)
			return
		}
	}

	web.SuccessResponse(c, req.Ids)
}

func _delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	if id == 0 {
		web.BadRequestResponse(c, "id is empty")
		return
	}
	err = service.WorkService.Delete(id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, id)
}
