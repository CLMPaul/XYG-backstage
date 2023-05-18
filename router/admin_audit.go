package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterAdminAudit(apiRouter *gin.RouterGroup) {

	apiRouter.GET("/getRegistration/", middleware.JWT(2), api.GetPendingAccount)
	apiRouter.POST("/postRegistration/", middleware.JWT(2), api.AuditAccount)
	apiRouter.Use(middleware.JWT(1))
	apiRouter.GET("/gettasklist/", api.GetTaskMidList)
	apiRouter.POST("/task/", api.PostTaskMid)
	apiRouter.GET("/taskinfo/", api.GetTaskMidInfo)
	apiRouter.GET("/getpeoplelist/", api.GetPeopleMidList)
	apiRouter.POST("/people/", api.PostPeopleMid)
}
