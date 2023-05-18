package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
)

func InitRouterBackstageReport(apiRouter *gin.RouterGroup) {
	//apiRouter.POST("/result/", api.PostReportResult)
	apiRouter.GET("/peoplelist/", api.GetReportPeopleList)
	apiRouter.GET("/tasklist/", api.GetReportTaskList)
	apiRouter.GET("/worklist/", api.GetReportWorkList)
	apiRouter.POST("/result/", api.PostReportResult)
}
