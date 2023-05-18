package authentication

import "github.com/gin-gonic/gin"

type orPermissionAuthentication []Authentication

func (a orPermissionAuthentication) Authenticate(c *gin.Context) (err error) {
	for _, auth := range []Authentication(a) {
		if err = auth.Authenticate(c); err == nil {
			return
		}
	}
	return
}

//goland:noinspection GoUnusedExportedFunction
func Or(auths ...Authentication) Authentication {
	return orPermissionAuthentication(auths)
}
