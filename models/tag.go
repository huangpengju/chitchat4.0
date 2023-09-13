package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null;unique"`
	Sort      int    `gorm:"type:integer"`
	SourceKey string `gorm:"type:varchar(50);not null;unique"`
	IconColor string `gorm:"type:varchar(50)"`
}
