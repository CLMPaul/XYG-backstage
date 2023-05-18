package errcode

import (
	"fmt"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	StatusCode int
	ErrorCode  string
	ErrorMsg   string
	ErrorArgs  []interface{}
}

func (e ErrorResponse) Error() string {
	if e.ErrorMsg != "" {
		return e.ErrorMsg
	}
	return e.ErrorCode
}

//goland:noinspection GoUnusedExportedFunction
func NewHttpError(statusCode int, errorMessage string, args ...interface{}) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		ErrorCode:  strconv.Itoa(statusCode),
		ErrorMsg:   errorMessage,
		ErrorArgs:  args,
	}
}

func NewFailure(errorCode, errorMessage string, args ...interface{}) ErrorResponse {
	return ErrorResponse{
		StatusCode: http.StatusOK,
		ErrorCode:  errorCode,
		ErrorMsg:   errorMessage,
		ErrorArgs:  args,
	}
}

// 简化形式：errorMessage 为空
//
//goland:noinspection GoUnusedExportedFunction
func NewSimpleFailure(errorCode string) ErrorResponse {
	return NewFailure(errorCode, "")
}

type FailureTemplate struct {
	ErrorCode   string
	ErrorFormat string
}

func (t FailureTemplate) Format(args ...interface{}) ErrorResponse {
	return ErrorResponse{
		StatusCode: http.StatusOK,
		ErrorCode:  t.ErrorCode,
		ErrorMsg:   fmt.Sprintf(t.ErrorFormat, args...),
		ErrorArgs:  args,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewFailureTemplate(errorCode, format string) FailureTemplate {
	return FailureTemplate{
		ErrorCode:   errorCode,
		ErrorFormat: format,
	}
}
