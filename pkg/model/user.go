package model

import (
	"encoding/json"
	"time"
)

const (
	UserAssociation         = "Users"     // user 关联
	UserAuthInfoAssociation = "AuthInfos" // user 授权信息关联
	GroupAssociation        = "Groups"    // 组关联
)

// User 用户结构
type User struct {
	ID       uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null;unique"`
	Password string `json:"-" gorm:"size:256"`
	Email    string `json:"email" gorm:"size:256"`
	Avatar   string `json:"avatar" gorm:"size:256"` // 头像

	AuthInfos []AuthInfo `json:"authInfos" gorm:"foreignKey:UserId;references:ID"`
	Groups    []Group    `json:"groups" gorm:"many2many:user_groups;"`
	Roles     []Role     `json:"roles" gorm:"many2many:user_roles;"`

	BaseModel
}

// TableName 设置表名 uses
func (*User) TableName() string {
	return "users"
}

// CacheKey 返回表名和id，格式： users:id
func (u *User) CacheKey() string {
	return u.TableName() + ":id"
}

// 设置 Redis时使用
// MarshalBinary 用于把 user 结构体实现 MarshalBinary 方法，MarshalBinary
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// 获取 Redis时使用
// UnmarshalBinary
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type AuthInfo struct {
	ID           uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	UserId       uint      `json:"userId" gorm:"size:256"`
	Url          string    `json:"url" gorm:"size:256"`
	AuthType     string    `json:"authType" gorm:"size:256"`
	AuthId       string    `json:"authId" gorm:"size:256"`
	AccessToken  string    `json:"-" gorm:"size:256"`
	RefreshToken string    `json:"-" gorm:"size:256"`
	Expiry       time.Time `json:"-"`

	BaseModel
}

func (*AuthInfo) TableName() string {
	return "auth_infos"
}

// CreatedUser 结构模型用于绑定前端传入的参数
type CreatedUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

// GetUser 使用 CreateUser 中的数据给 User 用户结构模型进行赋值，
// 返回 User 用户结构模型
func (u *CreatedUser) GetUser() *User {
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}
}

// UpdatedUser 结构用于绑定前端传入的参数
type UpdatedUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetUser 返回一个 User，使用UpdatedUser中的数据
func (u *UpdatedUser) GetUser() *User {
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	}
}

// AuthUser 授权User(登录)
type AuthUser struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	SetCookie bool   `json:"setCookie"`
	AuthType  string `json:"authType"`
	AuthCode  string `json:"authCode"`
}

// Users 变量是用户切片类型
type Users []User
