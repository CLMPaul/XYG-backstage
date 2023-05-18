package service

import (
	"xueyigou_demo/cache"
	"xueyigou_demo/global"
	"xueyigou_demo/serializer"
	"xueyigou_demo/tools"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sirupsen/logrus"
)

const (
	accessKeyId     = "LTAI5t7mo9mLUmboXeXwkvF9"
	accessKeySecret = "shaIaZ6UW7g1AfCVGrOBaw4yRdFuki"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendSms(sendSmsRequest *dysmsapi20170525.SendSmsRequest) interface{} {
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return serializer.BuildFailResponse(_err.Error())
	}
	code := tools.Code()
	sendSmsRequest.TemplateParam = tea.String("{\"code\":\"" + code + "\"}")
	cache.CaheSmsCode(*sendSmsRequest.PhoneNumbers, code, *sendSmsRequest.TemplateCode)
	global.Log.WithFields(logrus.Fields{
		"phone": *sendSmsRequest.PhoneNumbers,
		"code":  code,
	}).Info("genrate code")
	runtime := &util.RuntimeOptions{}
	_result, tryErr := func() (_result *dysmsapi20170525.SendSmsResponse, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_result, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return nil, _err
		}

		return _result, nil
	}()
	if tryErr != nil {
		global.Log.WithError(tryErr).Error("send sms")
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return serializer.BuildFailResponse(_err.Error())
		}
	}

	logrus.WithField("result", _result).Info("sms result")
	return serializer.BuildSendSMSResponse(_result.Body)
}

func AddSmsSign(addSmsSignRequest *dysmsapi20170525.AddSmsSignRequest) interface{} {
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return serializer.BuildFailResponse(_err.Error())
	}

	runtime := &util.RuntimeOptions{}
	_result, tryErr := func() (_result *dysmsapi20170525.AddSmsSignResponse, _err error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_err = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_result, _err = client.AddSmsSignWithOptions(addSmsSignRequest, runtime)
		if _err != nil {
			global.Log.WithError(_err).Info("add sms err")
			return
		}

		return
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return serializer.BuildFailResponse(_err.Error())
		}
	}
	return _result
}
