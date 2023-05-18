package authentication

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/internal/web/middlewares"
	"xueyigou_demo/service"
)

type actionAuthentication struct {
	action string
}

func (a actionAuthentication) Authenticate(c *gin.Context) error {
	account, err := middlewares.RequireLoginAccount(c)
	if err != nil {
		return err
	}

	granted, err := service.RoleService.AuthenticatePermission(account.UID, a.action)
	if err != nil {
		return err
	}
	if granted {
		return nil
	}
	return ErrPermissionDenied
}

//goland:noinspection GoUnusedExportedFunction
func Action(action string) Authentication {
	return actionAuthentication{action: action}
}
