package service

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
)

type hotSearchService struct {
	hotSearchRepository repository.HotSearchRepository
}

func NewHotSearchService(hotSearchRepository repository.HotSearchRepository) HotSearchService {
	return &hotSearchService{
		hotSearchRepository: hotSearchRepository,
	}
}

func (h *hotSearchService) List() ([]model.HotSearch, error) {
	hotSearchs := make([]model.HotSearch, 0)
	return hotSearchs, nil
}

func (h *hotSearchService) Create(tag *model.Tag, hotSearch *model.HotSearch) (*model.HotSearch, error) {
	return hotSearch, nil
}
