package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterSMS(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/verification/", api.SendSms)
}
