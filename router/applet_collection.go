package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"
)

func InitRouterAppletCollect(apiRouter *gin.RouterGroup) {
	// 收藏操作（取消收藏）
	apiRouter.POST("/", middleware.JWT(0), api.AppletCollection)
}
