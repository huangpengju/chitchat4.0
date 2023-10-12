package service

import (
	"errors"
	"fmt"
	"strconv"

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

// Create 创建 user 的服务
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
	return u.getUserByID(id)
}

// Validate 验证用户数据
func (u *userService) Validate(user *model.User) error {
	if user == nil {
		return errors.New("user 是空的")
	}
	if user.Name == "" {
		return errors.New("user 中 name 是空的")
	}
	if len(user.Password) < MinPasswordLength {
		return fmt.Errorf("密码长度不能小于%d", MinPasswordLength)
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

func (u *userService) getUserByID(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return u.userRepository.GetUserByID(uint(uid))
}
