package router

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api"
)

func InitRouterBackstageWelfare(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/modify/", api.PostWelfareModify)
	apiRouter.POST("/delete/", api.PostWelfareDelete)
}
