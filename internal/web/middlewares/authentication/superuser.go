package authentication

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/internal/web/middlewares"
)

type superuserAuthentication struct{}

func (superuserAuthentication) Authenticate(c *gin.Context) error {
	user, err := middlewares.RequireLoginAccount(c)
	if err != nil {
		return err
	}
	if user.IsSuperuser {
		return nil
	}
	return ErrPermissionDenied
}

//goland:noinspection GoUnusedExportedFunction
func Superuser() Authentication {
	return superuserAuthentication{}
}
