package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteForm struct {
	ObjectId   string `json:"object_id"`
	ObjectType int    `json:"object_type"`
}

func DeleteObject(c *gin.Context) {
	var form DeleteForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be DeleteForm")
		return
	}
	//objectId, _ := strconv.ParseInt(form.ObjectId, 10, 64)
	//response := service.DeleteObject(objectId, form.ObjectType)
	//c.JSON(http.StatusOK, response)
}
