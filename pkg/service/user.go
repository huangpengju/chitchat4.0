package service

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) List() (model.Users, error) {
	users := make(model.Users, 0)

	return users, nil
}

func (u *userService) Create(user *model.User) (*model.User, error) {
	// if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
	// 	return nil, err
	// }

	// u.setCacheUser(user)

	return user, nil
}
