package repository

import (
	"chitchat4.0/pkg/database"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, rdb *database.RedisDB) Repository {
	r := &repository{
		db:        db,
		rdb:       rdb,
		user:      newUserRepository(db, rdb),
		tag:       newTagRepository(db, rdb),
		hotSearch: newHotSearchRepository(db, rdb),
	}
	r.migrates = getMigrants(
		r.user,
		r.tag,
		r.hotSearch,
	)

	return r
}

func getMigrants(objs ...interface{}) []Migrant {
	var migrants []Migrant
	for _, obj := range objs {
		if m, ok := obj.(Migrant); ok {
			migrants = append(migrants, m)
		}
	}
	return migrants
}

type repository struct {
	user      UserRepository
	tag       TagRepository
	hotSearch HotSearchRepository

	db  *gorm.DB
	rdb *database.RedisDB

	migrates []Migrant // 用于各模型迁移
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) Tag() TagRepository {
	return r.tag
}

func (r *repository) HotSearch() HotSearchRepository {
	return r.hotSearch
}

func (r *repository) Migrate() error {
	for _, m := range r.migrates {
		if err := m.Migrate(); err != nil {
			return err
		}
	}
	return nil
}
