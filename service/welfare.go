package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"
)

func GetWelfareInfo(userId int64) serializer.WelfareListResponseForCollection {
	welfares := dao.GetWelfareInfoForCollection(userId)
	var welfareViewList []uint64
	for _, welfare := range welfares {
		welfareViewList = append(welfareViewList, dao.WelfareView(welfare.ID))
	}
	return serializer.BuildWelfareInfoResponseForCollection(welfares, welfareViewList)
}

func GetMyWelfareInfo(userId int64) serializer.MyWelfareListResponse {
	welfares := dao.GetMyWelfareList(userId)
	return serializer.BuildMyWelfareInfoResponse(welfares)
}
func GetWelfareNewsList() serializer.WelfareListResponse {
	welfares := dao.GetWelfareNewsList()
	var welfareViewList []uint64
	for _, welfare := range welfares {
		welfareViewList = append(welfareViewList, dao.WelfareView(welfare.ID))
	}
	return serializer.BuildWelfareListResponse(welfares, welfareViewList)
}

func GetWelfareActivity(wellfareId int64) serializer.WelfareActivityResponse {
	welfare := dao.GetWelfareActivity(wellfareId)
	return serializer.BuildWelfareAcyivityResponse(welfare)
}

func PostWelfareInfo(welfare models.Welfare) serializer.Response {
	err := dao.PostWelfareInfo(welfare)
	var response serializer.Response
	if err != nil {
		response.ResultMsg = "post welfareInfo failed"
		response.ResultStatus = 1
	} else {
		response.ResultMsg = "post welfareInfo success"
		response.ResultStatus = 0
	}
	return response
}

func GetWelfareActivityInfo(welfareId int64) serializer.WelfareResponse {
	welfare := dao.GetWelfareById(welfareId)
	//welfareView := dao.WelfareView(welfare.ID)
	//dao.AddWelfareView(welfare.ID)
	//welfare.WelfareLikes = dao.Likes_get(welfareId, 4)
	//welfare.WelfareLikes = dao.GetWelfareLikeCount(welfareId)
	welfareCollection := dao.GetWelfareCollection(welfareId)
	//pictures := dao.GetWelfarePicturesUrl(welfareId)
	return serializer.BuildWelafreInfoResponse(welfare, welfareCollection)
}

func GetWelfareHistory() serializer.WelfareHistoryResponse {
	welfareCount, welfarePeople := dao.GetWelfareHistory()
	return serializer.BuildWelfareHistoryResponse(welfareCount, welfarePeople)
}

func JoinWelfare(welfareId int64, userId int64) interface{} {
	welfareMember := models.WelfareMember{
		WelfareId: welfareId,
		UserId:    userId,
		Status:    0,
	}
	if err := dao.JoinWelfare(welfareMember); err != nil {
		return serializer.BuildFailResponse(err.Error())
	}
	return serializer.BuildSuccessResponse("Join Success")
}

func GetWelfarePeople(welfareId int64) interface{} {
	welfare := dao.GetWelfareById(welfareId)
	indentities := dao.GetWelfarePeople(welfareId)
	return serializer.BuildGetWelfarePeopleResponse(welfare, indentities)
}

func AddWelfareInfo(userId []int64, welfareId int64, welfareTime int) interface{} {
	if err := dao.AddWelfareInfo(userId, welfareId, welfareTime); err != nil {
		return serializer.BuildFailResponse("add failed")
	}
	return serializer.BuildSuccessResponse("add success")
}

func PostWelfareModify(welfare models.Welfare, welfareId int64, urls []models.WelfarePictureUrl) interface{} {
	if err := dao.ModifyWelfare(welfare, welfareId, urls); err != nil {
		return serializer.BuildFailResponse("modify failed")
	}
	return serializer.BuildSuccessResponse("modify success")
}
func PostWelfareDelete(welfareId int64) interface{} {
	if err := dao.DeleteWelfare(welfareId); err != nil {
		return serializer.BuildFailResponse("delete failed")
	}
	return serializer.BuildSuccessResponse("delete success")
}
