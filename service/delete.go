package service

import (
	"xueyigou_demo/dao"
	"xueyigou_demo/serializer"
)

func DeleteObject(objectId int64, objectType int) serializer.Response {
	switch objectType {
	case 0: //任务
		if urls, err := dao.DeleteTask(objectId); err == nil {
			//TODO fs删除图片
			var urls_ []string
			for _, item := range urls {
				urls_ = append(urls_, item.Url)
			}
			if res, err1 := Deletefile(urls_); err1 != nil {
				return res
			}
			//
			return serializer.BuildSuccessResponse("Delete task success")
		} else {
			return serializer.BuildFailResponse("Delete task failed")
		}
		return serializer.BuildFailResponse("delete task")
	case 1: //商品
		//if urls, err := WorkService.Delete(objectId); err == nil {
		//	//TODO fs删除图片
		//	var urls_ []string
		//	for _, item := range urls {
		//		urls_ = append(urls_, item.Url)
		//	}
		//	if res, err1 := Deletefile(urls_); err1 != nil {
		//		return res
		//	}
		//	//
		//	return serializer.BuildSuccessResponse("Delete work success")
		//} else {
		//	return serializer.BuildFailResponse("Delete work failed")
		//}
		return serializer.BuildFailResponse("delete work")
	default:
		return serializer.BuildSuccessResponse("Object type is not defined")
	}
	return serializer.BuildFailResponse("delete")
}
