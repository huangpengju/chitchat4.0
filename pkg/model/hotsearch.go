package model

// HostList 热搜列表结构
type HotSearch struct {
	ID    uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Title string `json:"title" gorm:"size:512;not null"`
	Link  string `json:"link" gorm:"size:512;not null"`
	Extra string `json:"extra" gorm:"size:256"`

	Tag   Tag  `json:"tag" gorm:"foreignKey:TagID"`
	TagID uint `json:"tagId"`

	BaseModel
}
