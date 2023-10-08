package service

import (
	"errors"
	"fmt"

	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 6 // 密码的长度
)

// userService 用户服务结构模型
type userService struct {
	userRepository repository.UserRepository
}

// NewUserService 创建一个用户服务
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// List 用户列表服务
func (u *userService) List() (model.Users, error) {
	users := make(model.Users, 0)

	return users, nil
}

// Create 创建用户服务
func (u *userService) Create(user *model.User) (*model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return u.userRepository.Create(user)
}

// Get 获取用户服务
func (u *userService) Get(id string) (*model.User, error) {
	user := &model.User{}
	return user, nil
}

// Validate 验证用户数据
func (u *userService) Validate(user *model.User) error {
	if user == nil {
		return errors.New("user is empty")
	}
	if user.Name == "" {
		return errors.New("user name is empty")
	}
	if len(user.Password) < MinPasswordLength {
		return fmt.Errorf("password length must great than %d", MinPasswordLength)
	}
	return nil
}

// Default 给用户模型中的Email设置默认值
func (u *userService) Default(user *model.User) {
	if user == nil || user.Name == "" {
		return
	}
	if user.Email == "" {
		user.Email = fmt.Sprintf("%s@qinng.io", user.Name)
	}
}
