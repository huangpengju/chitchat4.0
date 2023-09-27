package model

// HostList 热搜列表结构
type HostSearch struct {
	ID    uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Tag   Tag    `json:"tag" gorm:"foreignKey:TagID"`
	TagId uint   `json:"tagId"`
	Title string `json:"title" gorm:"size:512;not null"`
	Link  string `json:"link" gorm:"size:512;not null"`
	Extra string `json:"extra" gorm:"size:256"`

	BaseModel
}
