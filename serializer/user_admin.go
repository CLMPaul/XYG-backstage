package serializer

import (
	"strconv"
	"xueyigou_demo/models"
)

type AdminRegisterResponse struct {
	Response
}

func BuildAdminRegisterResponse() *AdminRegisterResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "等待审核",
	}
	return &AdminRegisterResponse{
		Response: response,
	}
}

type AdminLoginResponse struct {
	Response
	Token   string `json:"token"`
	IsSuper bool   `json:"isSuper"`
	UserId  string `json:"user_id"`
}

func BuildAdminLoginResponse(token string, user *models.SuperUser) *AdminLoginResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "success",
	}
	return &AdminLoginResponse{
		Response: response,
		Token:    token,
		IsSuper:  user.IsSuper == 1,
		UserId:   strconv.FormatInt(user.UserId, 10),
	}
}

type AuditAccountResponse struct {
	Response
}

func BuildAuditAccount() *AuditAccountResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "success",
	}
	return &AuditAccountResponse{
		Response: response,
	}
}

type GetPendingAccountResponse struct {
	Response
	UserList []pendingAccountList
}

type pendingAccountList struct {
	PhoneNumber string `json:"phonenumber"`
	Name        string `json:"name"`
}

func BUildGetPendingAccountResponse(super_user []models.SuperUser) GetPendingAccountResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "success",
	}
	var phone_list []pendingAccountList
	for _, user := range super_user {
		phone_list = append(phone_list, pendingAccountList{user.Telephone, user.Name})
	}
	return GetPendingAccountResponse{
		Response: response,
		UserList: phone_list,
	}
}
