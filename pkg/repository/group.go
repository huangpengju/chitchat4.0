package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

type groupRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

func newGroupRepository(db *gorm.DB, rdb *database.RedisDB) GroupRepository {
	return &groupRepository{
		db:  db,
		rdb: rdb,
	}
}

func (g *groupRepository) Migrate() error {
	return g.db.AutoMigrate(&model.Group{})
}
