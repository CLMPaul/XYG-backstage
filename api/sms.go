package api

import (
	"errors"
	"net/http"
	"xueyigou_demo/global"
	"xueyigou_demo/service"
	"xueyigou_demo/tools"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SendSmsForm struct {
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignName             *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SmsUpExtendCode      *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateParam        *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func SendSms(c *gin.Context) {
	form := &SendSmsForm{}
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	valid := tools.StringValid(global.PhonecPattern, *form.PhoneNumbers)
	if !valid {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("phone error")))
		return
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:     tea.String(*form.SignName),
		TemplateCode: tea.String(*form.TemplateCode),
		//TemplateParam: tea.String(*form.TemplateParam),
		// TemplateParam: tea.String("{\"code\":\"" + tools.Code() + "\"}"),
		PhoneNumbers: tea.String(*form.PhoneNumbers),
	}
	global.Log.WithFields(logrus.Fields{
		"phone number":  *sendSmsRequest.PhoneNumbers,
		"sign name":     *sendSmsRequest.SignName,
		"template code": *sendSmsRequest.TemplateCode,
		// "template param": *sendSmsRequest.TemplateParam,
	}).Info()
	response := service.SendSms(sendSmsRequest)
	c.JSON(http.StatusOK, response)
}

type AddSmsSignForm struct {
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SignFileList         []*AddSmsSignRequestSignFileList `json:"SignFileList,omitempty" xml:"SignFileList,omitempty" type:"Repeated"`
	SignName             *string                          `json:"SignName,omitempty" xml:"SignName,omitempty"`
	SignSource           *int32                           `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	SignType             *int32                           `json:"SignType,omitempty" xml:"SignType,omitempty"`
}

type AddSmsSignRequestSignFileList struct {
	FileContents *string `json:"FileContents,omitempty" xml:"FileContents,omitempty"`
	FileSuffix   *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
}

func AddSmsSign(c *gin.Context) {
	form := &AddSmsSignForm{}
	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	var sign_list []*dysmsapi20170525.AddSmsSignRequestSignFileList
	for _, item := range form.SignFileList {
		sign_list = append(sign_list,
			&dysmsapi20170525.AddSmsSignRequestSignFileList{
				FileContents: tea.String(*item.FileContents),
				FileSuffix:   tea.String(*item.FileSuffix),
			})
	}
	addSmsSignRequest := &dysmsapi20170525.AddSmsSignRequest{
		SignName:     tea.String(*form.SignName),
		SignSource:   tea.Int32(*form.SignSource),
		SignType:     tea.Int32(*form.SignType),
		SignFileList: sign_list,
	}
	response := service.AddSmsSign(addSmsSignRequest)
	c.JSON(http.StatusOK, response)
}
