package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

type hotSearchRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

func newHotSearchRepository(db *gorm.DB, rdb *database.RedisDB) HotSearchRepository {
	return &hotSearchRepository{
		db:  db,
		rdb: rdb,
	}
}
func (h *hotSearchRepository) List() ([]model.HotSearch, error) {
	hotSearchs := make([]model.HotSearch, 0)

	return hotSearchs, nil
}

func (h *hotSearchRepository) Create(user *model.Tag, hotSearch *model.HotSearch) (*model.HotSearch, error) {
	return hotSearch, nil
}

func (h *hotSearchRepository) Migrate() error {
	return h.db.AutoMigrate(&model.HotSearch{}, &model.Tag{})
}
