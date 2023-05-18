package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterCollection(apiRouter *gin.RouterGroup) {
	// 商品列表
	apiRouter.GET("/work/", api.GetCollectionWorkList)
	// 任务列表
	apiRouter.GET("/task/", api.GetCollectionTaskList)
	// 公益列表
	apiRouter.GET("/welfare/", api.GetWelfareList)
	//活动列表
	apiRouter.GET("/events/", api.GetCollectionEventsList)
	// 收藏操作（取消收藏）
	apiRouter.POST("/collection/", middleware.JWT(0), api.Collection)
}
