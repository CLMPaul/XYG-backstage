package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterWelfare(apiRouter *gin.RouterGroup) {
	// 公益活动列表
	apiRouter.GET("/mywelfare/", middleware.JWT(0), api.GetMyWelfareList)

	// 公益活动资讯列表
	apiRouter.GET("/welfarenewslist/", api.GetWelfareNewsList)
	// 公益活动详细信息
	apiRouter.GET("/welfareinfo/", api.GetWelfareActivityInfo)
	// 历史统计
	apiRouter.GET("/total/", api.GetWelfareHistory)
	// 公益活动信息
	apiRouter.GET("/eventinfo/", api.WelfareActivity)
	// 报名参加公益活动
	apiRouter.POST("/joinwelfare/", middleware.JWT(0), api.JoinWelfare)
	// 获取公益活动参加人员列表
	apiRouter.GET("/welfarepeoplelist/", api.GetWelfarePeople)
	// 添加个人的公益活动参加信息
	apiRouter.POST("addwelfareinfo/", api.AddWelfareInfo)
	// 发布公益活动
	apiRouter.POST("/postwelfare/", middleware.JWT(0), api.PostWelfareInfo)
}
