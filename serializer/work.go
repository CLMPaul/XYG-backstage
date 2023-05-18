package serializer

import (
	"strconv"
	"xueyigou_demo/internal"
	"xueyigou_demo/models"
)

type postUser struct {
	UserContribution int64  `json:"user_contribution"` // 用户公益贡献值
	UserID           string `json:"user_id"`           // 发布该商品的用户id
	UserIdentity     string `json:"user_identity"`     // 用户身份，用户身份，如：普通用户
	UserMedals       int64  `json:"user_medals"`       // 用户获得学易购奖章数
	UserName         string `json:"user_name"`         // 企业名/用户名
	UserHeadPhoto    string `json:"user_head_photo"`
}

type CollectionWork struct {
	WorkCollection int64  `json:"work_collection"` // 商品收藏数
	WorkIntroduce  string `json:"work_introduce"`  // 商品简介/信息
	WorkID         int64  `json:"work_id"`         // 商品id
	WorkLike       int64  `json:"work_like"`       // 商品点赞数
	WorkName       string `json:"work_name"`       // 商品名称
	WorkPicture    string `json:"work_picture"`    // 商品封面图片url
	//WorkPrice      int64      `json:"work_price"`      // 商品价格
	WorkMax     int64    `json:"work_max"`
	WorkMin     int64    `json:"work_min"`
	WorkSubject []string `json:"work_subject"` // 商品所属科目
	WorkView    uint64   `json:"work_view"`    // 商品浏览量
	UserId      int64    `json:"user_id"`
	UserName    string   `json:"user_name"`
	HeaderPhoto string   `json:"user_head_photo"`
}
type CollectionWorkListResponse struct {
	Response
	CollectionWorkList []CollectionWork `json:"collection_work_list"`
}

func BuildCollectionWorkListResponse(works []models.Work, workViewList []uint64) CollectionWorkListResponse {
	response := Response{
		ResultStatus: 0,
		ResultMsg:    "get collection_works_list success",
	}
	var collectionWorkList []CollectionWork
	for i, work := range works {
		collectionWork := CollectionWork{
			WorkCollection: work.Collect,
			WorkIntroduce:  work.Introduce,
			WorkID:         work.ID,
			WorkLike:       work.Like,
			WorkName:       work.Title,
			WorkPicture:    work.CoverPicture,
			//WorkMax:        work.WorkMax,
			//WorkMin:        work.WorkMin,
			WorkView:    workViewList[i],
			UserId:      work.PostUserId,
			UserName:    work.PostUser.Name,
			HeaderPhoto: work.PostUser.HeaderPhoto,
		}
		//for _, subject := range work.WorkSubject {
		//	collectionWork.WorkSubject = append(collectionWork.WorkSubject, subject.Item)
		//}
		collectionWorkList = append(collectionWorkList, collectionWork)
	}
	collectionWorkListResponse := CollectionWorkListResponse{
		Response:           response,
		CollectionWorkList: collectionWorkList,
	}
	return collectionWorkListResponse
}

type AddWorkPicturesUrlResponse struct {
	Response
	Url []string
}

func BuildAddWorkPicturesUrlResponse(urls []string) AddWorkPicturesUrlResponse {
	return AddWorkPicturesUrlResponse{
		Response: BuildSuccessResponse("work picture"),
		Url:      urls,
	}
}

type WorkView struct {
	internal.ViewModel
	Title        string `json:"title"`
	CoverPicture string `json:"coverPicture"` // 封面
	Introduce    string `json:"introduce"`    // 简介
	TypeID       int64  `json:"typeID"`
	Status       int    `json:"status"`  // 商品状态  0 待审核;1 审核通过;2 审核不通过;3 作品下架;9 作品强制下架
	Like         int64  `json:"like"`    // 商品点赞量
	Collect      int64  `json:"collect"` // 商品被收藏数
	View         int64  `json:"view"`    // 商品被浏览量
	PostUserId   int64  `json:"postUserId"`
	IsLike       bool   `json:"isLike"`
	IsCollect    bool   `json:"isCollect"`

	PicturesUrlList []models.WorkPicturesUrl `json:"picturesUrlList"` // 图像url
	WorkType        models.WorkType          `json:"workType"`
	LikeUser        []UserView               `json:"likeUser,omitempty"`
	CollectUser     []UserView               `json:"collectUser,omitempty"`
	PostUser        UserView                 `json:"postUser"` // 发布人
}

type UserView struct {
	UserId          string `json:"user_id" `
	Name            string `json:"user_name,omitempty"` //用户名
	Sex             string `json:"sex,omitempty"`
	HeaderPhoto     string `json:"header_photo,omitempty"`
	BackgroundPhoto string `json:"background_photo,omitempty"`
}

func BuildWorkViewList(data []models.Work, userId int64) []WorkView {
	var works []WorkView
	for _, d := range data {
		works = append(works, BuildWorkView(&d, userId))
	}
	return works
}

func BuildWorkView(w *models.Work, userId int64) WorkView {
	data := WorkView{
		ViewModel: internal.ViewModel{
			ID:        strconv.FormatInt(w.ID, 10),
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
			DeletedAt: w.DeletedAt,
		},
		Title:           w.Title,
		CoverPicture:    w.CoverPicture,
		Introduce:       w.Introduce,
		TypeID:          w.TypeID,
		Status:          w.Status,
		Like:            w.Like,
		Collect:         w.Collect,
		View:            w.View,
		PostUserId:      w.PostUserId,
		PicturesUrlList: w.PicturesUrlList,
		WorkType:        w.WorkType,
		PostUser: UserView{
			UserId:          strconv.FormatInt(w.PostUser.UserId, 10),
			Name:            w.PostUser.Name,
			Sex:             w.PostUser.Sex,
			HeaderPhoto:     w.PostUser.HeaderPhoto,
			BackgroundPhoto: w.PostUser.BackgroundPhoto,
		},
	}
	if len(w.LikeUser) > 0 {
		for _, user := range w.LikeUser {
			if user.UserId == userId {
				data.IsLike = true
			}
			data.LikeUser = append(data.LikeUser, UserView{
				UserId:          strconv.FormatInt(w.PostUser.UserId, 10),
				Name:            user.Name,
				Sex:             user.Sex,
				HeaderPhoto:     user.HeaderPhoto,
				BackgroundPhoto: user.BackgroundPhoto,
			})
		}
	}
	if len(w.CollectUser) > 0 {
		for _, user := range w.CollectUser {
			if user.UserId == userId {
				data.IsCollect = true
			}
			data.CollectUser = append(data.CollectUser, UserView{
				UserId:          strconv.FormatInt(w.PostUser.UserId, 10),
				Name:            user.Name,
				Sex:             user.Sex,
				HeaderPhoto:     user.HeaderPhoto,
				BackgroundPhoto: user.BackgroundPhoto,
			})
		}
	}
	return data
}
