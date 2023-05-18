package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterMessage(apiRouter *gin.RouterGroup) {
	apiRouter.Use(middleware.JWT(0))
	apiRouter.GET("/systemmessage/", api.GetSystemMessage)
	apiRouter.POST("/sendsystemmessage/", api.PostSystemMessage)
	apiRouter.POST("/deletesystemmessage/", api.DeleteMessage)
	apiRouter.GET("/eachmessage/", api.GetInteractiveMessage)
	apiRouter.POST("/sendeachmessage/", api.PostInteractiveMessage)
	apiRouter.GET("/messagelength/", api.GetMessageLength)
	apiRouter.POST("/postofficialmessage/", api.PostOfficialMessage)
	apiRouter.GET("/officialmessage/", api.GetOfficialMessage)
	apiRouter.POST("/changemessagereadstatus/", api.PostMessageRead)
}
