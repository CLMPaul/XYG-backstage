package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterDefalutpic(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/get/", api.GetDefaultPic)
	apiRouter.POST("/upload/", api.UploadDefaultPic)
	apiRouter.POST("/update/", api.UpdateDefaultPic)
	apiRouter.POST("/delete/", api.DeleteDefaultPic)
}
