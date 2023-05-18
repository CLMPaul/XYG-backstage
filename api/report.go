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

type ReportForm struct {
	SubjectItem   int    `json:"subject_item"`
	SubjectId     int64  `json:"subject_id"`
	ReportReason  int    `json:"report_reason"`
	ReportDetails string `json:"report_details"`
}

func PostReport(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	var r ReportForm
	if err := c.ShouldBind(&r); err != nil {
		c.String(http.StatusOK, "Body should be ReportForm")
		return
	}
	report := models.Report{
		ID:            global.Worker.GetId(),
		SubjectItem:   r.SubjectItem,
		SubjectId:     r.SubjectId,
		ReportReason:  r.ReportReason,
		ReportDetails: r.ReportDetails,
		ReporterId:    claim.Id,
	}
	response := service.PostReport(report)
	c.JSON(http.StatusOK, response)
}

func GetReportPeopleList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchChoiceFirst := c.Query("searchChoiceFirst")
	keywords := c.Query("keywords")
	searchChoiceSecond := c.Query("searchChoiceSecond")

	if keywords == "" { //无搜索返回作品
		response := service.GetReportPeopleList(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	} else { //返回搜索作品
		response := service.GetReportPeopleListWithSearch(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	}
}

func GetReportTaskList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchChoiceFirst := c.Query("searchChoiceFirst")
	keywords := c.Query("keywords")
	searchChoiceSecond := c.Query("searchChoiceSecond")

	if keywords == "" { //无搜索返回作品
		response := service.GetReportTaskList(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	} else { //返回搜索作品
		response := service.GetReportTaskListWithSearch(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	}
}

func GetReportWorkList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	searchChoiceSecond := c.Query("searchChoiceSecond")
	keywords := c.Query("keywords")
	searchChoiceFirst := c.Query("searchChoiceFirst")
	if keywords == "" { //无搜索返回作品
		response := service.GetReportWorkList(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	} else { //返回搜索作品
		response := service.GetReportWorkListWithSearch(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
		c.JSON(http.StatusOK, response)
	}
	//c.JSON(http.StatusOK, ErrorResponse(errors.New("query params error")))
}

func PostReportResult(c *gin.Context) {
	item_type, _ := strconv.ParseInt(c.Query("type"), 10, 64)
	item_id, _ := strconv.ParseInt(c.Query("item_id"), 10, 64)
	result, _ := strconv.ParseInt(c.Query("result"), 10, 64)
	report_id, _ := strconv.ParseInt(c.Query("report_id"), 10, 64)

	response := service.PostReportResult(item_type, item_id, result, report_id)
	c.JSON(http.StatusOK, response)
}
