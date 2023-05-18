package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterPersonalmsg(apiRouter *gin.RouterGroup) {
	// 获取个人资料
	apiRouter.GET("/getmain/", api.GetUserInfo)
	// 获取地址信息
	apiRouter.GET("/getaddress/", middleware.JWT(0), api.GetAllAddress)
	//获取资质审核结果
	apiRouter.GET("/getindentity/", middleware.JWT(0), api.GetIndentity)
	// 提交个人资料
	apiRouter.POST("/postmain/", middleware.JWT(0), api.SetUserInfo)
	// 提交地址信息
	apiRouter.POST("/postaddress/", middleware.JWT(0), api.PostAddress)
	// 提交资质审核
	apiRouter.POST("/postindentity/", middleware.JWT(0), api.PostIndentity)
	// 获取个人卡片
	apiRouter.GET("/mycard/", middleware.JWT(0), api.GetPeopleForMe)
	// 账号安全
	apiRouter.POST("/accountsafety/", api.Resetkey)
}
