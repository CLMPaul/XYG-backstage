package api

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/api/backstage/work"
	"xueyigou_demo/api/backstage/work_type"
	"xueyigou_demo/static"

	appletWork "xueyigou_demo/api/applet/work"
	appletWorkType "xueyigou_demo/api/applet/work_type"
)

func SetupRouter(router gin.IRouter) {
	static.SetupStatic(router)

	//api := router.Group("api/v1")
	backstageAPI := router.Group("backstage")
	work_type.SetupRouter(backstageAPI.Group("work_type"))
	work.SetupRouter(backstageAPI.Group("work"))

	appletAPI := router.Group("applet")
	appletWork.SetupRouter(appletAPI.Group("work"))
	appletWorkType.SetupRouter(appletAPI.Group("work_type"))
}
