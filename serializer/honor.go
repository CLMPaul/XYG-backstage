package serializer

import (
	"xueyigou_demo/models"
)

type HonorListResponse struct {
	Response
	UserList []userList `json:"user_list"`
}

type AddUserToHonorListResponse struct {
	Response
}

type userList struct {
	UserContribution int64  `json:"user_contribution"` // 用户的公益贡献值
	UserHead         string `json:"user_head"`         // 用户头像的url
	UserID           int64  `json:"user_id"`           // 用户id
	UserMedal        int64  `json:"user_medal"`        // 用户获得的奖牌数
	UserName         string `json:"user_name"`         // 用户名称
}

func BuildHonorListResponse(honor_list []models.HonorList) HonorListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "Get Honor List Success",
	}
	honor_list_response := HonorListResponse{
		Response: response,
	}

	for _, item := range honor_list {
		u := userList{
			UserHead:         item.HeaderPhoto,
			UserID:           item.UserID,
			UserName:         item.Name,
			UserContribution: item.Contribution,
			UserMedal:        item.Medals,
		}
		honor_list_response.UserList = append(honor_list_response.UserList, u)
	}
	return honor_list_response
}
