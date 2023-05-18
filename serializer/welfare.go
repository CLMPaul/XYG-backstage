package serializer

import (
	"fmt"
	"strconv"
	"time"
	"xueyigou_demo/models"
)

type WelfareListResponseForCollection struct {
	Response
	WelfareList []WelfareInfoForCollection `json:"collection_welfare_list"`
}
type WelfareInfoForCollection struct {
	WelfareId      string `json:"welfare_id,omitempty"`
	WelfareName    string `json:"welfare_name,omitempty"`
	WelfarePicture string `json:"welfare_picture"`
	WelfareDetails string `json:"welfare_details"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	WelfareView    uint64 `json:"welfare_view"`
	WelfareJoin    int64  `json:"welfare_join"`
	WelfareStatus  int    `json:"welfare_status"`
}

func BuildWelfareInfoResponseForCollection(welfares []models.Welfare, welfareViewList []uint64) WelfareListResponseForCollection {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get collection welfare list success",
	}
	var welfareList []WelfareInfoForCollection
	for i, welfare := range welfares {
		timeLayoutStr := "2006-01-02 15:04:05"
		startDate, _ := time.Parse(timeLayoutStr, welfare.StartDate)
		endDate, _ := time.Parse(timeLayoutStr, welfare.EndDate)
		t := time.Now()
		var welfareStatus int
		if t.Before(startDate) {
			welfareStatus = 0
		} else if t.After(startDate) && t.Before(endDate) {
			welfareStatus = 1
		} else {
			welfareStatus = 2
		}
		welfareInfoForCollection := WelfareInfoForCollection{
			WelfareId:      strconv.FormatInt(welfare.ID, 10),
			WelfareName:    welfare.WelfareName,
			WelfarePicture: welfare.WelfarePicture,
			WelfareDetails: welfare.WelfareDetails,
			StartDate:      welfare.StartDate,
			EndDate:        welfare.EndDate,
			WelfareView:    welfareViewList[i],
			WelfareJoin:    welfare.WelfareJoin,
			WelfareStatus:  welfareStatus,
		}
		welfareList = append(welfareList, welfareInfoForCollection)
	}
	welfare_list_response := WelfareListResponseForCollection{
		Response:    response,
		WelfareList: welfareList,
	}
	return welfare_list_response
}

type MyWelfareListResponse struct {
	Response
	WelfareList []MyWelfareListInfo `json:"welfare_list"`
}

type MyWelfareListInfo struct {
	WelfareId      string    `json:"welfare_id,omitempty"`
	WelfareName    string    `json:"welfare_name,omitempty"`
	WelfarePicture string    `json:"welfare_picture,omitempty"`
	WelfareDetails string    `json:"welfare_details,omitempty"`
	WelfareJoin    int64     `json:"welfare_join,omitempty"`
	WelfareStatus  int       `json:"welfare_status,omitempty"`
	StartDate      string    `json:"start_date,omitempty"`
	EndDate        string    `json:"end_date,omitempty"`
	WelfareView    int64     `json:"welfare_view"`
	CreatedTime    time.Time `json:"created_time"` //创建时间
}

func BuildMyWelfareInfoResponse(welfares []models.Welfare) MyWelfareListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get Mywelfare list success",
	}
	var welfareList []MyWelfareListInfo
	for _, welfare := range welfares {
		//timeLayoutStr := "2006-01-02 15:04:05"
		//startDate, _ := time.Parse(timeLayoutStr, welfare.StartDate)
		//endDate, _ := time.Parse(timeLayoutStr, welfare.EndDate)
		//t := time.Now()
		//var welfareStatus int
		//if t.Before(startDate) {
		//	welfareStatus = 0
		//} else if t.After(startDate) && t.Before(endDate) {
		//	welfareStatus = 1
		//} else {
		//	welfareStatus = 2
		//}
		welfareInfo := MyWelfareListInfo{
			WelfareId:      strconv.FormatInt(welfare.ID, 10),
			WelfareName:    welfare.WelfareName,
			WelfarePicture: welfare.WelfarePicture,
			WelfareDetails: welfare.WelfareDetails,
			WelfareJoin:    welfare.WelfareJoin,
			StartDate:      welfare.StartDate,
			EndDate:        welfare.EndDate,
			WelfareView:    welfare.WelfareView,
			CreatedTime:    welfare.CreatedAt,
		}
		welfareList = append(welfareList, welfareInfo)
	}
	welfare_list_response := MyWelfareListResponse{
		Response:    response,
		WelfareList: welfareList,
	}
	return welfare_list_response
}

type WelfareListResponse struct {
	Response
	WelfareList []WelfareNewsInfo `json:"welfare_list"`
}
type WelfareNewsInfo struct {
	WelfareId      string    `json:"welfare_id"`
	WelfareName    string    `json:"welfare_name"`
	WelfarePicture string    `json:"welfare_picture"`
	WelfareDetails string    `json:"welfare_details"`
	WelfareJoin    int64     `json:"welfare_join"`
	WelfareStatus  int       `json:"welfare_status"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	WelfareView    uint64    `json:"welfare_view"`
	CreatedTime    time.Time `json:"created_time"` //创建时间
}

func BuildWelfareListResponse(welfares []models.Welfare, welfareViewList []uint64) WelfareListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get welfare_news_list success",
	}
	var WelfareList []WelfareNewsInfo
	for i, welfare := range welfares {
		timeLayoutStr := "2006-01-02 15:04:05"
		startDate, _ := time.Parse(timeLayoutStr, welfare.StartDate)
		endDate, _ := time.Parse(timeLayoutStr, welfare.EndDate)
		t := time.Now()
		var welfareStatus int
		if t.Before(startDate) {
			welfareStatus = 0
		} else if t.After(startDate) && t.Before(endDate) {
			welfareStatus = 1
		} else {
			welfareStatus = 2
		}
		welfareNewsInfo := WelfareNewsInfo{
			WelfareId:      strconv.FormatInt(welfare.ID, 10),
			WelfareName:    welfare.WelfareName,
			WelfarePicture: welfare.WelfarePicture,
			WelfareDetails: welfare.WelfareDetails,
			WelfareJoin:    welfare.WelfareJoin,
			WelfareStatus:  welfareStatus,
			StartDate:      welfare.StartDate,
			EndDate:        welfare.EndDate,
			WelfareView:    welfareViewList[i],
			CreatedTime:    welfare.CreatedAt,
		}
		WelfareList = append(WelfareList, welfareNewsInfo)
	}
	welfareListResponse := WelfareListResponse{
		response,
		WelfareList,
	}
	return welfareListResponse
}

type WelfareActivityResponse struct {
	Response
	StartDate      string        `json:"start_date,omitempty"`
	EndDate        string        `json:"finish_date,omitempty"`
	WelfareId      string        `json:"welfare_id"`
	RecruitsPeople string        `json:"recruits_people,omitempty"`
	PosterInfo     PosterInfo    `json:"poster_info"`
	ConnectorInfo  ConnectorInfo `json:"connector_info"`
	CreatedTime    time.Time     `json:"created_time"` //创建时间
}

type PosterInfo struct {
	PosterId    string `json:"poster_id"`
	PosterName  string `json:"poster_name"`
	PosterPhoto string `json:"poster_photo"`
}

type ConnectorInfo struct {
	ConnectName string `json:"connect_name"`
	Telephone   string `json:"telephone"`
	QrCodePhoto string `json:"qr_code_photo"`
}

func BuildWelfareAcyivityResponse(welfare models.Welfare) WelfareActivityResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get welfare_activity success",
	}
	posterInfo := PosterInfo{
		PosterId:    strconv.FormatInt(welfare.PostUserId, 10),
		PosterPhoto: welfare.PostUser.HeaderPhoto,
		PosterName:  welfare.PostUser.Name,
	}
	connectorInfo := ConnectorInfo{
		ConnectName: welfare.ConnectorName,
		Telephone:   welfare.ConnectorTelephone,
		QrCodePhoto: welfare.ConnectorQrCodePhoto,
	}
	welfareActivityResponse := WelfareActivityResponse{
		Response:       response,
		WelfareId:      strconv.FormatInt(welfare.ID, 10),
		StartDate:      welfare.StartDate,
		EndDate:        welfare.EndDate,
		RecruitsPeople: welfare.RecruitsPeople,
		PosterInfo:     posterInfo,
		ConnectorInfo:  connectorInfo,
		CreatedTime:    welfare.CreatedAt,
	}
	return welfareActivityResponse
}

type WelfareResponse struct {
	Response
	WelfareList `json:"welfare_list"`
}
type WelfareList struct {
	WelfareInfo       string    `json:"welfare_info"`
	WelfareId         string    `json:"welfare_id"`
	WelfareName       string    `json:"welfare_name"`
	PicturesUrlList   []string  `json:"pictures_url_list"`
	WelfareDetails    string    `json:"welfare_details"`
	WelfareDate       string    `json:"welfare_date"`
	WelfareView       uint64    `json:"welfare_view"`
	WelfareCollection int64     `json:"welfare_collection"`
	WelfareLike       int64     `json:"welfare_like"`
	CreatedTime       time.Time `json:"created_time"` //创建时间
}

func BuildWelafreInfoResponse(welfare models.Welfare, collectionCount int64) WelfareResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get welfare_activityInfo success",
	}
	var pictures []string
	for _, url := range welfare.PicturesUrlList {
		pictures = append(pictures, url.Url)
	}
	welfareList := WelfareList{
		WelfareInfo:       welfare.WelfareInfo,
		WelfareId:         strconv.FormatInt(welfare.ID, 10),
		WelfareName:       welfare.WelfareName,
		PicturesUrlList:   pictures,
		WelfareDetails:    welfare.WelfareDetails,
		WelfareCollection: collectionCount,
		CreatedTime:       welfare.CreatedAt,
		WelfareLike:       welfare.WelfareLikes,
	}
	welfareInfoResponse := WelfareResponse{
		Response:    response,
		WelfareList: welfareList,
	}
	return welfareInfoResponse
}

type WelfareHistoryResponse struct {
	Response
	HistoryCount  int64 `json:"history_count"`
	HistoryPeople int64 `json:"history_people"`
}

func BuildWelfareHistoryResponse(welfareCount int64, welfarePeople int64) WelfareHistoryResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get welfare_history success",
	}
	welfareHistoryResponse := WelfareHistoryResponse{
		Response:      response,
		HistoryCount:  welfareCount,
		HistoryPeople: welfarePeople,
	}
	return welfareHistoryResponse
}

type MemberList []WelfareMember
type GetWelfarePeopleResponse struct {
	Response
	MemberList `json:"member_list"`
}

type WelfareMember struct {
	UserID              string `json:"user_id"`                        // 参加的用户id
	UserRealname        string `json:"user_realname"`                  // 参加的用户真实姓名
	UserTelephone       string `json:"user_telephone"`                 // 参加的用户电话号码
	WelfareContribution int64  `json:"welfare_contribution,omitempty"` // 参加的公益活动贡献值
	WelfareTime         int64  `json:"welfare_time,omitempty"`         // 参加的公益活动公益时长
}

func BuildGetWelfarePeopleResponse(welfare models.Welfare, indentities []models.Indentity) GetWelfarePeopleResponse {
	var members []WelfareMember
	for _, indentity := range indentities {
		member := WelfareMember{
			UserRealname:        indentity.UserName,
			UserTelephone:       indentity.UserTelenum,
			WelfareContribution: welfare.WelfareTime * 10,
			WelfareTime:         welfare.WelfareTime,
		}
		member.UserID = fmt.Sprint(indentity.UserId)
		members = append(members, member)
	}
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "Get WelfarePeopleList Success",
	}
	res := GetWelfarePeopleResponse{
		Response:   response,
		MemberList: members,
	}
	return res
}
