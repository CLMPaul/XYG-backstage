package router

import (
	"xueyigou_demo/api"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouterfile(apiRouter *gin.RouterGroup) {
	apiRouter.Use(middleware.JWT(0))
	apiRouter.POST("/upload/", api.UploadPicture)
	apiRouter.DELETE("/delete/", api.DeletePicture)
	//apiRouter.POST("uplo", api.UploadFile)
}
