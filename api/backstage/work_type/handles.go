package work_type

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
	"xueyigou_demo/service"
)

func getPaged(c *gin.Context) {
	var req proto.WorkTypeRequest
	err := c.BindQuery(&req)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	data, err := service.SubjectItemService.GetPaged(req)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func getAll(c *gin.Context) {
	data, err := service.SubjectItemService.GetAll()
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func create(c *gin.Context) {
	var s models.WorkType
	err := c.BindJSON(&s)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	err = service.SubjectItemService.Create(&s)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, nil)
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
	err = service.SubjectItemService.Delete(id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, nil)
}

func update(c *gin.Context) {
	var s models.WorkType
	err := c.BindJSON(&s)
	if err != nil {
		web.BadRequestResponse(c, err.Error())
		return
	}
	err = service.SubjectItemService.Update(&s)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, nil)
}
