package service

import (
	"xueyigou_demo/global"
	"xueyigou_demo/models"
	"xueyigou_demo/proto"
	"xueyigou_demo/serializer"
)

var WorkService workService

type workService struct {
}

func (workService) GetPaged(req proto.WorkRequest) (*serializer.PagedResponse[serializer.WorkView], error) {
	total, err := models.WorkModel.GetTotal(req)
	if err != nil {
		return nil, err
	}
	data, err := models.WorkModel.GetList(req)
	if err != nil {
		return nil, err
	}
	return &serializer.PagedResponse[serializer.WorkView]{Data: serializer.BuildWorkViewList(data, req.UserId), Count: total}, nil
}

func (workService) GetPagedByUser(req proto.WorkRequest, currentUserId int64) (*serializer.PagedResponse[serializer.WorkView], error) {
	total, err := models.WorkModel.GetTotalByUser(req)
	if err != nil {
		return nil, err
	}
	data, err := models.WorkModel.GetPagedByUser(req)
	if err != nil {
		return nil, err
	}
	return &serializer.PagedResponse[serializer.WorkView]{Data: serializer.BuildWorkViewList(data, currentUserId), Count: total}, nil
}

func (workService) FindById(id int, userId int64) (*serializer.WorkView, error) {
	data, err := models.WorkModel.FindById(id)
	if err != nil {
		return nil, err
	}
	work := serializer.BuildWorkView(data, userId)
	return &work, err
}

func (workService) Add(w *models.Work) error {
	w.ID = global.Worker.GetId()
	for _, url := range w.PicturesUrlList {
		url.WorkId = w.ID
	}
	return w.Add()
}

func (workService) Delete(id int) error {
	return models.WorkModel.Delete(id)
}

func (workService) Update(w *models.Work) error {
	return w.Update()
}

func (workService) UpdateState(id int, state int) error {
	return models.WorkModel.UpdateState(id, state)
}
