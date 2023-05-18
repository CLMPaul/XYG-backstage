package service

import (
	"fmt"
	"xueyigou_demo/dao"
	"xueyigou_demo/models"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"
)

// 网站收藏接口
func Collection(userId int64, objectId int64, objectTypeId string) serializer.Response {
	switch objectTypeId {
	case "0": //对公益活动进行收藏
		if err := dao.AddWelfareCollection(objectId, userId); err == nil {
			if err = dao.AddWelfareCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Welfare Collection success")
			}
		}
		return serializer.BuildFailResponse("Welfare Collection failed")
	case "1": //对商品进行收藏
		if err := models.WorkModel.AddWorkCollection(objectId, userId); err == nil {
			if err = dao.AddWorkCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Work Collection success")
			}
		}
		return serializer.BuildFailResponse(" Work Collection failed")
	case "2": //对任务进行收藏
		if err := dao.AddTaskCollection(objectId, userId); err == nil {
			if err = dao.AddTaskCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Task Collection success")
			}
		}
		return serializer.BuildFailResponse("Task Collection failed")
	case "3": //对活动进行收藏
		if err := dao.AddEventCollection(objectId, userId); err == nil {
			if err = dao.AddEventCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Event Collection success")
			}
		}
		return serializer.BuildFailResponse("Event Collection failed")
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}

func CancelCollection(userId int64, objectId int64, objectTypeId string) serializer.Response {
	switch objectTypeId {
	case "0": //对公益活动进行取消收藏
		if err := dao.CancelWelfareCollection(objectId, userId); err == nil {
			if err = dao.CancelWelfareCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Welfare Cancel collection success")
			}
		}
		return serializer.BuildFailResponse("Welfare Cancel collection failed")
	case "1": //对商品进行取消收藏
		if err := models.WorkModel.CancelWorkCollection(objectId, userId); err == nil {
			if err = dao.CancelWorkCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Work Cancel Collection success")
			}
		}
		return serializer.BuildFailResponse("Work Cancel Collection failed")
	case "2": //对任务进行取消收藏
		if err := dao.CancelTaskCollection(objectId, userId); err == nil {
			if err = dao.CancelTaskCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Task Cancel Collection success")
			}
		}
		return serializer.BuildFailResponse("Task Cancel Collection failed")
	case "3": //对活动进行取消收藏
		if err := dao.CancelEventCollection(objectId, userId); err == nil {
			if err = dao.CancelEventCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Event Cancel Collection success")
			}
		}
		return serializer.BuildFailResponse("Event Cancel Collection failed")
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}

func GetCollectionTaskList(userId int64) serializer.CollectionTaskListResponse {
	tasks := dao.GetCollectionTaskList(userId)
	return serializer.BuildCollectionTaskListResponse(tasks)
}

func GetCollectionWorkList(userId int64) serializer.CollectionWorkListResponse {
	works := models.WorkModel.GetCollectionWorkList(userId)
	var workViewList []uint64
	for _, work := range works {
		workViewList = append(workViewList, dao.WorkView(work.ID))
	}
	return serializer.BuildCollectionWorkListResponse(works, workViewList)
}

// 小程序收藏接口
func AppletCollection(userId int64, objectId int64, objectTypeId string) serializer.Response {
	fmt.Println(objectTypeId)
	switch objectTypeId {
	case "0": //对作品进行收藏
		if err := models.WorkModel.AddWorkCollection(objectId, userId); err == nil {
			if err = dao.AddWorkCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Work Collection success")
			} else {
				return serializer.BuildErrorResponse(e.ErrorDatabase)
			}
		} else {
			return serializer.BuildErrorResponse(e.ErrorDatabase)
		}
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}
func AppletCancelCollection(userId int64, objectId int64, objectTypeId string) serializer.Response {
	switch objectTypeId {
	case "0": //对作品进行取消收藏
		if err := models.WorkModel.CancelWorkCollection(objectId, userId); err == nil {
			if err = dao.CancelWorkCollectionCount(objectId); err == nil {
				return serializer.BuildSuccessResponse("Work Cancel Collection success")
			}
		} else {
			return serializer.BuildErrorResponse(e.ErrorDatabase)
		}
		return serializer.BuildFailResponse("Work Cancel Collection failed")
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}
