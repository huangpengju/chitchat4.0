package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null;unique;comment:'热搜标签id'"`
	Sort      int    `gorm:"type:tinyint(3);comment:'热搜标签排序'"`
	SourceKey string `gorm:"type:varchar(50);not null;unique;comment:'热搜标签类型'"`
	IconColor string `gorm:"type:varchar(50);comment:'热搜标签图标颜色'"`
}

/*
	Name string `gorm:"type:varchar(50);not null;unique;comment:'热搜标签id'"`
	Sort      int    `gorm:"type:tinyint(3);comment:'热搜标签排序'"`
	SourceKey string `gorm:"type:varchar(50);not null;unique;comment:'热搜标签类型'"`
	IconColor string `gorm:"type:varchar(50);comment:'热搜标签图标颜色'"`
*/
