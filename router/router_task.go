package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterTask(apiRouter *gin.RouterGroup) {
	// 任务信息
	apiRouter.GET("/info/", api.GetTaskInfoForVisitor)

	apiRouter.GET("/people/", api.GetPeopleInfo)

	apiRouter.POST("/changeorderstatus/", api.ChangeStatus)
	// 获取分页的任务列表
	apiRouter.GET("/slicelist/", api.GetTaskSliceList)
	//获取分页的人才列表
	apiRouter.GET("/slicepeople/", api.GetPeopleSliceList)
	//操作任务候选人接单
	apiRouter.POST("/ordercandidate/", middleware.JWT(0), api.PostCandidate)
	apiRouter.GET("/candidatelist/", api.GetCandidate)
}
