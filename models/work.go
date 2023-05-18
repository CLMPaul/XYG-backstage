package models

import (
	"errors"
	"gorm.io/gorm"
	"xueyigou_demo/db"
	"xueyigou_demo/internal"
	"xueyigou_demo/proto"
	"xueyigou_demo/tools"
)

var WorkModel Work

type Work struct {
	internal.ModelLogically
	Title           string            `json:"title"`
	CoverPicture    string            `json:"coverPicture"` // 封面
	Introduce       string            `json:"introduce"`    // 简介
	TypeID          int64             `json:"typeID"`
	Status          int               `json:"status"`  // 商品状态  0 待审核;1 审核通过;2 审核不通过;3 作品下架;9 作品强制下架
	Like            int64             `json:"like"`    // 商品点赞量
	Collect         int64             `json:"collect"` // 商品被收藏数
	View            int64             `json:"view"`    // 商品被浏览量
	PostUserId      int64             `json:"postUserId"`
	PicturesUrlList []WorkPicturesUrl `json:"picturesUrlList" gorm:"constraint:OnDelete:CASCADE;"` // 图像url

	WorkType    WorkType `json:"workType" gorm:"foreignKey:TypeID;constraint:OnDelete:CASCADE;"`
	LikeUser    []*User  `json:"likeUser,omitempty" gorm:"many2many:work_like"`
	CollectUser []*User  `json:"collectUser,omitempty" gorm:"many2many:work_collection"`
	PostUser    User     `gorm:"foreignKey:PostUserId;references:UserId"` // 发布人
}

type WorkPicturesUrl struct {
	ID     int64  `json:"id"`
	WorkId int64  `json:"-"`
	Url    string `json:"url"`
}

func (*Work) FindByID(id int64) (*Work, error) {
	var data Work
	err := db.DB.First(&data, "id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &data, err
}

func (*Work) GetCollectionWorkList(userId int64) []Work {
	var works []Work
	var user User
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if err = db.DB.Model(&user).Association("CollectionWork").Find(&works); err != nil {
			panic(err)
		}
	}
	for index, work := range works {
		var w Work
		if err := db.DB.Where("id = ?", work.ID).Preload("WorkSubject").Find(&w).Error; err != nil {
			panic(err)
		}
		works[index].WorkType = w.WorkType
		works[index].PostUser = user
	}
	return works
}

func (*Work) CancelWorkCollection(workId int64, userId int64) error {
	var user User
	work, _ := WorkModel.FindByID(workId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionWork").Delete(&work)
		return err
	}
	return nil
}

func (*Work) AddWorkCollection(workId int64, userId int64) error {
	var user User
	work, _ := WorkModel.FindByID(workId)
	if err := db.DB.Where("user_id = ?", userId).Find(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.DB.Model(&user).Association("CollectionWork").Append(&work)
		return err
	}
	return nil
}

//func (work *Work) View() uint64 {
//	countStr, _ := cache.GetClient().Get(cache.WelfareViewKey(work.ID)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}

// AddView 商品游览
//func (work *Work) AddView() {
//	// 增加商品浏览数
//	cache.GetClient().Incr(cache.WorkViewKey(work.ID))
//
//}

func (*Work) GetList(req proto.WorkRequest) (list []Work, err error) {
	err = db.DB.Select("collect", "introduce", "id",
		"like", "title", "cover_picture", "view", "post_user_id", "type_id", "status").
		Preload("WorkType").Preload("PostUser").Preload("PicturesUrlList").
		Scopes(WorkModel.workPagingQuery(req), tools.Paging(req.CurrentPage, req.PageSize), WorkModel.workPagingOrder(req.Mode)).Find(&list).Error
	return
}

func (*Work) GetTotal(req proto.WorkRequest) (total int64, err error) {
	err = db.DB.Model(WorkModel).Scopes(WorkModel.workPagingQuery(req)).Count(&total).Error
	return
}

func (*Work) FindById(id int) (*Work, error) {
	var w Work
	err := db.DB.Preload("WorkType").Preload("PostUser").Preload("PicturesUrlList").First(&w, "id=?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &w, err
}

func (*Work) GetPagedByUser(req proto.WorkRequest) (list []Work, err error) {
	err = db.DB.Select("collect", "id", "introduce",
		"like", "title", "cover_picture", "view", "type_id", "post_user_id", "status").
		Preload("WorkType").Preload("PostUser").Preload("PicturesUrlList").
		Scopes(WorkModel.workByUserPagingQuery(req), tools.Paging(req.CurrentPage, req.PageSize)).Find(&list).Error
	return list, err
}

func (*Work) GetTotalByUser(req proto.WorkRequest) (total int64, err error) {
	err = db.DB.Model(WorkModel).Scopes(WorkModel.workByUserPagingQuery(req)).Count(&total).Error
	return
}

func (w *Work) Add() error {
	return db.DB.Create(w).Error
}

func (w *Work) Update() error {
	return db.DB.Select([]string{"introduce", "title", "cover_picture", "type_id"}).
		Updates(w).Error
}

func (*Work) UpdateState(id int, state int) error {
	return db.DB.Model(WorkModel).Where("id=?", id).Update("status", state).Error
}

func (*Work) Delete(id int) error {
	return db.DB.Delete(&WorkModel, "id=?", id).Error
}

func (*Work) workPagingQuery(req proto.WorkRequest) func(*gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		if req.Keywords != "" {
			query = query.Where("introduce like ? or title like ?", db.QuoteForContain(req.Keywords), db.QuoteForContain(req.Keywords))
		}
		if req.TypeID != nil {
			query = query.Where("type_id=?", req.TypeID)
		}
		if req.State != nil {
			query = query.Where("status=?", req.State)
		}
		return query
	}
}

func (*Work) workByUserPagingQuery(req proto.WorkRequest) func(*gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		query = query.Where("post_user_id = ?", req.UserId)

		if req.Keywords != "" {
			query = query.Where("introduce like ? or title like ?", db.QuoteForContain(req.Keywords), db.QuoteForContain(req.Keywords))
		}
		if req.TypeID != nil {
			query = query.Where("type_id=?", req.TypeID)
		}
		if req.State != nil {
			query = query.Where("status=?", req.State)
		}
		return query
	}
}

func (*Work) workPagingOrder(mode int) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		switch mode {
		case 0:
			query = query.Order("created_at desc")
		case 1:
			query = query.Order("`like` desc")
		}
		return query
	}
}
