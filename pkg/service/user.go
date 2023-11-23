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

// List 实现获取用户列表服务
func (u *userService) List() (model.Users, error) {
	// 调用user仓库，完成具体细节
	return u.userRepository.List()
}

// Get 用户服务（获取单个用户）
func (u *userService) Get(id string) (*model.User, error) {
	//
	return u.getUserByID(id)
}

// Create 实现 Create user 的服务
func (u *userService) Create(user *model.User) (*model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)
	return u.userRepository.Create(user)
}

// Update 实现 Update user 的服务
func (u *userService) Update(id string, new *model.User) (*model.User, error) {
	old, err := u.getUserByID(id)
	if err != nil {
		return nil, err
	}
	if new.ID != 0 && old.ID != new.ID {
		return nil, fmt.Errorf("update user id %s not match（不匹配）", id)
	}
	new.ID = old.ID

	if len(new.Password) > 0 {
		passwrd, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		new.Password = string(passwrd)
	}
	return u.userRepository.Update(new)
}

func (u *userService) Delete(id string) error {
	user, err := u.getUser(id)
	if err != nil {
		return err
	}
	return u.userRepository.Delete(user)
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

// Auth() 授权，通过接收到的参数实现登录验证的服务
func (u *userService) Auth(auser *model.AuthUser) (*model.User, error) {
	if auser == nil || auser.Name == "" || auser.Password == "" {
		return nil, fmt.Errorf("name or password is empty")
	}
	// 通过name查询user是否存在
	user, err := u.userRepository.GetUserByName(auser.Name)
	if err != nil {
		return nil, err
	}
	// 数据库用户密码和登录用户密码进行对比
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auser.Password)); err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

// getUserByID 通过ID获取用户的服务，接收用户id后调用user仓库
func (u *userService) getUserByID(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	// 在仓库中实现具体的用户查询操作
	return u.userRepository.GetUserByID(uint(uid))
}

func (u *userService) getUser(id string) (*model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return &model.User{ID: uint(uid)}, nil
}
