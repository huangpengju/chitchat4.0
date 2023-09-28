package service

import "chitchat4.0/pkg/model"

type UserService interface {
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
}

type TagService interface {
	List() ([]model.Tag, error)
	Create(*model.User, *model.Tag) (*model.Tag, error)
}

type HotSearchService interface {
	List() ([]model.HotSearch, error)
	Create(*model.Tag, *model.HotSearch) (*model.HotSearch, error)
}
