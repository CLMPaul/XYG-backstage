package authentication

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/internal/web/middlewares"
)

type administratorAuthentication struct{}

func (administratorAuthentication) Authenticate(c *gin.Context) error {
	user, err := middlewares.RequireLoginAccount(c)
	if err != nil {
		return err
	}
	if user.IsAdministrator {
		return nil
	}
	return ErrPermissionDenied
}

//goland:noinspection GoUnusedExportedFunction
func Administrator() Authentication {
	return administratorAuthentication{}
}
