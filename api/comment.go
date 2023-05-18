package api

import (
	"errors"
	"net/http"
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CommentActionForm struct {
	ActionType          string  `json:"action_type"`               // 评论操作，1-发布评论，2-删除评论
	CommentContent      *string `json:"comment_content,omitempty"` // 用户填写的评论内容（action_type =1）
	CommentID           *string `json:"comment_id,omitempty"`      // 要删除的评论id（action_type = 2）
	UserID              *string `json:"user_id,omitempty"`         // 发布二级评论时被评论的用户id
	ObjectId            string  `json:"object_id"`                 // 作品id
	ResponseToCommentId *string `json:"response_to_comment_id"`    //发布二级评论时被评论的comment_id
}

// 添加评论
func CommentAction(c *gin.Context) {
	var form CommentActionForm
	var user_id int64
	var comment_id int64
	var comment_level bool //true for first level, false for second level
	var response_to_comment_id string
	var content string
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be CommentActionForm")
		return
	}
	global.Log.WithFields(logrus.Fields{
		"ActionType":          form.ActionType,
		"CommentContent":      form.CommentContent,
		"CommentID":           form.CommentID,
		"UserID":              form.UserID,
		"ObjectId":            form.ObjectId,
		"ResponseToCommentId": form.ResponseToCommentId,
	}).Info()
	actionType := form.ActionType
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)

	if form.CommentContent == nil {
		content = ""
	} else {
		content = *form.CommentContent
	}

	if form.UserID == nil {
		user_id = 0
	} else {
		user_id, _ = strconv.ParseInt(*form.UserID, 10, 64)
	}

	if form.ResponseToCommentId == nil {
		response_to_comment_id = ""
	} else {
		response_to_comment_id = *form.ResponseToCommentId
	}

	//handle object id
	object_type, object_id, err := HandleObjectId(form.ObjectId)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	//end handle

	//handle comment id
	if form.CommentID == nil {
		comment_id = 0
	} else {
		if len(*form.CommentID) <= 2 {
			c.JSON(http.StatusOK, ErrorResponse(errors.New("参数错误")))
			return
		}
		switch (*form.CommentID)[0:2] {
		case "f_":
			comment_level = true
		case "s_":
			comment_level = false
		default:
			c.JSON(http.StatusOK, ErrorResponse(errors.New("参数错误")))
			return
		}
		cid_string := (*form.CommentID)[2:]
		var err error
		comment_id, err = strconv.ParseInt(cid_string, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(errors.New("参数错误")))
			return
		}
	}
	//end handle
	comment_service := service.CommentService{
		Claim:      claim,
		Actiontype: actionType,
		// Content:             global.IllegalWords.Replace(content, '*'),
		Content:             content,
		WorkId:              object_id,
		Commentid:           comment_id,
		Userid:              user_id,
		ObjectType:          object_type,
		CommentLevel:        comment_level,
		ResponseToCommentId: response_to_comment_id,
	}

	response := comment_service.CommentAction()
	c.JSON(http.StatusOK, response)
}

// 获取评论
func CommentList(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	o_id := c.Query("object_id")
	object_type, object_id, err := HandleObjectId(o_id)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	comment_service := service.CommentService{
		Claim:      claim,
		WorkId:     object_id,
		ObjectType: object_type,
	}
	comment_list := comment_service.GetCommentList()
	c.JSON(http.StatusOK, comment_list)
}
