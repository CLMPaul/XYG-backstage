package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
)

func InitRouterBackstageEvent(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/submit/", api.PostEventSubmit)
	apiRouter.POST("/edit/", api.PostEventEdit)
	apiRouter.POST("/delete/", api.PostEventDelete)
}
func InitRouterEvent(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/eventsmain/", api.GetEvent)
	apiRouter.GET("/eventsleftmain/", api.GetEventSelf)
	apiRouter.GET("/eventslist/", api.GetEventList)
	apiRouter.GET("/eventsdata/", api.GetEvent)
}
