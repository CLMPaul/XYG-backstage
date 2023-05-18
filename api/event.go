package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/service"
)

func PostEventSubmit(c *gin.Context) {
	var eventform models.Eventform
	if err := c.ShouldBind(&eventform); err == nil {
		event := models.Event{
			EventId:            global.Worker.GetId(),
			ConnectorID:        eventform.ConnectorID,
			ConnectorName:      eventform.ConnectorName,
			ConnectorQrCode:    eventform.ConnectorQrCode,
			ConnectorTelephone: eventform.ConnectorTelephone,
			EventsBeforejoin:   eventform.EventsBeforejoin,
			EventsCover:        eventform.EventsCover,
			EventsDetails:      eventform.EventsDetails,
			EventsEndDate:      eventform.EventsEndDate,
			EventsName:         eventform.EventsName,
			EventsPictureList:  nil,
			EventsPoster:       eventform.EventsPoster,
			EventsQA:           eventform.EventsQA,
			EventsRewards:      eventform.EventsRewards,
			EventsRules:        eventform.EventsRules,
			EventsStartDate:    eventform.EventsStartDate,
			EventsWorkRequire:  eventform.EventsWorkRequire,
			IsOnline:           eventform.IsOnline,
			IsSchool:           eventform.IsSchool,
		}
		event.EventsPictureList = make([]models.Eventpicurl, 0)
		for _, i := range eventform.EventsPictureList {
			event.EventsPictureList = append(event.EventsPictureList, models.Eventpicurl{
				Url:     i,
				EventId: event.EventId,
			})
		}

		res := service.EventSubmit(event)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func PostEventEdit(c *gin.Context) { //TODO：较为低效实现update
	var eventform models.Eventform
	if err := c.ShouldBind(&eventform); err == nil {
		T, _ := strconv.ParseInt(eventform.EventId, 10, 64)
		service.EventDelete(T)
		event := models.Event{
			EventId:            T,
			ConnectorID:        eventform.ConnectorID,
			ConnectorName:      eventform.ConnectorName,
			ConnectorQrCode:    eventform.ConnectorQrCode,
			ConnectorTelephone: eventform.ConnectorTelephone,
			EventsBeforejoin:   eventform.EventsBeforejoin,
			EventsCover:        eventform.EventsCover,
			EventsDetails:      eventform.EventsDetails,
			EventsEndDate:      eventform.EventsEndDate,
			EventsName:         eventform.EventsName,
			EventsPictureList:  nil,
			EventsPoster:       eventform.EventsPoster,
			EventsQA:           eventform.EventsQA,
			EventsRewards:      eventform.EventsRewards,
			EventsRules:        eventform.EventsRules,
			EventsStartDate:    eventform.EventsStartDate,
			EventsWorkRequire:  eventform.EventsWorkRequire,
		}
		event.EventsPictureList = make([]models.Eventpicurl, 0)
		for _, i := range eventform.EventsPictureList {
			event.EventsPictureList = append(event.EventsPictureList, models.Eventpicurl{
				Url: i,
			})
		}
		res := service.EventSubmit(event)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
func PostEventDelete(c *gin.Context) {
	var eventform models.EventDeleteform
	if err := c.ShouldBind(&eventform); err == nil {
		T, _ := strconv.ParseInt(eventform.EventId, 10, 64)
		res := service.EventDelete(T)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
func GetEvent(c *gin.Context) {
	eventid := c.Query("event_id")
	T, _ := strconv.ParseInt(eventid, 10, 64)
	//fmt.Println("----------------", T)
	res := service.EventGet(T)
	c.JSON(200, res)
}
func GetEventSelf(c *gin.Context) {
	eventid := c.Query("event_id")
	T, _ := strconv.ParseInt(eventid, 10, 64)
	//fmt.Println("----------------", T)
	res := service.EventGetSelf(T)
	c.JSON(200, res)
}
func GetEventList(c *gin.Context) {
	res := service.EventGetList()
	c.JSON(200, res)
}
func GetCollectionEventsList(c *gin.Context) {
	Id := c.Query("user_id")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	res := service.GetCollectionEventsList(userId)
	c.JSON(200, res)
}

func ActivitiyList(c *gin.Context) {
	isOnline, _ := strconv.Atoi(c.Query("isOnline"))
	isSchool, _ := strconv.Atoi(c.Query("isSchool"))
	isBegin, _ := strconv.Atoi(c.Query("isBegin"))
	keyword := c.Query("keyword")
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	res := service.GetEventsListBySearch(isOnline, isSchool, isBegin, keyword, currentPage, pageSize)
	c.JSON(200, res)
}

func GetMembers(c *gin.Context) {
	eventid := c.Query("event_id")
	T, _ := strconv.ParseInt(eventid, 10, 64)
	res := service.GetEventMembers(T)
	c.JSON(200, res)
}

func PostActivitySign(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	eventid := c.Query("event_id")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)

	T, _ := strconv.ParseInt(eventid, 10, 64)
	res := service.PostActivitySign(T, strconv.FormatInt(claim.Id, 10), claim.Phone, claim.Name)
	c.JSON(200, res)
}
