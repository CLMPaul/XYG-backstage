package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterUser(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/follower/list/", api.UserFans)

	apiRouter.POST("/register/", api.UserRegister)
	apiRouter.POST("/login/", api.UserLogin)
	apiRouter.POST("/smslogin/", api.UserLoginViaCode)

	//apiRouter.GET("/follow/list/", api.UserAttention)
	apiRouter.POST("/follow/action/", api.UserAttentionAction)
	apiRouter.GET("follow/list/", api.UserAttention)
	apiRouter.GET("/", api.UserInfo)

	apiRouter.POST("/addordertask/", api.PostOrderTask)
	apiRouter.GET("/taskorderlist/", api.GetOrderTaskList)
	apiRouter.POST("/finishordertask/", middleware.JWT(0), api.FinishOrderTask)

	apiRouter.POST("/tokenRefresh/", api.RefreshToken)
	apiRouter.POST("/addorderwork/", api.PostOrderWork)
	apiRouter.GET("/workorderlist/", api.GetOrderWorkList)

	apiRouter.POST("/deleteorderitem/", api.DelOrderItem)
}
