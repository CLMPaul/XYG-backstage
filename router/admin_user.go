package router

import (
	"xueyigou_demo/api"

	"github.com/gin-gonic/gin"
)

func InitRouterAdminUser(apiRouter *gin.RouterGroup) {
	apiRouter.POST("/register/", api.AdminRegister)
	apiRouter.POST("/login/", api.AdminLogin)
	apiRouter.POST("/passwordlogin/", api.AdminLoginWithPassword)
	apiRouter.POST("/resetpassword/", api.AdminPasswordReset)
}
