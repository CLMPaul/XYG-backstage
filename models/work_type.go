package models

import (
	"gorm.io/gorm"
	"xueyigou_demo/db"
	"xueyigou_demo/proto"
	"xueyigou_demo/tools"
)

var WorkTypeModel WorkType

// WorkType 作品类别
type WorkType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

func (s *WorkType) Create() error {
	return db.DB.Create(s).Error
}

func (*WorkType) Delete(id int) error {
	return db.DB.Delete(&WorkTypeModel, "id=?", id).Error
}

func (s *WorkType) Update() error {
	return db.DB.Select([]string{"name", "sort"}).Updates(s).Error
}

func (*WorkType) GetList(req proto.WorkTypeRequest) ([]WorkType, error) {
	var data []WorkType
	err := db.DB.Scopes(WorkTypeModel.Query(req), tools.Paging(req.CurrentPage, req.PageSize)).Order("sort").Find(&data).Error
	return data, err
}

func (*WorkType) GetTotal(req proto.WorkTypeRequest) (int64, error) {
	var total int64
	err := db.DB.Scopes(WorkTypeModel.Query(req)).Count(&total).Error
	return total, err
}

func (*WorkType) GetAll() ([]WorkType, error) {
	var data []WorkType
	err := db.DB.Find(&data).Order("sort").Error
	return data, err
}

func (*WorkType) Query(req proto.WorkTypeRequest) func(*gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		if req.SearchText != "" {
			query = query.Where("name like ?", db.QuoteForContain(req.SearchText))
		}
		return query
	}
}
