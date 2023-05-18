package api

import (
	"errors"
	"net/http"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

func GetSystemMessage(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetSystemMessage(claim.Id)
	c.JSON(http.StatusOK, response)
}

type MessageForm struct {
	MessageTitle   string `json:"message_title"`
	MessageDetails string `json:"message_details"`
	MessageDate    string `json:"message_date"`
	TargetType     int    `json:"target_type"`
	TargetId       string `json:"target_id"`
	TargetPosterId string `json:"target_poster_id"`
	TargetPostName string `json:"target_poster_name"`
	TargetStatus   int    `json:"target_status"`
	UserId         string `json:"user_id"`
	UserName       string `json:"user_name"`
}

func PostSystemMessage(c *gin.Context) {
	var message MessageForm
	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	targetId, _ := strconv.ParseInt(message.TargetId, 10, 64)
	posterId, _ := strconv.ParseInt(message.TargetPosterId, 10, 64)
	systemMessage := models.SystemMessage{
		ID:             global.Worker.GetId(),
		MessageTitle:   message.MessageTitle,
		MessageDetails: message.MessageDetails,
		MessageDate:    message.MessageDate,
		TargetType:     message.TargetType,
		TargetId:       targetId,
		PosterId:       posterId,
		PostName:       message.TargetPostName,
		TargetStatus:   message.TargetStatus,
		UserName:       message.UserName,
		MessageStatus:  0,
	}
	systemMessage.UserId, _ = strconv.ParseInt(message.UserId, 10, 64)
	response := service.PostSystemMessage(systemMessage)
	c.JSON(http.StatusOK, response)
}

type OfficialMessageForm struct {
	MessageTitle   string `json:"message_title"`
	MessageDetails string `json:"message_details"`
	MessageDate    string `json:"message_date"`
	MessageId      int64  `json:"message_id"`
	PosterName     string `json:"poster_name"`
}

func PostOfficialMessage(c *gin.Context) {
	var message OfficialMessageForm
	if err := c.ShouldBind(&message); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	officialMessage := models.OfficialMessage{
		ID:             global.Worker.GetId(),
		MessageTitle:   message.MessageTitle,
		MessageDetails: message.MessageDetails,
		MessageDate:    message.MessageDate,
		PosterName:     message.PosterName,
	}
	response := service.PostOfficialMessage(officialMessage)
	c.JSON(http.StatusOK, response)
}

func GetOfficialMessage(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetOfficialMessage(claim.Id)
	c.JSON(http.StatusOK, response)
}

// prefix 00_ for SystemMessage, 01_ for InteractiveMessage
func handlerMessageId(message_id string) (int64, uint64, error) {
	if len(message_id) < 3 {
		return 0, 0, errors.New("message id error")
	}
	m_type, err := strconv.ParseUint(message_id[0:2], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	m_id, err := strconv.ParseInt(message_id[3:], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return m_id, m_type, nil
}

type Delete struct {
	MessageId string `json:"message_id"`
}

func DeleteMessage(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	var deleted Delete
	if err := c.ShouldBind(&deleted); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	m_id, m_type, err := handlerMessageId(deleted.MessageId)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	var response interface{}
	switch m_type {
	case 00:
		//SystemMessage
		response = service.DeleteSystemMessage(m_id)
	case 01:
		//InteractiveMessage
		response = service.DeleteInteractiveMessage(m_id)
	case 02:
		//OfficialMessage
		response = service.DeleteOfficialMessage(m_id, claim.Id)
	default:
		response = ErrorResponse(errors.New("unsupported message prefix"))
	}
	c.JSON(http.StatusOK, response)
}

type InteractiveForm struct {
	MessageDate     string `json:"message_date"`      // 发送消息的日期
	MessageDetails  string `json:"message_details"`   // 消息的具体内容
	ObjectID        string `json:"object_id"`         // 该对象的id
	ObjectTypeID    int64  `json:"object_type_id"`    // 用户发送互动消息的对象种类id，1-商品 ，2-任务，3-公益活动，4-一级评论，5-二级评论
	PosterHeadPhoto string `json:"poster_head_photo"` // 发送消息的用户头像
	PosterID        int64  `json:"poster_id"`         // 发送消息的用户id
	PosterName      string `json:"poster_name"`       // 发送消息的用户名称
	UserID          string `json:"user_id"`           // 发送互动消息的目标用户id
}

func PostInteractiveMessage(c *gin.Context) {
	var form InteractiveForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	object_id, err := strconv.ParseInt(form.ObjectID, 10, 64)
	if err != nil {
		global.Log.WithError(err).Error("PostInteractiveMessage")
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
	message := models.InteractiveMessage{
		MessageDate:     form.MessageDate,
		MessageDetails:  form.MessageDetails,
		ObjectID:        object_id,
		ObjectTypeID:    form.ObjectTypeID,
		PosterID:        form.PosterID,
		UserID:          form.UserID,
		PosterName:      form.PosterName,
		PosterHeadPhoto: form.PosterHeadPhoto,
		ID:              global.Worker.GetId(),
	}
	response := service.PostInteractiveMessage(message)
	c.JSON(http.StatusOK, response)
}

func GetInteractiveMessage(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetInteractiveMessage(claim.Id)
	c.JSON(http.StatusOK, response)
}

func GetMessageLength(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetMessageLength(claim.Id)
	c.JSON(http.StatusOK, response)
}

type Read struct {
	MessageId string `json:"message_id"`
}

func PostMessageRead(c *gin.Context) {
	//userClaim, exist := c.Get("userClaim")
	//if !exist {
	//	c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
	//	return
	//}
	//claim := userClaim.(*middleware.UserClaims)
	var read Read
	if err := c.ShouldBind(&read); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	m_id, m_type, err := handlerMessageId(read.MessageId)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	var response interface{}
	switch m_type {
	case 00:
		//SystemMessage
		response = service.ReadSystemMessage(m_id)
	case 01:
		//InteractiveMessage
		response = service.ReadInteractiveMessage(m_id)
	default:
		response = ErrorResponse(errors.New("unsupported message prefix"))
	}
	c.JSON(http.StatusOK, response)
}
