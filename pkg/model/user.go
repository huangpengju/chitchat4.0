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

type Users []User
