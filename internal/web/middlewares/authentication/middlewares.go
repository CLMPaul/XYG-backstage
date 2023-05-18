package authentication

import (
	"github.com/gin-gonic/gin"
	"xueyigou_demo/internal/web"
)

func Using(a Authentication) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.IsAborted() {
			return
		}
		if err := a.Authenticate(c); err != nil {
			web.HandleError(c, err)
		}
	}
}

//goland:noinspection GoUnusedExportedFunction
func UsingAdministrator() gin.HandlerFunc {
	return Using(Administrator())
}

//goland:noinspection GoUnusedExportedFunction
func UsingAction(action string) gin.HandlerFunc {
	return Using(Action(action))
}
