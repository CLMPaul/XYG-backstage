package serializer

import (
	"strconv"
	"xueyigou_demo/models"
)

type SystemMessageListResponse struct {
	Response
	SystemMessageList []SystemMessage
}

type SystemMessage struct {
	MessageId         string `json:"message_id"`
	MessageTitle      string `json:"message_title"`
	MessageDetails    string `json:"message_details"`
	MessageDate       string `json:"message_date"`
	MessageStatus     int    `json:"message_status"`
	TargetType        int    `json:"target_type"`
	TargetId          string `json:"target_id"`
	TargetStatus      int    `json:"target_status"`
	UserId            string `json:"user_id"`
	UserName          string `json:"user_name"`
	UserTelephone     string `json:"user_telephone"`
	MessageReadStatus int    `json:"message_read_status"`
	TargetPoster
}

type TargetPoster struct {
	PosterId string `json:"poster_id"`
	PostName string `json:"post_name"`
}

func BuildSystemMessageListResponse(messages []models.SystemMessage) SystemMessageListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get system message list success",
	}
	var messageList []SystemMessage
	for _, message := range messages {
		posterId := strconv.FormatInt(message.PosterId, 10)
		targetPoster := TargetPoster{
			PosterId: posterId,
			PostName: message.PostName,
		}
		messageId := strconv.FormatInt(message.ID, 10)
		targetId := strconv.FormatInt(message.TargetId, 10)
		userId := strconv.FormatInt(message.UserId, 10)
		systemMessage := SystemMessage{
			MessageId:         messageId,
			MessageTitle:      message.MessageTitle,
			MessageDetails:    message.MessageDetails,
			MessageDate:       message.MessageDate,
			MessageStatus:     message.MessageStatus,
			TargetType:        message.TargetType,
			TargetId:          targetId,
			TargetStatus:      message.TargetStatus,
			UserId:            userId,
			UserName:          message.UserName,
			UserTelephone:     message.UserTelephone,
			TargetPoster:      targetPoster,
			MessageReadStatus: message.MessageReadStatus,
		}
		messageList = append(messageList, systemMessage)
	}
	systemMessageListResponse := SystemMessageListResponse{
		response,
		messageList,
	}
	return systemMessageListResponse
}

type OfficialMessage struct {
	MessageId      int64  `json:"message_id"`
	MessageTitle   string `json:"message_title"`
	MessageDetails string `json:"message_details"`
	MessageDate    string `json:"message_date"`
	PosterName     string `json:"poster_name"`
}
type OfficialMessageListResponse struct {
	Response
	OfficialMessageList []OfficialMessage
}

func BuildOfficialMessageListResponse(messages []models.OfficialMessage) OfficialMessageListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get official message list success",
	}
	var messageList []OfficialMessage
	for _, message := range messages {
		officialMessage := OfficialMessage{
			MessageId:      message.ID,
			MessageTitle:   message.MessageTitle,
			MessageDetails: message.MessageDetails,
			MessageDate:    message.MessageDate,
			PosterName:     message.PosterName,
		}
		messageList = append(messageList, officialMessage)
	}
	officialMessageListResponse := OfficialMessageListResponse{
		response,
		messageList,
	}
	return officialMessageListResponse
}

type GetInteractiveMessagelistResponse struct {
	Response
	InteractiveMessageList []interactiveMessage `json:"eachmessage_list"`
}

type GetMessageLengthResponse struct {
	Response
	models.MessageLength `json:"message_length"`
}

type interactiveMessage struct {
	MessageId         string `json:"message_id"`
	MessageDate       string `json:"message_date"`        // 发送消息的日期
	MessageDetails    string `json:"message_details"`     // 消息的具体内容
	ObjectID          int64  `json:"object_id"`           // 该对象的id
	ObjectTypeID      int64  `json:"object_type_id"`      // 用户发送互动消息的对象种类id，1-商品 ，2-任务，3-公益活动，4-一级评论，5-二级评论
	PosterHeadPhoto   string `json:"poster_head_photo"`   // 用户头像地址url
	PosterID          int64  `json:"poster_id"`           // 用户id
	PosterName        string `json:"poster_name"`         // 用户名称
	UserID            string `json:"user_id"`             // 发送互动消息的目标用户id
	MessageReadStatus int    `json:"message_read_status"` //判断消息是否已读
}

func BuildGetInteractiveMessagelistResponse(messages []models.InteractiveMessage) GetInteractiveMessagelistResponse {
	response := BuildSuccessResponse("get interactive message")

	ret := GetInteractiveMessagelistResponse{
		Response: response,
	}
	for _, message := range messages {
		ret.InteractiveMessageList = append(ret.InteractiveMessageList, interactiveMessage{
			MessageDate:       message.MessageDate,
			MessageDetails:    message.MessageDetails,
			ObjectID:          message.ObjectID,
			ObjectTypeID:      message.ObjectTypeID,
			PosterHeadPhoto:   message.PosterHeadPhoto,
			PosterName:        message.PosterName,
			PosterID:          message.PosterID,
			UserID:            message.UserID,
			MessageId:         strconv.Itoa(int(message.ID)),
			MessageReadStatus: message.MessageReadStatus,
		})
	}
	return ret
}

type PostInteractiveMessagelistResponse struct {
	Response
}

func BuildPostInteractiveMessagelistResponse() PostInteractiveMessagelistResponse {
	response := BuildSuccessResponse("post interactive message")
	return PostInteractiveMessagelistResponse{
		Response: response,
	}
}

func BuildGetMessageLengthResponse(Length *models.MessageLength) GetMessageLengthResponse {
	response := BuildSuccessResponse("get message length success")
	return GetMessageLengthResponse{
		Response:      response,
		MessageLength: *Length,
	}
}

type DeleteInteractiveMessageResponse struct {
	Response
}

func BuildDeleteInteractiveMessageResponse() DeleteInteractiveMessageResponse {
	return DeleteInteractiveMessageResponse{
		Response: BuildSuccessResponse("DeleteInteractiveMessage"),
	}
}
