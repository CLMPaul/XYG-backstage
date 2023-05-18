package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("panic: %v\n\n%s\n", err, debug.Stack())
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
