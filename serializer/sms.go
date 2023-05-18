package serializer

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
)

type SendSMSResponse struct {
	Response
	BizId     string
	RequestId string
}

func BuildSendSMSResponse(body *dysmsapi20170525.SendSmsResponseBody) *SendSMSResponse {
	var bizid string
	if body.BizId == nil {
		bizid = ""
	} else {
		bizid = *body.BizId
	}
	var request_id string
	if body.RequestId == nil {
		request_id = ""
	} else {
		request_id = *body.RequestId
	}
	var msg string
	if body.Message == nil {
		msg = ""
	} else {
		msg = *body.Message
	}
	return &SendSMSResponse{
		Response:  BuildSuccessResponse(msg),
		BizId:     bizid,
		RequestId: request_id,
	}
}
