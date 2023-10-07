package model

// User 用户结构
type User struct {
	ID       uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null;unique"`
	Password string `json:"-" gorm:"size:256"`
	Email    string `json:"email" gorm:"size:256"`
	Avatar   string `json:"avatar" gorm:"size:256"`

	BaseModel
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
