package work_type

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(api gin.IRouter) {
	api.GET("", getPaged)
	api.GET("select", getAll)
	api.POST("", create)
	api.DELETE(":id", _delete)
	api.PUT("", update)
}
