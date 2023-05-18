package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"xueyigou_demo/internal/errcode"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthTokenNotSpecified = errcode.NewHttpError(http.StatusUnauthorized, "auth token is not specified")
	ErrBadToken              = errcode.NewHttpError(http.StatusUnauthorized, "bad auth token")
)

type contextKeyAuthToken struct{}

func GetAuthToken(c *gin.Context) string {
	if token, err := GetRequestContextValue[string](c, contextKeyAuthToken{}); err == nil {
		return token
	}
	return ""
}

func SetAuthToken(c *gin.Context, token string) {
	SetRequestContextValue[string](c, contextKeyAuthToken{}, token)
}

func RequireAuthToken(c *gin.Context) (token string, err error) {
	token, err = GetRequestContextValue[string](c, contextKeyAuthToken{})
	if err == nil || !errors.Is(err, ErrContextNotSet) {
		return
	}
	err = nil

	defer func() {
		if err != nil {
			SetRequestContextError(c, contextKeyAuthToken{}, err)
		} else {
			SetAuthToken(c, token)
		}
	}()

	token = c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("auth_token")
	}
	if token == "" {
		err = ErrAuthTokenNotSpecified
	} else {
		// token, _ = url.QueryUnescape(token)
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		} else {
			err = ErrBadToken
		}
	}

	return
}
