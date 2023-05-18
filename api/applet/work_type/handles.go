package work_type

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/service"
)

func getAll(c *gin.Context) {
	data, err := service.SubjectItemService.GetAll()
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}
