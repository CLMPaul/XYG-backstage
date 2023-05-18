package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"
)

func InitRouterAppletLike(apiRouter *gin.RouterGroup) {
	// 点赞操作（取消点赞）
	apiRouter.POST("/", middleware.JWT(0), api.AppletLike)
}
