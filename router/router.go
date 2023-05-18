package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r gin.IRouter) {
	// public directory is used to serve static resources
	r.Static("/image", "./public/")

	// apiRouter := r.Group("/xueyisi") // 路由组加入中间件

	InitRouterUser(r.Group("/xueyigou/user"))
	InitRouterComment(r.Group("/xueyigou/comment"))
	InitRouterHonor(r.Group("/xueyigou/honor"))

	InitRouterTask(r.Group("/xueyigou/task"))
	InitRouterCollection(r.Group("/xueyigou/collection"))
	InitRouterPersonalmsg(r.Group("/xueyigou/personalmsg"))
	InitRouterWelfare(r.Group("/xueyigou/welfare"))
	InitRouterLikes(r.Group("/xueyigou/likes"))

	InitRouterMessage(r.Group("/xueyigou/message"))

	//apiRouter.POST("/personalCenter/", controller.UserBasicInfo)
	InitRouterfile(r.Group("/xueyigou/file"))

	InitRouterSMS(r.Group("/xueyigou/sms/"))

	InitRouterTest(r.Group("/xueyigou/test"))

	InitRouterSSO(r.Group("/xueyigou/sso"))

	InitRouterDefalutpic(r.Group("/xueyigou/defaultpic"))

	InitRouterReport(r.Group("/xueyigou/report"))
	InitRouterEvent(r.Group("/xueyigou/events"))
	//后台端口
	InitRouterBackstageMessage(r.Group("/backstage/message"))
	InitRouterBackstageEvent(r.Group("/backstage/events"))
	InitRouterAdminUser(r.Group("/backstage/user"))

	InitRouterAdminAudit(r.Group("/backstage/audit"))
	InitRouterBackstageWelfare(r.Group("/backstage/welfare"))
	InitRouterBackstageReport(r.Group("/backstage/report"))

	//小程序端口
	InitRounterAppletSchoolactivity(r.Group("/applet/schoolactivity"))
	InitRouterAppletUser(r.Group("/applet/user"))
	InitRouterAppletComment(r.Group("/applet/comment"))
	InitRouterAppletCollect(r.Group("/applet/collect"))
	InitRouterAppletLike(r.Group("/applet/like"))
}
