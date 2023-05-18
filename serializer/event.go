package serializer

import (
	"strconv"
	"xueyigou_demo/models"
)

type EventsMain struct {
	EventsBeforejoin  string   `json:"events_beforejoin"`
	EventsDetails     string   `json:"events_details"`
	EventsName        string   `json:"events_name"`
	EventsPictureList []string `json:"events_picture_list"`
	EventsQA          string   `json:"events_QA"`
	EventsRewards     string   `json:"events_rewards"`
	EventsRules       string   `json:"events_rules"`
	EventsWorkRequire string   `json:"events_work_require"`
}
type EventResponse struct {
	ResultStatus int32  `json:"result_status"`
	ResultMsg    string `json:"result_msg,omitempty"`
	EventsMain   `json:"events_main"`
}

func BuildEventResponse(event models.Event) EventResponse {
	var urlList []string
	for _, picture := range event.EventsPictureList {
		urlList = append(urlList, picture.Url)
	}
	eventMain := EventsMain{
		EventsBeforejoin:  event.EventsBeforejoin,
		EventsDetails:     event.EventsDetails,
		EventsName:        event.EventsName,
		EventsPictureList: urlList,
		EventsQA:          event.EventsQA,
		EventsRewards:     event.EventsRewards,
		EventsRules:       event.EventsRules,
		EventsWorkRequire: event.EventsWorkRequire,
	}
	response := EventResponse{
		ResultStatus: 0,
		ResultMsg:    "get events_main success",
		EventsMain:   eventMain,
	}
	return response
}

type EventSlefResponse struct {
	ResultStatus int32  `json:"result_status"`
	ResultMsg    string `json:"result_msg,omitempty"`
	EventsInfo   `json:"events_info"`
}

// 活动信息
type EventsInfo struct {
	ConnectorInfo    Connector `json:"connector_info"` // 联系人信息
	EventsCollectnum int64     `json:"events_collectnum"`
	EventsEndDate    string    `json:"events_end_date"` // 活动结束时间
	EventsLikenum    int64     `json:"events_likenum"`
	EventsPoster     string    `json:"events_poster"`     // 活动举办方
	EventsStartDate  string    `json:"events_start_date"` // 活动开始时间
}

// 联系人信息
type Connector struct {
	ConnectorName      string `json:"connector_name"`
	ConnectorTelephone string `json:"connector_telephone"`
	ConnectorQrCode    string `json:"connector_qr_code"`
}

func BuildEventsLefResponse(event models.Event, likeCount int64) EventSlefResponse {
	connector := Connector{
		ConnectorName:      event.ConnectorName,
		ConnectorTelephone: event.ConnectorTelephone,
		ConnectorQrCode:    event.ConnectorQrCode,
	}
	eventsInfo := EventsInfo{
		ConnectorInfo:    connector,
		EventsCollectnum: event.EventCollection,
		EventsEndDate:    event.EventsEndDate,
		EventsLikenum:    likeCount,
		EventsPoster:     event.EventsPoster,
		EventsStartDate:  event.EventsStartDate,
	}
	eventSlefResponse := EventSlefResponse{
		ResultStatus: 0,
		ResultMsg:    "get events left success",
		EventsInfo:   eventsInfo,
	}
	return eventSlefResponse
}

type Events struct {
	EventsCover     string `json:"events_cover"`      // 活动封面图片
	EventsDetails   string `json:"events_details"`    // 活动简介
	EventsEndDate   string `json:"events_end_date"`   // 活动结束日期
	EventsID        string `json:"events_id"`         // 活动id
	EventsName      string `json:"events_name"`       // 活动名称
	EventsPoster    string `json:"events_poster"`     // 活动举办方
	EventsStartDate string `json:"events_start_date"` // 活动开始日期
}
type EventListResponse struct {
	ResultStatus int32    `json:"result_status"`
	ResultMsg    string   `json:"result_msg,omitempty"`
	EventsList   []Events `json:"events_list"`
}

func BuildEventListResponse(eventList []models.Event) EventListResponse {
	var eventsList []Events
	for _, event := range eventList {
		eventId := strconv.FormatInt(event.EventId, 10)
		events := Events{
			EventsCover:     event.EventsCover,
			EventsDetails:   event.EventsDetails,
			EventsEndDate:   event.EventsEndDate,
			EventsID:        eventId,
			EventsName:      event.EventsName,
			EventsPoster:    event.EventsPoster,
			EventsStartDate: event.EventsStartDate,
		}
		eventsList = append(eventsList, events)
	}
	response := EventListResponse{
		ResultStatus: 0,
		ResultMsg:    "get events list success",
		EventsList:   eventsList,
	}
	return response
}

type EventsListResponse struct {
	ResultStatus         int          `json:"result_status"`
	ResultMsg            string       `json:"result_msg"`
	CollectionEventsList []EventsList `json:"collection_events_list"`
}
type EventsList struct {
	EventsCover     string `json:"events_cover"`      // 活动封面
	EventsDetails   string `json:"events_details"`    // 活动信息
	EventsEndDate   string `json:"events_end_date"`   // 活动结束日期
	EventsID        string `json:"events_id"`         // 活动id
	EventsName      string `json:"events_name"`       // 活动名称
	EventsPoster    string `json:"events_poster"`     // 活动举办方
	EventsStartDate string `json:"events_start_date"` // 活动开始日期
}

func BuildEventsListResponse(eventsList []models.Event) EventsListResponse {
	var list []EventsList
	for _, event := range eventsList {
		id := strconv.FormatInt(event.EventId, 10)
		events := EventsList{
			EventsCover:     event.EventsCover,
			EventsDetails:   event.EventsDetails,
			EventsEndDate:   event.EventsEndDate,
			EventsID:        id,
			EventsName:      event.EventsName,
			EventsPoster:    event.EventsPoster,
			EventsStartDate: event.EventsStartDate,
		}
		list = append(list, events)
	}
	response := EventsListResponse{
		ResultStatus:         0,
		ResultMsg:            "get events list success",
		CollectionEventsList: list,
	}
	return response
}

type EventMemberList struct {
	id       string `json:"id"`       // 头像
	Phone    string `json:"phone"`    // 手机号
	Username string `json:"username"` // 用户名
}
type EventMembersresponse struct {
	MemberList   []EventMemberList `json:"memberList"`
	ResultMsg    string            `json:"result_msg"`
	ResultStatus int64             `json:"result_status"`
}

func BuildEventMembersResponse(members []models.EventMember) EventMembersresponse {
	var list []EventMemberList
	for _, member := range members {
		member_list := EventMemberList{
			id:       member.UserID,
			Phone:    member.Phone,
			Username: member.UserName,
		}
		list = append(list, member_list)
	}
	response := EventMembersresponse{
		ResultStatus: 0,
		ResultMsg:    "get events members success",
		MemberList:   list,
	}
	return response
}
