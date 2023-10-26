package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

type rbacRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

func newRBACRepository(db *gorm.DB, rdb *database.RedisDB) RBACRepository {
	return &rbacRepository{
		db:  db,
		rdb: rdb,
	}
}

func (rbac *rbacRepository) Migrate() error {
	return rbac.db.AutoMigrate(&model.Role{}, &model.Resource{})
}
