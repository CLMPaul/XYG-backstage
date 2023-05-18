package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterAppletUser(apiRouter *gin.RouterGroup) {
	//登录
	apiRouter.POST("/login", api.WeChatAppLetCallBack)

	apiRouter.Use(middleware.JWT(0))
	// 个人的基本信息
	apiRouter.GET("information", api.UserInfo)
	// 我发布的作品列表
	apiRouter.GET("publish", api.GetUserWorks)
	// 我的活动列表
	apiRouter.GET("activities", api.GetUserActivities)
	// 关注操作
	apiRouter.POST("followaction", api.UserFollowAction)
	// 我关注的用户列表
	apiRouter.GET("followlist", api.UserAttention)
	// 用户反馈
	apiRouter.POST("feedback", api.Feedback)
}
