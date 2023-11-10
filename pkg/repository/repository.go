package repository

import (
	"context"

	"chitchat4.0/pkg/database"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, rdb *database.RedisDB) Repository {
	r := &repository{
		db:        db,
		rdb:       rdb,
		user:      newUserRepository(db, rdb), // user 数据仓库
		group:     newGroupRepository(db, rdb),
		rbac:      newRBACRepository(db, rdb),
		tag:       newTagRepository(db, rdb),
		hotSearch: newHotSearchRepository(db, rdb),
	}
	r.migrates = getMigrants(
		r.user,
		r.tag,
		r.hotSearch,
		r.group,
		r.rbac,
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
	group     GroupRepository
	rbac      RBACRepository
	tag       TagRepository
	hotSearch HotSearchRepository

	db  *gorm.DB
	rdb *database.RedisDB

	migrates []Migrant // 用于各模型迁移
}

func (r *repository) User() UserRepository {
	return r.user
}
func (r *repository) Group() GroupRepository {
	return r.group
}
func (r *repository) RBAC() RBACRepository {
	return r.rbac
}

func (r *repository) Tag() TagRepository {
	return r.tag
}

func (r *repository) HotSearch() HotSearchRepository {
	return r.hotSearch
}

// Ping 是使用 *repository 接收器定义的方法，
// 作用：实现了 Repository 仓库接口的 Ping 方法
// 查看数据库的连接状态
func (r *repository) Ping(ctx context.Context) error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	// PingContext 验证到 Postgres 数据库的连接是否仍然存在，并在必要时建立连接
	if err = db.PingContext(ctx); err != nil {
		return err
	}

	if r.rdb == nil {
		return nil
	}
	// 查看 redis 的连接状态
	if _, err := r.rdb.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

func (r *repository) Migrate() error {
	for _, m := range r.migrates {
		if err := m.Migrate(); err != nil {
			return err
		}
	}
	return nil
}
