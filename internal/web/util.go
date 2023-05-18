package web

import (
	"errors"
	"net/http"
	"xueyigou_demo/internal/errcode"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	ErrorCode    string        `json:"code,omitempty"`
	ErrorMessage string        `json:"msg,omitempty"`
	ErrArgs      []interface{} `json:"errArgs,omitempty"`
	Data         interface{}   `json:"data,omitempty"`
	Success      bool          `json:"success,omitempty"` // 兼容旧版本
}

func BadRequestResponse(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
		ErrorCode:    "400",
		ErrorMessage: msg,
	})
}

func InternalErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, APIResponse{
		ErrorCode:    "500",
		ErrorMessage: err.Error(),
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{Data: data, Success: true})
}

func FailureResponse(c *gin.Context, code, msg string, args ...interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    code,
		ErrorMessage: msg,
		ErrArgs:      args,
	})
}

// HandleError 统一使用 errcode 包下面的错误码实现
func HandleError(c *gin.Context, e error) {
	var errorResponse errcode.ErrorResponse
	if !errors.As(e, &errorResponse) {
		InternalErrorResponse(c, e)
	} else {
		c.AbortWithStatusJSON(errorResponse.StatusCode, APIResponse{
			ErrorCode:    errorResponse.ErrorCode,
			ErrorMessage: errorResponse.ErrorMsg,
			ErrArgs:      errorResponse.ErrorArgs,
		})
	}
}
