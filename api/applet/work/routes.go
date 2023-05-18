package work

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/middleware"
)

func SetupRouter(api gin.IRouter) {
	api.Use(middleware.JWT(0))
	api.GET("find", getPaged)
	api.GET("append", getPagedByUser)
	api.GET("details/:id", getInfo)
	api.POST("post", add)
	api.DELETE(":id", _delete)
	api.PUT("", update)
	api.PUT("take_down/:id", TakeDown)
}
