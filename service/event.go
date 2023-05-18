package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
)

func EventSubmit(event models.Event) interface{} {
	if dao.AddEvent(event) != nil {
		return serializer.BuildFailResponse("event submit fail")
	}
	return serializer.BuildSuccessResponse("event submit success")
}

func EventDelete(eventid int64) interface{} {
	if dao.DeleteEvent(eventid) != nil {
		return serializer.BuildFailResponse("event delete fail")
	}
	return serializer.BuildSuccessResponse("event delete success")
}

func EventGet(eventid int64) interface{} {
	event, err := dao.GetEvent(eventid)
	if err != nil {
		return serializer.BuildFailResponse("event get fail")
	}
	return serializer.BuildEventResponse(event)
}

func EventGetSelf(eventid int64) interface{} {
	event, err := dao.GetEvent(eventid)
	if err != nil {
		return serializer.BuildFailResponse("event get fail")
	}
	//likeCount := dao.GetEventLikeCount(eventid)
	return serializer.BuildEventsLefResponse(event, 0)
}
func EventGetList() interface{} {
	eventlist, err := dao.GetEventList()
	if err != nil {
		return serializer.BuildFailResponse("event get fail")
	}
	return serializer.BuildEventListResponse(eventlist)
}
func GetCollectionEventsList(userId int64) serializer.EventsListResponse {
	eventsList := dao.GetCollectionEventsList(userId)
	return serializer.BuildEventsListResponse(eventsList)
}

func GetEventsListBySearch(isOnline int, isSchool int, isBegin int, keyword string, currentPage int, pageSize int) serializer.EventsListResponse {
	eventsList := dao.GetEventsListBySearch(isOnline, isSchool, isBegin, keyword, currentPage, pageSize)
	return serializer.BuildEventsListResponse(eventsList)
}

func GetEventMembers(eventid int64) interface{} {
	members := dao.GetEventMembers(eventid)
	return serializer.BuildEventMembersResponse(members)
}

func PostActivitySign(eventid int64, id string, phone string, name string) interface{} {
	err := dao.PostActivitySign(eventid, id, phone, name)
	if err != nil {
		return serializer.BuildFailResponse(" fail")
	}
	return serializer.BuildSuccessResponse("success")

}
