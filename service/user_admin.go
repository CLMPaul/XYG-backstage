package service

import (
	"errors"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"

	"gorm.io/gorm"
)

func AdminRegister(phone string, code string, name string, password string) interface{} {
	_, _, err := dao.GetAdminIdByPhone(phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildErrorResponse(e.ErrorExistPhone)
	}
	if _, err := dao.GetAdminByName(name); !errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildErrorResponse(e.ErrorExistNick)
	}

	err = VerifSmsCodeWithPhone(phone, code, global.TemplateCodeForRegister)
	if err != nil {
		return serializer.BuildFailResponseWithCode(err.Error(), e.ErrorCode)
	}

	admin := models.SuperUser{
		UserId:    global.Worker.GetId(),
		Telephone: phone,
		Name:      name,
		Password:  middleware.Md5Crypt(password, global.SlatSuperUserPassword),
	}
	for _, v := range global.HeadPhotourls {
		admin.HeaderPhoto = v
		break
	}
	err = dao.AdminRegister(&admin)
	if err != nil {
		global.Log.WithError(err).Error("AdminRegister")
		return serializer.BuildErrorResponse(e.ErrorDatabase)
	}
	return serializer.BuildAdminRegisterResponse()
}

func AdminLogin(phone string, code string) interface{} {
	super_user, err := dao.GetAdminByPhone(phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildFailResponseWithCode("not registered", 1)
	}
	if super_user.IsPass == 0 {
		return serializer.BuildFailResponseWithCode("等待审核", 2)
	}
	err = VerifSmsCodeWithPhone(phone, code, global.TemplateCodeForLogin)
	if err != nil {
		return serializer.BuildFailResponseWithCode(err.Error(), e.ErrorCode)
	}
	privilege := 1
	if super_user.IsSuper == 1 {
		privilege = 2
	}
	newClaim := middleware.UserClaims{Id: super_user.UserId, Phone: phone, Privilege: privilege}
	token, _ := middleware.GenerateToken(&newClaim)
	return serializer.BuildAdminLoginResponse(token, super_user)
}

func AdminLoginWithPassword(phone string, code string, password string) interface{} {
	super_user, err := dao.GetAdminByPhone(phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildFailResponseWithCode("not registered", 1)
	}
	if super_user.IsPass == 0 {
		return serializer.BuildFailResponseWithCode("等待审核", 2)
	}
	if super_user.Password != middleware.Md5Crypt(password, global.SlatSuperUserPassword) {
		return serializer.BuildFailResponseWithCode("密码错误", 3)
	}
	privilege := 1
	if super_user.IsSuper == 1 {
		privilege = 2
	}
	newClaim := middleware.UserClaims{Id: super_user.UserId, Phone: phone, Privilege: privilege}
	token, _ := middleware.GenerateToken(&newClaim)
	return serializer.BuildAdminLoginResponse(token, super_user)
}

func AdminPasswordReset(phone string, code string, password string) interface{} {
	super_user, err := dao.GetAdminByPhone(phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.BuildFailResponseWithCode("not registered", 1)
	}
	if super_user.IsPass == 0 {
		return serializer.BuildFailResponseWithCode("等待审核", 2)
	}
	err = VerifSmsCodeWithPhone(phone, code, global.TemplateCodeForLogin)
	if err != nil {
		return serializer.BuildFailResponseWithCode(err.Error(), e.ErrorCode)
	}
	err = dao.UpdateAdminPassword(super_user, middleware.Md5Crypt(password, global.SlatSuperUserPassword))
	if err != nil {
		global.Log.WithError(err).Error("AdminPasswordReset")
		return serializer.BuildErrorResponse(e.ErrorDatabase)
	}
	return serializer.BuildSuccessResponse("password reset success")
}
func GetPendingAccount() interface{} {
	super_users, err := dao.GetPendingAccount()
	if err != nil {
		return serializer.BuildErrorResponse(e.ErrorDatabase)
	}
	return serializer.BUildGetPendingAccountResponse(super_users)
}

func AuditAccount(phone string, is_pass int) interface{} {
	err := dao.AuditAccount(phone, is_pass)
	if err != nil {
		global.Log.WithError(err).Error("audit account")
		return serializer.BuildErrorResponse(e.ErrorDatabase)
	}
	return serializer.BuildAuditAccount()
}
