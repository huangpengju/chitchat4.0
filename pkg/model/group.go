package model

type Group struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string `json:"name" gorm:"size:100;not null;unique"`
	Kind      string `json:"kind" gorm:"size:100"`
	Describe  string `json:"describe" gorm:"size:1024;"`
	CreatorId uint   `json:"creatorId"`
	UpdaterId uint   `json:"updaterId"`
	Users     []User `json:"users" gorm:"many2many:user_groups;"`
	Roles     []Role `json:"roles" gorm:"many2many:group_roles;"`

	BaseModel
}
