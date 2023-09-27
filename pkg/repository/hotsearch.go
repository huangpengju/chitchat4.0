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
func (h *hotSearchRepository) List() ([]model.HostSearch, error) {
	hotSearchs := make([]model.HostSearch, 0)

	return hotSearchs, nil
}

func (h *hotSearchRepository) Create(user *model.Tag, hotSearch *model.HostSearch) (*model.HostSearch, error) {
	return hotSearch, nil
}

func (h *hotSearchRepository) Migrate() error {
	return h.db.AutoMigrate(&model.HostSearch{}, &model.Tag{})
}
