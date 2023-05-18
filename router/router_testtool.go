package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterTest(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/enableatoken/", api.EnAbleAToken)
	apiRouter.POST("/disableatoken/", api.DisAbleAToken)
	apiRouter.POST("/enablertoken/", api.EnAbleRToken)
	apiRouter.POST("/disablertoken/", api.DisAbleRToken)
	apiRouter.POST("/atokenvalid/", middleware.JWT(0), api.ATokenValid)
	apiRouter.POST("/rtokenvalid/", api.RTokenValid)
}
