package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"
)

func GetWelfareList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	response := service.GetWelfareInfo(userId)
	c.JSON(http.StatusOK, response)
}
func GetMyWelfareList(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetMyWelfareInfo(claim.Id)
	c.JSON(http.StatusOK, response)
}
func GetWelfareNewsList(c *gin.Context) {
	response := service.GetWelfareNewsList()
	c.JSON(http.StatusOK, response)
}

func WelfareActivity(c *gin.Context) {
	welfareId, _ := strconv.Atoi(c.Query("welfare_id"))
	response := service.GetWelfareActivity(int64(welfareId))
	c.JSON(http.StatusOK, response)
}

type WelfareIntro struct {
	WelfareName          string   `json:"welfare_name,omitempty"`
	WelfarePicture       string   `json:"welfare_picture,omitempty"`
	WelfareDetails       string   `json:"welfare_details,omitempty"`
	WelfareInfo          string   `json:"welfare_info,omitempty"`
	WelfareJoin          int64    `json:"welfare_join,omitempty"`
	WelfareStatus        int      `json:"welfare_status,omitempty"`
	WelfareAddress       string   `json:"welfare_address,omitempty"`
	StartDate            string   `json:"start_date,omitempty"`
	EndDate              string   `json:"end_date,omitempty"`
	RecruitsPeople       string   `json:"recruits_people,omitempty"`
	UserType             int      `json:"user_type,omitempty"`
	ConnectorName        string   `json:"connector_name"`
	ConnectorTelephone   string   `json:"connector_telephone"`
	ConnectorQrCodePhoto string   `json:"connector_qr_code_photo"`
	WelfareDate          string   `json:"welfare_date"`
	PicturesUrlList      []string `json:"pictures_url_list"`
}

func PostWelfareInfo(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	var w WelfareIntro
	if err := c.ShouldBind(&w); err != nil {
		c.String(http.StatusOK, "Body should be CollectionForm")
		return
	}
	welfare := models.Welfare{
		ID:                   global.Worker.GetId(),
		WelfareInfo:          w.WelfareInfo,
		WelfareJoin:          w.WelfareJoin,
		WelfareDetails:       w.WelfareDetails,
		WelfarePicture:       w.WelfarePicture,
		WelfareName:          w.WelfareName,
		WelfareAddress:       w.WelfareAddress,
		StartDate:            w.StartDate,
		EndDate:              w.EndDate,
		RecruitsPeople:       w.RecruitsPeople,
		ConnectorName:        w.ConnectorName,
		ConnectorQrCodePhoto: w.ConnectorQrCodePhoto,
		ConnectorTelephone:   w.ConnectorTelephone,
	}
	for _, url := range w.PicturesUrlList {
		welfare.PicturesUrlList = append(welfare.PicturesUrlList, models.WelfarePictureUrl{Url: url})
	}
	welfare.PostUserId = claim.Id
	response := service.PostWelfareInfo(welfare)
	if response.ResultStatus == 1 {
		serializer.BuildUserDoesNotExitResponse("post welfareInfo")
	}
	c.JSON(http.StatusOK, response)
}

func GetWelfareActivityInfo(c *gin.Context) {
	welfareId, _ := strconv.Atoi(c.Query("welfare_id"))
	response := service.GetWelfareActivityInfo(int64(welfareId))
	c.JSON(http.StatusOK, response)
}
func GetWelfareHistory(c *gin.Context) {
	response := service.GetWelfareHistory()
	c.JSON(http.StatusOK, response)
}

func JoinWelfare(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	welfareId, _ := strconv.ParseInt(c.Query("welfare_id"), 10, 64)
	response := service.JoinWelfare(welfareId, claim.Id)
	c.JSON(http.StatusOK, response)
}

func GetWelfarePeople(c *gin.Context) {
	welfareId, _ := strconv.ParseInt(c.Query("welfare_id"), 10, 64)
	response := service.GetWelfarePeople(welfareId)
	c.JSON(http.StatusOK, response)
}

type WelfareInfo struct {
	UserId      []int64 `json:"user_id"`
	WelfareId   int64   `json:"welfare_id"`
	WelfareTime int     `json:"welfare_time"`
}

func AddWelfareInfo(c *gin.Context) {
	var form WelfareInfo
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be CollectionForm")
		return
	}
	response := service.AddWelfareInfo(form.UserId, form.WelfareId, form.WelfareTime)
	c.JSON(http.StatusOK, response)
}

type ModifyInfo struct {
	WelfareId            int64    `json:"welfare_id"`
	ConnectorName        string   `json:"connector_name"`
	ConnectorQrCodePhoto string   `json:"connector_qr_code_photo"`
	ConnectorTelephone   string   `json:"connector_telephone"`
	WelfareAddress       string   `json:"welfare_address"`
	WelfareDetails       string   `json:"welfare_details"`
	WelfareEndDate       string   `json:"welfare_end_date"`
	WelfareJoin          int64    `json:"welfare_join"`
	WelfareName          string   `json:"welfare_name"`
	WelfarePictureList   []string `json:"welfare_picture_list"`
	WelfareStartDate     string   `json:"welfare_start_date"`
}

func PostWelfareModify(c *gin.Context) {
	var m ModifyInfo
	if err := c.ShouldBind(&m); err != nil {
		c.String(http.StatusOK, "Body should be ModifyForm")
		return
	}
	welfare := models.Welfare{
		ConnectorName:        m.ConnectorName,
		ConnectorQrCodePhoto: m.ConnectorQrCodePhoto,
		ConnectorTelephone:   m.ConnectorTelephone,
		WelfareAddress:       m.WelfareAddress,
		WelfareDetails:       m.WelfareDetails,
		EndDate:              m.WelfareEndDate,
		WelfareJoin:          m.WelfareJoin,
		WelfareName:          m.WelfareName,
		StartDate:            m.WelfareStartDate,
	}
	var urls []models.WelfarePictureUrl
	for _, url := range m.WelfarePictureList {
		urls = append(urls, models.WelfarePictureUrl{Url: url, WelfareId: m.WelfareId})
	}
	response := service.PostWelfareModify(welfare, m.WelfareId, urls)
	c.JSON(http.StatusOK, response)
}

type DeleteWelfare struct {
	WelfareId int64 `json:"welfare_id"`
}

func PostWelfareDelete(c *gin.Context) {
	var d DeleteWelfare
	if err := c.ShouldBind(&d); err != nil {
		c.String(http.StatusOK, "Body should be DeleteForm")
		return
	}
	response := service.PostWelfareDelete(d.WelfareId)
	c.JSON(http.StatusOK, response)
}
