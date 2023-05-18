package api

import (
	"strings"
	"time"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"

	"github.com/gin-gonic/gin"
)

func DisAbleAToken(c *gin.Context) {
	global.Expires = time.Now().Add(time.Hour * 24 * 31)
}

func EnAbleAToken(c *gin.Context) {
	global.Expires = time.Time{}
}

func DisAbleRToken(c *gin.Context) {
	global.RTokenExpires = time.Now().Add(time.Hour * 24 * 31)
}

func EnAbleRToken(c *gin.Context) {
	global.RTokenExpires = time.Time{}
}

func ATokenValid(c *gin.Context) {
	c.JSON(200, gin.H{
		"result_status": 0,
		"result_msg":    "token valid",
	})
}

func RTokenValid(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer", "", -1)
	res, err := middleware.Refresh(token)
	if err == nil {
		c.JSON(200, gin.H{
			"result_status": 0,
			"result_msg":    "Rtoken valid",
		})
	} else {
		c.JSON(200, res)
	}

}
