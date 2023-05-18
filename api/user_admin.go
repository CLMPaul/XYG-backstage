package api

import (
	"errors"
	"net/http"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/service"
	"xueyigou_demo/tools"

	"github.com/gin-gonic/gin"
)

// for login and register
type adminUserForm struct {
	Phone    string `json:"phone"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AdminRegister(c *gin.Context) {
	var form adminUserForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(errors.New("phone format error")))
		return
	}
	if len(form.Password) < 6 {
		c.JSON(200, ErrorResponse(errors.New("password too short")))
		return
	}
	res := service.AdminRegister(form.Phone, form.Code, form.Name, form.Password)
	c.JSON(http.StatusOK, res)
}

func AdminLogin(c *gin.Context) {
	var form adminUserForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(errors.New("phone format error")))
		return
	}
	res := service.AdminLogin(form.Phone, form.Code)
	c.JSON(http.StatusOK, res)
}

func AdminLoginWithPassword(c *gin.Context) {
	var form adminUserForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(errors.New("phone format error")))
		return
	}
	if len(form.Password) < 6 {
		c.JSON(200, ErrorResponse(errors.New("password too short")))
		return
	}
	res := service.AdminLoginWithPassword(form.Phone, form.Code, form.Password)
	c.JSON(http.StatusOK, res)
}

type adminPasswordResetForm struct {
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
	Phone       string `json:"phone"`
}

func AdminPasswordReset(c *gin.Context) {
	var form adminPasswordResetForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(errors.New("phone format error")))
		return
	}
	if len(form.NewPassword) < 6 {
		c.JSON(200, ErrorResponse(errors.New("password too short")))
		return
	}
	res := service.AdminPasswordReset(form.Phone, form.Code, form.NewPassword)
	c.JSON(http.StatusOK, res)
}

func GetPendingAccount(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}

	claim := userClaim.(*middleware.UserClaims)
	if claim.Privilege != 2 {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("非特权用户")))
		return
	}

	res := service.GetPendingAccount()
	c.JSON(http.StatusOK, res)
}

type auditForm struct {
	Phone  string `json:"phone"`
	IsPass int    `json:"isPass"`
}

func AuditAccount(c *gin.Context) {
	var form auditForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	vaild := tools.StringValid(global.PhonecPattern, form.Phone)
	if !vaild {
		c.JSON(200, ErrorResponse(errors.New("phone format error")))
		return
	}

	if form.IsPass != 0 && form.IsPass != 1 {
		c.JSON(200, ErrorResponse(errors.New("unknown ispass")))
		return
	}
	res := service.AuditAccount(form.Phone, form.IsPass)
	c.JSON(http.StatusOK, res)
}
