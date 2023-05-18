package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
)

func GetSystemMessage(userId int64) interface{} {
	messages, err := dao.GetMessageList(userId)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildSystemMessageListResponse(messages)
}

func PostSystemMessage(message models.SystemMessage) serializer.Response {
	user, _ := dao.GetUserById(message.UserId)
	//detect sensitive words
	// if global.IllegalWords.ContainsAny(message.MessageDetails) {
	// 	message.MessageDetails = global.IllegalWords.Replace(message.MessageDetails, '*')
	// }
	message.UserTelephone = user.Telephone
	err := dao.PostSystemMessage(message)

	var response serializer.Response
	if err != nil {
		response.ResultMsg = "post system message failed"
		response.ResultStatus = 1
	} else {
		response.ResultMsg = "post system message success"
		response.ResultStatus = 0
	}
	return response
}

func PostOfficialMessage(message models.OfficialMessage) serializer.Response {
	err := dao.PostOfficialMessage(message)
	var response serializer.Response
	if err != nil {
		response.ResultMsg = "post official message failed"
		response.ResultStatus = 1
	} else {
		response.ResultMsg = "post official message success"
		response.ResultStatus = 0
	}
	return response
}
func GetOfficialMessage(userId int64) interface{} {
	messages, err := dao.GetOfficialMessageList(userId)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildOfficialMessageListResponse(messages)
}
func DeleteSystemMessage(messageId int64) interface{} {
	err := dao.DeleteSystemMessage(messageId)
	var response serializer.Response
	if err != nil {
		response.ResultMsg = "delete system message failed"
		response.ResultStatus = 1
	} else {
		response.ResultMsg = "delete system message success"
		response.ResultStatus = 0
	}
	return response
}
func DeleteOfficialMessage(messageId int64, userId int64) interface{} {
	err := dao.DeleteOfficialMessage(messageId, userId)
	var response serializer.Response
	if err != nil {
		response.ResultMsg = "delete official message failed"
		response.ResultStatus = 1
	} else {
		response.ResultMsg = "delete official message success"
		response.ResultStatus = 0
	}
	return response
}
func GetInteractiveMessage(user_id int64) interface{} {
	message_list, err := dao.Getinteractivemessagelist(user_id)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildGetInteractiveMessagelistResponse(message_list)
}

func PostInteractiveMessage(message models.InteractiveMessage) interface{} {
	err := dao.PostInteractiveMessage(message)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildPostInteractiveMessagelistResponse()
}

func GetMessageLength(user_id int64) interface{} {
	MessageLength, err := dao.GetMessageLength(user_id)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildGetMessageLengthResponse(MessageLength)
}

func DeleteInteractiveMessage(message_id int64) interface{} {
	err := dao.DeleteInteractiveMessage(message_id)
	if err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildDeleteInteractiveMessageResponse()
}
func ReadSystemMessage(messageId int64) interface{} {
	if err := dao.SystemMessageRead(messageId); err != nil {
		return serializer.BuildFailResponse("set system message read failed")
	}
	return serializer.BuildSuccessResponse("set system message read success")
}
func ReadInteractiveMessage(messageId int64) interface{} {
	if err := dao.InteractiveMessageRead(messageId); err != nil {
		return serializer.BuildFailResponse("set interactive message read failed")
	}
	return serializer.BuildSuccessResponse("set interactive message read success")
}
