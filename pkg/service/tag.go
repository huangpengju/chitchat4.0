package service

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
)

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepository,
	}
}

func (t *tagService) List() ([]model.Tag, error) {
	tags := make([]model.Tag, 0)

	return tags, nil
}

func (t *tagService) Create(user *model.User, tag *model.Tag) (*model.Tag, error) {
	return tag, nil
}
