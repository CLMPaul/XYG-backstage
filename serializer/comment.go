package serializer

import (
	"strconv"
	"xueyigou_demo/global"
	"xueyigou_demo/models"
)

type CommentListResponse struct {
	Response
	CommentList []CommentList `json:"comment_list,omitempty"`
	ThisID      string        `json:"thisID,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type CommentListUserInfo struct {
	Name   string `json:"user_name,omitempty"`
	UserId int64  `json:"user_id,omitempty"  gorm:"primaryKey"`
}

type Comment struct {
	CommentID   string          `json:"comment_id"`  // 评论id
	Commentator CommentUserInfo `json:"commentator"` // 评论者
	Content     string          `json:"content"`     // 评论内容
	CreateDate  string          `json:"create_date"` // 评论发布日期，mm-dd
}

type CommentUserInfo struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	HeaderPhoto         string `json:"header_photo"`
	ResponseToCommentId string `json:"comment_id,omitempty"`
}

type CommentList struct {
	CommentID   string          `json:"comment_id"`  // 评论id
	Commentator CommentUserInfo `json:"commentator"` // 评论者
	Content     string          `json:"content"`     // 评论内容
	CreateDate  string          `json:"create_date"` // 评论发布日期，mm-dd
	Respons     []response      `json:"response"`    // 评论回复
	Likes       int64           `json:"likes"`       //点赞
}

type response struct {
	CommentID    string          `json:"comment_id"`    // 评论id
	Content      string          `json:"content"`       // 评论内容
	Respondent   CommentUserInfo `json:"respondent"`    // 回复者
	ResponsTo    CommentUserInfo `json:"response_to"`   // 回复对象
	ResponseDate string          `json:"response_date"` // 回复日期，mm-dd
	Likes        int64           `json:"likes"`         //点赞
}

func BuildCommentList(flc *models.FirstLevelComment, slcs []models.SecondLevelComment) CommentList {
	respons := make([]response, 0)
	for _, slc := range slcs {
		respondent := CommentUserInfo{
			Name:        slc.Replier.Name,
			Id:          slc.ReplierId,
			HeaderPhoto: slc.Replier.HeaderPhoto,
		}
		ResponsTo := CommentUserInfo{
			Name:                slc.ReplyTo.Name,
			Id:                  slc.ReplytoId,
			HeaderPhoto:         slc.ReplyTo.HeaderPhoto,
			ResponseToCommentId: slc.ReplyToCommentId,
		}

		respon := response{
			CommentID:    "s_" + strconv.Itoa(int(slc.ID)),
			Content:      slc.Content,
			Respondent:   respondent,
			ResponsTo:    ResponsTo,
			ResponseDate: slc.CreatedAt.Format(global.Timelayout),
			Likes:        slc.CommentLikes,
		}
		respons = append(respons, respon)
	}

	commenter := CommentUserInfo{
		Name:        flc.Commenter.Name,
		Id:          flc.CommenterId,
		HeaderPhoto: flc.Commenter.HeaderPhoto,
	}
	comment_list := CommentList{
		CommentID:   "f_" + strconv.Itoa(int(flc.ID)),
		Content:     flc.Content,
		Commentator: commenter,
		CreateDate:  flc.CreatedAt.Format(global.Timelayout),
		Respons:     respons,
		Likes:       flc.CommentLikes,
	}
	return comment_list
}

func BuildComment(flc *models.FirstLevelComment, slc *models.SecondLevelComment) Comment {
	if flc != nil {
		Commentator := CommentUserInfo{
			Id:          flc.CommenterId,
			Name:        flc.Commenter.Name,
			HeaderPhoto: flc.Commenter.HeaderPhoto,
		}
		return Comment{
			CommentID:   "f_" + strconv.Itoa(int(flc.ID)),
			Content:     flc.Content,
			CreateDate:  flc.CreatedAt.Format(global.Timelayout),
			Commentator: Commentator,
		}
	}
	Commentator := CommentUserInfo{
		Id:          slc.ReplierId,
		Name:        slc.Replier.Name,
		HeaderPhoto: slc.Replier.HeaderPhoto,
	}
	return Comment{
		CommentID:   "s_" + strconv.Itoa(int(slc.ID)),
		Content:     slc.Content,
		Commentator: Commentator,
		CreateDate:  slc.CreatedAt.Format(global.Timelayout),
	}
}
