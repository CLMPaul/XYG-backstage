package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterReport(apiRouter *gin.RouterGroup) {
	apiRouter.Use(middleware.JWT(0))
	apiRouter.POST("/", api.PostReport)
}
