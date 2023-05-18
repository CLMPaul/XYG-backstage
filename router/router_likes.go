package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterLikes(apiRouter *gin.RouterGroup) {
	// 点赞操作
	apiRouter.Use(middleware.JWT(0))
	apiRouter.POST("/click/", api.Like)
	apiRouter.GET("/likeslist/", api.GetLikeList)
}
