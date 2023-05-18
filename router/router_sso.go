package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterSSO(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/github/", api.GitHubLogin)
	apiRouter.POST("/githubcallbackhandle/", api.GitHubLoginCallback)

	apiRouter.GET("/wechat/", api.WeChatLogin)
	apiRouter.POST("wechatcallback", api.WeChatLoginCallBack)

	apiRouter.GET("/qq/", api.QQLogin)
	apiRouter.POST("qqcallback", api.QQLoginCallBack)

	apiRouter.POST("/bind", api.BindPhoneNum)

}
