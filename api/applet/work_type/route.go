package work_type

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(api gin.IRouter) {
	api.GET("select", getAll)
}
