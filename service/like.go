package service

import (
	"fmt"
	"xueyigou_demo/dao"
	"xueyigou_demo/serializer"
)

// 网站点赞接口
func Like(userid int64, workid int64, obejcttype string) serializer.Response {
	switch obejcttype {
	case "0": //对一级评论进行点赞
		if err := dao.AddCommentL1Like(workid, userid); err == nil {
			fmt.Println(dao.Likes_get(workid, 0))
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	case "1": //二级评论
		if err := dao.AddCommentL2Like(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	case "2": //商品
		if err := dao.AddWorkLike(workid, userid); err == nil {
			dao.AddWorkLikeCount(workid)
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	case "3": //任务
		if err := dao.AddTaskLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Like success")
		} else {
			//panic(err)
			return serializer.BuildFailResponse("Like failed")
		}
	case "4": //公益
		if err := dao.AddWelfareLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	case "5": //活动
		if err := dao.AddEventsLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	}

	return serializer.BuildSuccessResponse("Object type is not defined")
}

func Dislike(userid int64, workid int64, obejcttype string) serializer.Response {
	switch obejcttype {
	case "0": //对一级评论取消点赞
		if err := dao.DeleteCommentL1Like(workid, userid); err == nil {
			fmt.Println(dao.Likes_get(workid, 0))
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	case "1":
		if err := dao.DeleteCommentL2Like(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("DisLike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	case "2":
		if err := dao.DeleteWorkLike(workid, userid); err == nil {
			dao.DeleteWorkLikeCount(workid)
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	case "3":
		if err := dao.DeleteTaskLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	case "4":
		//fmt.Println(dao.Likes_get(workid, 4))
		if err := dao.DeleteWelfareLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	case "5":
		//fmt.Println(dao.Likes_get(workid, 4))
		if err := dao.DeleteEventsLike(workid, userid); err == nil {
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}
func GetLikeList(userId int64, objectTypeId int64) serializer.LikeListResponse {
	var likeList []int64
	switch objectTypeId {
	case 0: //返回一级评论对象列表
		likeList = dao.GetCommentL1Like(userId)
	case 1:
		likeList = dao.GetCommentL2Like(userId)
	case 2:
		likeList = dao.GetWorkLike(userId)
	case 3:
		likeList = dao.GetTaskLike(userId)
	case 4:
		likeList = dao.GetWelfareLike(userId)
	case 5:
		likeList = dao.GetEventLike(userId)
	}
	return serializer.BuildLikeListResponse(likeList)
}

// 小程序点赞接口
func AppletLike(userid int64, workid int64, obejcttype string) serializer.Response {
	switch obejcttype {
	case "0": //作品
		if err := dao.AppletAddWorkLike(workid, userid); err == nil {
			dao.AddWorkLikeCount(workid)
			return serializer.BuildSuccessResponse("Like success")
		}
		return serializer.BuildFailResponse("Like failed")
	}

	return serializer.BuildSuccessResponse("Object type is not defined")
}

func AppletDislike(userid int64, workid int64, obejcttype string) serializer.Response {
	switch obejcttype {
	case "0": //作品
		if err := dao.AppletCancelWorkLike(workid, userid); err == nil {
			dao.DeleteWorkLikeCount(workid)
			return serializer.BuildSuccessResponse("Dislike success")
		}
		return serializer.BuildFailResponse("Dislike failed")
	}
	return serializer.BuildSuccessResponse("Object type is not defined")
}
