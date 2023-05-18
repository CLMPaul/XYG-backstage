package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xueyigou_demo/middleware"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"
)

type CollectionForm struct {
	ActionType   int    `json:"action_type"` // 收藏操作，1-收藏，2-取消收藏
	ObjectId     string `json:"object_id"`   // 作品id
	ObjectTypeId string `json:"object_type"` //对象类型id,1-评论 ，2-二级评论(不同对象要严格区分)
}

func Collection(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	userId := claim.Id
	var form CollectionForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be CollectionForm")
		return
	}
	actiontype := form.ActionType
	//userClaim, _ := c.Get("userClaim")
	//claim := userClaim.(*middleware.UserClaims)
	objectId, _ := strconv.Atoi(form.ObjectId)
	switch actiontype {
	case 0:
		res := service.Collection(userId, int64(objectId), form.ObjectTypeId)
		c.JSON(200, res)
	case 1:
		res := service.CancelCollection(userId, int64(objectId), form.ObjectTypeId)
		c.JSON(200, res)
	default:
		res := serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "action_type should be 1 or 2",
		}
		c.JSON(http.StatusNotFound, res)
	}
}

func GetCollectionTaskList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	response := service.GetCollectionTaskList(int64(userId))
	c.JSON(http.StatusOK, response)
}

func GetCollectionWorkList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	response := service.GetCollectionWorkList(int64(userId))
	c.JSON(http.StatusOK, response)
}

// 小程序收藏接口
func AppletCollection(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	userId := claim.Id
	var form CollectionForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be CollectionForm")
		return
	}
	actiontype := form.ActionType
	objectId, _ := strconv.Atoi(form.ObjectId)
	switch actiontype {
	case 0:
		res := service.AppletCollection(userId, int64(objectId), form.ObjectTypeId)
		c.JSON(200, res)
	case 1:
		res := service.AppletCancelCollection(userId, int64(objectId), form.ObjectTypeId)
		c.JSON(200, res)
	default:
		res := serializer.Response{
			ResultStatus: 1,
			ResultMsg:    "action_type should be 0 or 1",
		}
		c.JSON(http.StatusNotFound, res)
	}
}
