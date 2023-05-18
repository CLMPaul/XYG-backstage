package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"xueyigou_demo/config"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				ResultStatus: 400,
				ResultMsg:    fmt.Sprintf("%s%s", field, tag),
				Error:        fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			ResultStatus: 400,
			ResultMsg:    "JSON类型不匹配",
			Error:        fmt.Sprint(err),
		}
	}

	return serializer.Response{
		ResultStatus: 400,
		ResultMsg:    "参数错误",
		Error:        fmt.Sprint(err),
	}
}

func HandleObjectId(o_id string) (models.ObjectType, int64, error) {
	var object_type models.ObjectType
	var object_id int64
	if len(o_id) <= 3 {
		return object_type, object_id, e.ParmError{Parm: "object id"}
	}
	o_type := o_id[:3]
	switch o_type {
	case models.PTASK:
		object_type = models.TASK
	case models.PWORK:
		object_type = models.WORK
	case models.PWELFARE:
		object_type = models.WELFARE
	case models.PEVENT:
		object_type = models.EVENT
	default:
		return object_type, object_id, e.ParmError{Parm: "object id"}
	}
	object_id, err := strconv.ParseInt(o_id[3:], 10, 64)
	if err != nil {
		return object_type, object_id, e.ParmError{Parm: "object id"}
	}
	return object_type, object_id, nil
}

type refreshForm struct {
	Token serializer.Token
}

func RefreshToken(c *gin.Context) {
	var token refreshForm
	if err := c.ShouldBind(&token); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
	global.Log.WithField("r", token.Token.RefreshToken).Info("1")
	response, _ := middleware.Refresh(token.Token.RefreshToken)
	c.JSON(http.StatusOK, response)
}
