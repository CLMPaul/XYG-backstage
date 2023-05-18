package cache

import (
	"fmt"
	"time"
	"xueyigou_demo/global"

	"github.com/sirupsen/logrus"
)

func SmsTemplateCodeKey(template_code string, phone_num string) string {
	return fmt.Sprintf("sms:template_code:%s:phone_num%s", template_code, phone_num)
}

func CaheSmsCode(phone_num string, code string, template_code string) {
	SetByTTL(SmsTemplateCodeKey(template_code, phone_num), code, time.Minute*30)
	global.Log.WithField("key", SmsTemplateCodeKey(template_code, phone_num)).Info("set sms key")
}

// return 0 for code expiration 1 for code not match 2 for success
func VerifSmsCode(phone_num string, code string, template_code string) int {
	res, err := Get(SmsTemplateCodeKey(template_code, phone_num))
	if err != nil {
		global.Log.WithFields(logrus.Fields{
			"phone":  phone_num,
			"errmsg": err.Error(),
		}).Info("err occured while verifying msm code")
		return 0
	}
	if res != code {
		global.Log.WithFields(logrus.Fields{
			"phone":       phone_num,
			"param code":  code,
			"cahced code": res,
		}).Info("code not match")
		return 1
	}
	return 2
}
