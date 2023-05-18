package service

import (
	"errors"
	"xueyigou_demo/cache"
)

func VerifSmsCodeWithPhone(phone_num string, code string, template_code string) error {
	// 测试后门
	if code == "Xueyigou" {
		return nil
	}
	verif := cache.VerifSmsCode(phone_num, code, template_code)

	//验证码过期
	if verif == 0 {
		return errors.New("验证码过期")
	} else if verif == 1 {
		return errors.New("验证码错误")
	}
	return nil
}
