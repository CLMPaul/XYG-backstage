package api

import (
	"net/http"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

func GetHonorList(c *gin.Context) {
	response := service.GetHonorList()
	c.JSON(http.StatusOK, response)
}
