package serializer

import (
	"strconv"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
)

type PagedResponse[T any] struct {
	Data  []T   `json:"data"`
	Count int64 `json:"count"`
}

// Response 基础序列化器
type Response struct {
	ResultStatus int    `json:"result_status"`
	ResultMsg    string `json:"result_msg,omitempty"`
	Error        string `json:"error,omitempty"`
}

func BuildSuccessResponse(msg string) Response {
	return Response{
		ResultStatus: 0,
		ResultMsg:    msg,
	}
}
func BuildFailResponse(msg string) Response {
	return Response{
		ResultStatus: 1,
		ResultMsg:    msg,
	}
}

func BuildErrorResponse(code int) Response {
	return Response{
		ResultStatus: code,
		ResultMsg:    e.GetMsg(code),
	}
}

func BuildFailResponseWithCode(msg string, code int) Response {
	return Response{
		ResultStatus: code,
		ResultMsg:    msg,
	}
}

func BuildUserDoesNotExitResponse(msg string) Response {
	return Response{
		ResultStatus: 1,
		ResultMsg:    msg,
		Error:        "User Does Not Exit",
	}
}

func GenrateObjectId(o_type models.ObjectType, o_id int64) string {
	//TODO: map object to prefix_object
	return strconv.Itoa(int(o_id))
	// switch o_type {
	// case model.TASK:
	// 	return model.PTASK + strconv.Itoa(int(o_id))
	// case model.WELFARE:
	// 	return model.PWELFARE + strconv.Itoa(int(o_id))
	// case model.WORK:
	// 	return model.PWORK + strconv.Itoa(int(o_id))
	// default:
	// 	panic("ObjectType")
	// }
}
