package model

import (
	"encoding/json"
)

// User 用户结构
type User struct {
	ID       uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null;unique"`
	Password string `json:"-" gorm:"size:256"`
	Email    string `json:"email" gorm:"size:256"`
	Avatar   string `json:"avatar" gorm:"size:256"`

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

// MarshalBinary 用于把 user 结构体实现 MarshalBinary 方法，MarshalBinary
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

// Users 变量是用户切片类型
type Users []User

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
