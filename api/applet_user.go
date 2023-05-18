package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xueyigou_demo/internal/web"
	"xueyigou_demo/middleware"
	"xueyigou_demo/proto"
	"xueyigou_demo/serializer"
	"xueyigou_demo/service"
)

func GetUserActivities(c *gin.Context) {
	var req proto.UserActivityRequest
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	resp := serializer.BuildAppletActivityPagedResp(service.AppletActivitySrv.GetPagedByUser(req))
	c.JSON(http.StatusOK, resp)
}

func GetUserWorks(c *gin.Context) {
	var req proto.WorkRequest
	err := c.BindQuery(&req)
	userClaim, exist := c.Get("userClaim")
	if !exist {
		web.BadRequestResponse(c, "user claim not in token")
		return
	}
	claim := userClaim.(*middleware.UserClaims)
	data, err := service.WorkService.GetPagedByUser(req, claim.Id)
	if err != nil {
		web.HandleError(c, err)
		return
	}
	web.SuccessResponse(c, data)
}

func UserFollowAction(c *gin.Context) {
	var req proto.UserFollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	resp := service.ParseAttentionAction(req)
	c.JSON(http.StatusOK, resp)
}
