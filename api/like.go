package api

import (
	"errors"
	"net/http"
	"strconv"
	"xueyigou_demo/middleware"
	"xueyigou_demo/service"

	"github.com/gin-gonic/gin"
)

type LikeForm struct {
	ActionType int    `json:"action_type"` // 点赞操作，0-点赞，1-取消点赞
	WorkID     string `json:"object_id"`   // 作品id
	ObjectType string `json:"object_type"` //对象类型id,1-评论 ，2-二级评论(不同对象要严格区分)
}

// 网站点赞接口
func Like(c *gin.Context) {
	var form LikeForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be LikeForm")
		return
	}
	actiontype := form.ActionType
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	work_id, _ := strconv.ParseInt(form.WorkID, 10, 64)
	//fmt.Println(dao.Likes_get(work_id, 0))
	switch actiontype {
	case 0:
		res := service.Like(claim.Id, work_id, form.ObjectType)
		c.JSON(200, res)
	case 1:
		res := service.Dislike(claim.Id, work_id, form.ObjectType)
		c.JSON(200, res)
	}
}

func GetLikeList(c *gin.Context) {
	userClaim, exist := c.Get("userClaim")
	objectTypeId, _ := strconv.Atoi(c.Query("object_type_id"))
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	response := service.GetLikeList(claim.Id, int64(objectTypeId))
	c.JSON(http.StatusOK, response)
}

// 小程序点赞接口
func AppletLike(c *gin.Context) {
	var form LikeForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusOK, "Body should be LikeForm")
		return
	}
	actiontype := form.ActionType
	userClaim, exist := c.Get("userClaim")
	if !exist {
		c.JSON(http.StatusOK, ErrorResponse(errors.New("user claim not in token")))
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	work_id, _ := strconv.ParseInt(form.WorkID, 10, 64)
	switch actiontype {
	case 0:
		res := service.AppletLike(claim.Id, work_id, form.ObjectType)
		c.JSON(200, res)
	case 1:
		res := service.AppletDislike(claim.Id, work_id, form.ObjectType)
		c.JSON(200, res)
	}
}
