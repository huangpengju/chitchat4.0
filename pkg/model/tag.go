package model

// Tag 标签结构
type Tag struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:256;not null;unique"`
	Sort      int    `json:"sort" `
	SourceKey string `json:"source_key" gorm:"size:100" `
	IconColor string `json:"icon_color" gorm:"size:100"`

	Creator   User `json:"creator" gorm:"foreignKey:CreatorID"`
	CreatorID uint `json:"creatorId"`

	BaseModel
}
