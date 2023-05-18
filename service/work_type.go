package service

import (
	"xueyigou_demo/global"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
	"xueyigou_demo/serializer"
)

var SubjectItemService subjectItemService

type subjectItemService struct {
}

func (subjectItemService) Create(s *models.WorkType) error {
	s.Id = global.Worker.GetId()
	return s.Create()
}
func (subjectItemService) Delete(id int) error {
	return models.WorkTypeModel.Delete(id)
}

func (subjectItemService) Update(s *models.WorkType) error {
	return s.Update()
}

func (subjectItemService) GetPaged(req proto.WorkTypeRequest) (*serializer.PagedResponse[models.WorkType], error) {
	total, err := models.WorkTypeModel.GetTotal(req)
	if err != nil {
		return nil, err
	}
	data, err := models.WorkTypeModel.GetList(req)
	if err != nil {
		return nil, err
	}
	return &serializer.PagedResponse[models.WorkType]{Data: data, Count: total}, nil
}

func (subjectItemService) GetAll() ([]models.WorkType, error) {
	return models.WorkTypeModel.GetAll()
}
