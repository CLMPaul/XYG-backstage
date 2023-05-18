package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xueyigou_demo/internal/errcode"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/service"
)

var (
	ErrInvalidToken   = errcode.NewHttpError(http.StatusUnauthorized, "invalid auth token")
	ErrLicenseExpired = errcode.ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  "license_not_active",
		ErrorMsg:   "license not active or expired",
	}
)

type LoginAccount = service.LoginAccount

type contextKeyLoginAccount struct{}

//goland:noinspection GoUnusedExportedFunction
func GetLoginAccount(c *gin.Context) *LoginAccount {
	if user, err := GetRequestContextValue[*LoginAccount](c, contextKeyLoginAccount{}); err == nil {
		return user
	}
	return nil
}

func SetLoginAccount(c *gin.Context, user *LoginAccount) {
	SetRequestContextValue(c, contextKeyLoginAccount{}, user)
}

func SetLoginAccountError(c *gin.Context, err error) {
	SetRequestContextError(c, contextKeyLoginAccount{}, err)
}

// RequireLoginAccount 验证用户信息，如果返回 err 为空，则 user 必不为空
func RequireLoginAccount(c *gin.Context) (user *LoginAccount, err error) {
	defer func() {
		if err != nil {
			SetLoginAccountError(c, err)
		} else {
			SetLoginAccount(c, user)
		}
	}()

	var token string
	token, err = RequireAuthToken(c)
	if err != nil {
		return
	}

	user, err = service.AccountService.Authenticate(token)
	if err != nil {
		return
	}

	if user == nil {
		err = ErrInvalidToken
	}

	return
}

// LoginRequired 中间件，要求接口调用 RequireLoginAccount 并返回 err 为空，否则会中止请求
func LoginRequired(c *gin.Context) {
	if _, err := RequireLoginAccount(c); err != nil {
		web.HandleError(c, err)
		return
	}
}
