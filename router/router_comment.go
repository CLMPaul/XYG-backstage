package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterComment(apiRouter *gin.RouterGroup) {
	apiRouter.Use(middleware.JWT(0))
	apiRouter.GET("/list/", api.CommentList)
	apiRouter.POST("/action/", api.CommentAction)
}
