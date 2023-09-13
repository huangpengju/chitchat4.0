package models

import "github.com/jinzhu/gorm"

type HotList struct {
	gorm.Model
	Tag   Tag    `gorm:"foreignKey:TagID"`
	TagId uint   `gorm:"not null"`
	Title string `gorm:"type:varchar(200);not null"`
	Link  string `gorm:"type:varchar(300);not null"`
	Extra string `gorm:"type:varchar(50)"`
}
