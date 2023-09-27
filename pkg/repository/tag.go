package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

type tagRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

func newTagRepository(db *gorm.DB, rdb *database.RedisDB) TagRepository {
	return &tagRepository{
		db:  db,
		rdb: rdb,
	}
}

func (t *tagRepository) List() ([]model.Tag, error) {
	tags := make([]model.Tag, 0)
	return tags, nil
}

func (t *tagRepository) Create(user *model.User, tag *model.Tag) (*model.Tag, error) {
	return tag, nil
}

func (t *tagRepository) Migrate() error {
	return t.db.AutoMigrate(&model.Tag{}, &model.User{})
}
