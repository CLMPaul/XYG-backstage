package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterBackstageMessage(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/postofficialmessage/", api.PostOfficialMessage)
}
