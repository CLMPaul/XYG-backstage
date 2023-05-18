package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
)

func InitRounterAppletSchoolactivity(apiRouter *gin.RouterGroup) {
	//apiRouter.GET("/find/", api.WorkFind)

	//apiRouter.GET("/eventworks/", api.EventWorks)
	apiRouter.GET("/activityList/", api.ActivitiyList)
	apiRouter.GET("/activityItem/", api.GetEvent)
	apiRouter.GET("/activityMembers/", api.GetMembers)
	apiRouter.POST("/activitySign/", api.PostActivitySign)

}
