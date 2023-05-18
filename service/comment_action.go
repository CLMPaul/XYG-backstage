package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"xueyigou_demo/dao"
	"xueyigou_demo/global"
	"xueyigou_demo/middleware"
	"xueyigou_demo/models"
	"xueyigou_demo/serializer"

	"gorm.io/gorm"
)

type CommentService struct {
	Claim               *middleware.UserClaims
	Actiontype          string
	Content             string
	WorkId              int64
	Commentid           int64
	NewcommentL1        *models.FirstLevelComment
	NewcommentL2        *models.SecondLevelComment
	Userid              int64
	ObjectType          models.ObjectType
	CommentLevel        bool //true for first level, false for second level
	ResponseToCommentId string
}

type CommentServiceRet int

const (
	Addsucceed CommentServiceRet = iota
	Deletesucceed
	Failed
	UserNotExist
)

func (service *CommentService) CommentAction() interface{} {

	user, err := dao.GetUserById(service.Claim.Id)
	if err == nil {
		if service.Actiontype == "0" {
			t := time.Now()
			if service.Commentid == 0 {
				//一级评论
				service.NewcommentL1 = &models.FirstLevelComment{
					GoodId:     service.WorkId,
					Commenter:  user,
					PostDate:   strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()),
					Content:    service.Content,
					ObjectType: service.ObjectType,
					ID:         global.Worker.GetId(),
				}
				if response := service.AddCommentL1(); response != nil {
					return response
				}

			} else {
				//二级评论
				replyto_id := service.Userid
				fmt.Println(replyto_id)
				reply_to, err := dao.GetUserById(replyto_id)
				fmt.Println(err)
				if errors.Is(err, gorm.ErrRecordNotFound) {
					reply_to = dao.FindL1Commenter(service.Commentid)
					replyto_id = reply_to.UserId
				}
				fmt.Println(reply_to)
				service.NewcommentL2 = &models.SecondLevelComment{
					FirstLevelCommentId: service.Commentid,
					GoodId:              service.WorkId,
					Replier:             user,
					ReplierId:           user.UserId,
					ReplyDate:           strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day()),
					Content:             service.Content,
					ReplyTo:             reply_to,
					ReplytoId:           replyto_id,
					ObjectType:          service.ObjectType,
					ReplyToCommentId:    service.ResponseToCommentId,
					ID:                  global.Worker.GetId(),
				}
				if response := service.AddCommentL2(); response != nil {
					return response
				}
			}
			return serializer.BuildFailResponse("add fail")
		} else if service.Actiontype == "1" {
			if err := service.DeleteComment(); err != nil {
				return serializer.BuildFailResponse("delete fail")
			}
			return serializer.BuildFailResponse("delete success")
		}
		return serializer.BuildFailResponse("unkonw actiontype")
	}
	return serializer.BuildFailResponse("UserNotExist")
}
func (service *CommentService) AddCommentL1() *serializer.CommentListResponse {
	err := dao.AddCommentL1(service.NewcommentL1)
	if err != nil {
		return nil
	}
	flc := dao.FindL1CommentById(service.NewcommentL1.ID)
	slcs := dao.GetSecondLevelCommand(service.NewcommentL1.ID, service.ObjectType)
	comment_list := serializer.BuildCommentList(&flc, slcs)
	comment_lists := []serializer.CommentList{comment_list}
	return &serializer.CommentListResponse{
		Response:    serializer.BuildSuccessResponse("get comment list success"),
		CommentList: comment_lists,
		ThisID:      "f_" + strconv.Itoa(int(flc.ID)),
	}
}
func (service *CommentService) AddCommentL2() *serializer.CommentListResponse {
	err := dao.AddCommentL2(service.NewcommentL2)
	if err != nil {
		return nil
	}
	flc := dao.FindL1CommentById(service.NewcommentL2.FirstLevelCommentId)
	slcs := dao.GetSecondLevelCommand(service.NewcommentL2.FirstLevelCommentId, service.ObjectType)
	comment_list := serializer.BuildCommentList(&flc, slcs)
	comment_lists := []serializer.CommentList{comment_list}
	return &serializer.CommentListResponse{
		Response:    serializer.BuildSuccessResponse("get comment list success"),
		CommentList: comment_lists,
		ThisID:      "s_" + strconv.Itoa(int(service.NewcommentL2.ID)),
	}
}

func (service *CommentService) DeleteComment() error {
	err := dao.DeleteCommentById(service.Commentid, service.ObjectType, service.CommentLevel)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) GetCommentList() serializer.CommentListResponse {
	commentList := make([]serializer.CommentList, 0)
	FLCs := dao.GetFirstLevelCommand(service.WorkId, service.ObjectType)
	for _, FLC := range FLCs {
		SLCs := dao.GetSecondLevelCommand(int64(FLC.ID), service.ObjectType)
		comment := serializer.BuildCommentList(&FLC, SLCs)
		commentList = append(commentList, comment)
	}
	return serializer.CommentListResponse{
		Response:    serializer.BuildSuccessResponse("get comment list success"),
		CommentList: commentList,
	}
}
