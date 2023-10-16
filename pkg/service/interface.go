package service

import "chitchat4.0/pkg/model"

type UserService interface {
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Get(string) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
	// Delete(string) error
	Validate(*model.User) error
	Auth(*model.AuthUser) (*model.User, error)
	Default(*model.User)
}

type TagService interface {
	List() ([]model.Tag, error)
	Create(*model.User, *model.Tag) (*model.Tag, error)
	// Get(string) (*model.Tag, error)
	// Update(string, *model.Tag) (*model.Tag, error)
	// Delete(string) error
	// Validate(*model.Tag) error
}

type HotSearchService interface {
	List() ([]model.HotSearch, error)
	Create(*model.Tag, *model.HotSearch) (*model.HotSearch, error)
	// Get(string) (*model.HotSearch, error)
	// Update(string, *model.HotSearch) (*model.HotSearch, error)
	// Delete(string) error
	// Validate(*model.HotSearch) error
}
