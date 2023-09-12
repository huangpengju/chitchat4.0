package models

import "github.com/jinzhu/gorm"

type HotList struct {
	gorm.Model
	Tag   Tag
	TagId uint   `gorm:"not null;comment:'热搜标签id'"`
	Title string `gorm:"type:varchar(200);not null;comment:'热搜标题'"`
	Link  string `gorm:"type:varchar(300);not null;comment:'热搜地址'"`
	Extra string `gorm:"type:varchar(50);comment:'额外信息'"`
}
