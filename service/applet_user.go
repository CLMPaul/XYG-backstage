package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/dao/applet"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
	"xueyigou_demo/serializer"
)

var AppletActivitySrv appletActivitySrv

type appletActivitySrv struct{}

func (appletActivitySrv) GetPagedByUser(req proto.UserActivityRequest) ([]models.Event, int64, error) {
	data, err := applet.ActivityDao.GetEventsByUser(req)
	if err != nil {
		return nil, 0, err
	}

	total := int64(len(data))
	return data, total, nil
}

func ParseAttentionAction(req proto.UserFollowRequest) serializer.Response {
	if req.ActionType == 1 {
		dao.AddAttention(req.UserId, req.OtherId)
		dao.AddFan(req.UserId, req.OtherId)
		return serializer.Response{ResultStatus: 1, ResultMsg: "Success"}
	}

	if req.ActionType == 2 {
		dao.DeleteAttention(req.UserId, req.OtherId)
		dao.DeleteFan(req.UserId, req.OtherId)
		return serializer.Response{ResultStatus: 1, ResultMsg: "Success"}
	}

	return serializer.Response{ResultStatus: 0, ResultMsg: "参数传入错误"}
}
