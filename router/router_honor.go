package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterHonor(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/pricelist/", api.GetHonorList)

}
