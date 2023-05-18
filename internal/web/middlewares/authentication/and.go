package authentication

import "github.com/gin-gonic/gin"

type andPermissionAuthentication []Authentication

func (a andPermissionAuthentication) Authenticate(c *gin.Context) (err error) {
	for _, auth := range []Authentication(a) {
		if err = auth.Authenticate(c); err != nil {
			return
		}
	}
	return
}

func And(auths ...Authentication) Authentication {
	return andPermissionAuthentication(auths)
}
