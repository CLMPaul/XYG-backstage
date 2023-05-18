package serializer

import (
	"strconv"
	"time"
	"xueyigou_demo/models"
)

type UserRegisterResponse struct {
	ResultStatus int32       `json:"result_status"`
	ResultMsg    string      `json:"result_msg,omitempty"`
	User         models.User `json:"user"`
}

// 登录与注册
type UserLoginResponse struct {
	ResultStatus int32  `json:"result_status"`
	ResultMsg    string `json:"result_msg,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	Token        Token  `json:"token,omitempty"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserAttention struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Sign        string `json:"sign,omitempty"`
	HeaderPhoto string `json:"header_photo,omitempty"`
}

type UserAttentionResponse struct {
	ResultStatus int32           `json:"result_status"`
	ResultMsg    string          `json:"result_msg,omitempty"`
	UserList     []UserAttention `json:"user_list,omitempty"`
}

type UserFansResponse struct {
	ResultStatus int32           `json:"result_status"`
	ResultMsg    string          `json:"result_msg,omitempty"`
	UserList     []UserAttention `json:"user_list,omitempty"`
}

type GetUserInfoResponse struct {
	UserName            string   `json:"user_name"`
	UserHeadPhoto       string   `json:"user_header_photo"`
	UserBackgroundPhoto string   `json:"user_background_photo"`
	UserGender          string   `json:"user_gender"`
	UserHometown        string   `json:"user_hometown"`
	UserProvince        string   `json:"user_province"`
	UserCity            string   `json:"user_city"`
	UserSign            string   `json:"user_sign"`
	UserGraduate        string   `json:"user_graduate"`
	UserSubject         string   `json:"user_subject"`
	UserQQ              string   `json:"user_qq"`
	Userdetail          string   `json:"user_detail"`
	UserWechat          string   `json:"user_wechat"`
	UserLabels          []string `json:"user_labels"`
	Response
}
type GoodsInfo struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	GoodsName    string `json:"goodsName,omitempty"`
	GoodsDate    string `json:"goodsDate,omitempty"`
	GoodsMsg     string `json:"goodsMsg,omitempty"`
	GoodsPicture string `json:"goodsPicture,omitempty"`
	PosterName   string `json:"posterName,omitempty"`
}

type UserGoodsListResponse struct {
	ResultStatus int32        `json:"result_status"`
	ResultMsg    string       `json:"result_msg,omitempty"`
	GoodsList    []*GoodsInfo `json:"goods_list,omitempty"`
}

type UserInfoResponse struct {
	Response
	UserInfo
}
type UserInfo struct {
	UserId          string              `json:"user_id"`
	Name            string              `json:"user_name"`
	Sex             string              `json:"sex"`
	Sign            string              `json:"sign"`
	HeaderPhoto     string              `json:"header_photo"`
	BackgroundPhoto string              `json:"background_photo"`
	Telephone       string              `json:"telephone"`
	Email           string              `json:"email"`
	Wechat          string              `json:"wechat"`
	LabelList       []models.UserLables `json:"label_list"`
	FollowCount     int64               `json:"follow_count"`
	FollowerCount   int64               `json:"follower_count"`
	IsFollow        bool                `json:"is_follow"`
	Level           uint                `json:"level"`
	Hot             uint64              `json:"hot"`
	Contribution    uint64              `json:"contribution"`
	Medals          uint                `json:"medals"`
	Deals           int64               `json:"deals"`
	WorkCount       int64               `json:"work_count"`
	LikeCount       int64               `json:"like_count"`
}

type GetUserAddressResponse struct {
	Response
	AddressList []addressInfo `json:"address_list"`
}

type addressInfo struct {
	ReceiverAddress string `json:"receiver_address"` // 收件人地址详细信息
	ReceiverName    string `json:"receiver_name"`    // 收件人姓名
	ReceiverStatus  int64  `json:"receiver_status"`  // 该地址的状态，0-默认，其他值-非默认
	ReceiverTelenum string `json:"receiver_telenum"` // 收件人电话号码
}

type PostAddressResponse struct {
	Response
}

func BuildUserInfoResponse(user models.User, followCount, followerCount int64, isFollow bool, work_cnt, like_cnt int64) UserInfoResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get UserInfo success",
	}
	userInfo := UserInfo{
		UserId:          strconv.FormatInt(user.UserId, 10),
		Name:            user.Name,
		Sex:             string(user.Sex),
		Sign:            user.Sign,
		HeaderPhoto:     user.HeaderPhoto,
		BackgroundPhoto: user.BackgroundPhoto,
		Telephone:       user.Telephone,
		Email:           user.Email,
		Wechat:          user.Wechat,
		LabelList:       user.LabelList,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        isFollow,
		Level:           user.Level,
		Hot:             user.Hot,
		Medals:          user.Medals,
		Contribution:    user.Contribution,
		Deals:           user.Deals,
		LikeCount:       like_cnt,
		WorkCount:       work_cnt,
	}
	userInfoResponse := UserInfoResponse{
		response,
		userInfo,
	}
	return userInfoResponse
}

func BuildGetAddress(addresses []models.Useraddress) GetUserAddressResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get UserAddrees success",
	}
	var address_list []addressInfo
	for _, address := range addresses {
		item := addressInfo{
			ReceiverAddress: address.AddressDetails,
			ReceiverName:    address.AddressName,
			ReceiverStatus:  address.AddressStatus,
			ReceiverTelenum: address.AddressTelenum,
		}
		address_list = append(address_list, item)
	}
	return GetUserAddressResponse{
		Response:    response,
		AddressList: address_list,
	}
}

func BuildPostAddress() PostAddressResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "post Useraddress success",
	}
	return PostAddressResponse{
		Response: response,
	}
}

type GetIndentityResponse struct {
	Response
	IndentityItem1
}

type IndentityItem struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserId       int64
	UserTelenum  string  `json:"user_telenum"`         // 用户电话号码
	UserIdentity string  `json:"user_identity"`        // 学生身份认证，0-在校学生，1-非在校学生
	UserIdnum    string  `json:"user_idnum"`           // 用户学生证号
	UserName     string  `json:"user_name"`            // 用户姓名
	UserPhoto    *string `json:"user_photo,omitempty"` // 用户学生证照片地址url
}

// for return
type IndentityItem1 struct {
	UserId       string
	UserTelenum  string  `json:"user_telenum"`         // 用户电话号码
	UserIdentity string  `json:"user_identity"`        // 学生身份认证，0-在校学生，1-非在校学生
	UserIdnum    string  `json:"user_idnum"`           // 用户学生证号
	UserName     string  `json:"user_name"`            // 用户姓名
	UserPhoto    *string `json:"user_photo,omitempty"` // 用户学生证照片地址url
}

func BuildGetIndentityResponse(indentity IndentityItem) GetIndentityResponse {
	item := IndentityItem1{
		UserId:       strconv.FormatInt(indentity.UserId, 10),
		UserTelenum:  indentity.UserTelenum,
		UserIdentity: indentity.UserIdentity,
		UserIdnum:    indentity.UserIdnum,
		UserName:     indentity.UserName,
		UserPhoto:    indentity.UserPhoto,
	}
	return GetIndentityResponse{
		Response:       BuildSuccessResponse("Get Indentity"),
		IndentityItem1: item,
	}
}

type PostIndentityResponse struct {
	Response
}

type PostPeopleResponse struct {
	Response
}

func BuildPostIndentityResponse() PostIndentityResponse {
	return PostIndentityResponse{
		Response: BuildSuccessResponse("post indentity"),
	}
}

func BuildPostPeopleResponse() PostPeopleResponse {
	return PostPeopleResponse{
		Response: BuildSuccessResponse("post people"),
	}
}

func BuildGetUserDetailInfoResponse(user models.User, labels []models.UserLables) GetUserInfoResponse {
	var labelslist []string
	for _, t := range labels {
		labelslist = append(labelslist, t.Label)
	}
	return GetUserInfoResponse{
		UserName:            user.Name,
		UserHeadPhoto:       user.HeaderPhoto,
		UserBackgroundPhoto: user.BackgroundPhoto,
		UserGender:          user.Sex,
		UserHometown:        user.Hometown,
		UserProvince:        user.Province,
		UserCity:            user.City,
		UserSign:            user.Sign,
		UserGraduate:        user.Graduate,
		UserSubject:         user.Subject,
		UserQQ:              user.QQ,
		Userdetail:          user.Detail,
		UserWechat:          user.Wechat,
		UserLabels:          labelslist,
		Response: Response{
			ResultStatus: 0,
			ResultMsg:    "success get user info",
		},
	}
}

type OrderTaskListResponse struct {
	Response
	OrderTaskIdList []int64 `json:"task_id_list"` // 用户接单的任务
}

func BuildGetOrderTaskListResponse(taskIds []int64) OrderTaskListResponse {
	return OrderTaskListResponse{
		BuildSuccessResponse("Get OrderList Success"),
		taskIds,
	}
}

type OrderTaskResponse struct {
	Response
}

type FinishTaskResponse struct {
	Response
}

func BuildOrderTaskResponse() OrderTaskResponse {
	return OrderTaskResponse{
		Response: BuildSuccessResponse("Order Task"),
	}
}

func BuildFinishTaskResponse() FinishTaskResponse {
	return FinishTaskResponse{
		Response: BuildSuccessResponse("Finish Task"),
	}
}

type OrderWorkListResponse struct {
	Response
	OrderWorkIdList []int64 `json:"work_id_list"`
}

func BuildGetOrderWorkListResponse(workIds []int64) OrderWorkListResponse {
	return OrderWorkListResponse{
		BuildSuccessResponse("Get OrderList Success"),
		workIds,
	}
}

type MyPeopleResponse struct {
	Response
	models.People `json:"people"`
}

func BuildMyPeople(people models.People) MyPeopleResponse {
	return MyPeopleResponse{
		BuildSuccessResponse("Get People Success"),
		people,
	}
}
