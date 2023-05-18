package service

//import (
//	"strconv"
//	"xueyigou_demo/dao"
//	"xueyigou_demo/global"
//	"xueyigou_demo/middleware"
//	"xueyigou_demo/model"
//	"xueyigou_demo/models"
//	"xueyigou_demo/serializer"
//)
//
//func GetWorkInfo(Work_id int64) serializer.WorkInfoResponse {
//	work, works := models.WorkModel.GetWorkInfo(Work_id)
//
//	dao.AddWorkView(work.ID)
//	workView := dao.WorkView(work.ID) / 3
//	return serializer.BuildWorkResponse(work, works, workView)
//}
//
//func GetWorkMidInfo(Work_id int64) serializer.WorkInfoResponse {
//	work := dao.GetWorkMidInfo(Work_id)
//	return serializer.BuildWorkMidResponse(work)
//}
//
//func GetWorkListForMe(user_id int64) *serializer.WorkListResponseForMe {
//	_, err := dao.GetUserById(user_id)
//	if err == nil {
//		Work_list := dao.GetWorkListForMe(user_id)
//		var workViewList []uint64
//		for _, work := range Work_list {
//			workViewList = append(workViewList, dao.WorkView(work.ID))
//		}
//		ret := serializer.BuildWorkListResponseForme(Work_list, workViewList)
//		return &ret
//	}
//	return &serializer.WorkListResponseForMe{
//		Response: serializer.BuildUserDoesNotExitResponse("get work list"),
//	}
//}
//
//func GetWorkMidList() serializer.WorkMidListResponse {
//	workMidList := dao.GetWorkMidList()
//	return serializer.BuildWorkMidListResponse(workMidList)
//}
//func GetWorkSliceList(status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.WorkListResponse {
//	var work_list []models.Work
//	var count int64
//	if status == 0 {
//		work_list, count = dao.GetWorkSliceListByTime(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
//	} else if status == 1 {
//		work_list, count = dao.GetWorkSliceListByLike(currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
//	}
//	var workViewList []uint64
//	for _, work := range work_list {
//		workViewList = append(workViewList, dao.WorkView(work.ID))
//	}
//	return serializer.BuildWorkSliceListResponse(work_list, workViewList, count)
//}
//func GetWorkSliceLictWithSearch(keywords string, status int, currentPage int, pageSize int, searchChoiceFirst string, searchChoiceSecond string) serializer.WorkListResponse {
//	var work_list []models.Work
//	var count int64
//	if status == 0 {
//		work_list, count = dao.GetWorkSliceListWithSearchByTime(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
//	} else if status == 1 {
//		work_list, count = dao.GetWorkSliceListWithSearchByLike(keywords, currentPage, pageSize, searchChoiceFirst, searchChoiceSecond)
//	}
//	var workViewList []uint64
//	for _, work := range work_list {
//		workViewList = append(workViewList, dao.WorkView(work.ID))
//	}
//	return serializer.BuildWorkSliceListWithSearchResponse(work_list, workViewList, count)
//}
//func AddWork(Work models.WorkMid, Claim *middleware.UserClaims) *serializer.AddWorkResponse {
//	user, err := dao.GetUserById(Claim.Id)
//	if err == nil {
//		Work.PostUser = user
//		Work.PostUserId = Claim.Id
//		Work.ID = global.Worker.GetId()
//		dao.AddWork(Work)
//		ret := serializer.BUildAddWorkResponse(Work.ID)
//		return &ret
//	}
//	return &serializer.AddWorkResponse{
//		Response: serializer.BuildUserDoesNotExitResponse("add work"),
//	}
//}
//
//func PostWorkMid(workId int64, ispassed int64) interface{} {
//	if ispassed == 1 {
//		workMid := dao.GetWorkMidById(workId)
//		Work := models.Work{
//			Introduce: workMid.Introduce,
//			WorkDetails:   workMid.WorkDetails,
//			Title:      workMid.Title,
//			WorkType:      workMid.WorkType,
//			WorkDonation:  workMid.WorkDonation,
//			WorkPercent:   workMid.WorkPercent,
//			WorkMax:       workMid.WorkMax,
//			WorkMin:       workMid.WorkMin,
//			PostUser:      workMid.PostUser,
//			PostUserId:    workMid.PostUserId,
//			ID:            workId,
//			CoverPicture:   workMid.CoverPicture,
//		}
//		for _, pic := range workMid.PicturesUrlList {
//			Work.PicturesUrlList = append(Work.PicturesUrlList, models.WorkPicturesUrl{Url: pic.Url})
//		}
//		for _, subject := range workMid.WorkSubject {
//			Work.WorkSubject = append(Work.WorkSubject, models.WorkSubjectItem{Item: subject.Item})
//		}
//		dao.AddWorkMid(Work)
//
//	}
//	//图片就不用删了
//	if _, err := dao.DeleteWorkMid(workId); err == nil {
//		//TODO fs删除图片
//		// var urls_ []string
//		// for _, item := range urls {
//		// 	urls_ = append(urls_, item.Url)
//		// }
//		// if res, err1 := Deletefile(urls_); err1 != nil {
//		// 	return res
//		// }
//		//
//		return serializer.BuildSuccessResponse("Delete work success")
//	} else {
//		return serializer.BuildFailResponse("Delete work failed")
//	}
//	return serializer.BuildSuccessResponse("audit success")
//}
//
//func AddWorkPictureUrl(work_id int64, urls []string, p_urls []string) interface{} {
//	var pic_urls []models.WorkPicturesUrl
//	for _, url := range urls {
//		pic_urls = append(pic_urls, models.WorkPicturesUrl{Url: url})
//	}
//	err := dao.AddWorkPictureUrl(work_id, pic_urls)
//	if err == nil {
//		return serializer.BuildAddWorkPicturesUrlResponse(p_urls)
//	}
//	return serializer.BuildFailResponse("work picture error")
//}
//
//func GetWorkSliceLictWithEvent(event_id string, currentPage int, pageSize int) serializer.WorkListResponse {
//	var workID_list []models.WorkSubjectItem
//	var count int64
//
//	T, _ := strconv.ParseInt(event_id, 10, 64)
//	event, _ := dao.GetEvent(T)
//	workID_list, count = dao.GetWorksIDListWithSearchByEvent(event.EventsName, currentPage, pageSize)
//
//	var work_list []models.Work
//	for i := range workID_list {
//		work_list = append(work_list, dao.GetWorkById(workID_list[i].WorkId))
//	}
//
//	var workViewList []uint64
//	for _, work := range work_list {
//		workViewList = append(workViewList, dao.WorkView(work.ID))
//	}
//	return serializer.BuildWorkSliceListWithSearchResponse(work_list, workViewList, count)
//}
