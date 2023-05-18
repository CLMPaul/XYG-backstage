package authentication

import (
	"net/http"
	"xueyigou_demo/internal/errcode"

	"github.com/gin-gonic/gin"
)

type Authentication interface {
	Authenticate(c *gin.Context) error
}

var ErrPermissionDenied = errcode.NewHttpError(http.StatusForbidden, "permission denied")
