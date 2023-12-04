package repository

import (
	"context"

	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *repository) Close() error {
	db, _ := r.db.DB()
	if db != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}

	if r.rdb != nil {
		if err := r.rdb.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) Init() error {
	resources := []model.Resource{
		{
			Name:  model.ContainerResource,
			Scope: model.ClusterScope,
		},
		{
			Name:  model.ContainerResource + "/log",
			Scope: model.ClusterScope,
		},
		{
			Name:  model.ContainerResource + "/exec",
			Scope: model.ClusterScope,
		},
		{
			Name:  model.ContainerResource + "/proxy",
			Scope: model.ClusterScope,
		},
		// {
		// 	Name:  model.PostResource,
		// 	Scope: model.ClusterScope,
		// },
		{
			Name:  model.GroupResource,
			Scope: model.ClusterScope,
		},
		{
			Name:  model.UserResource,
			Scope: model.ClusterScope,
		},
		{
			Name:  model.RoleResource,
			Scope: model.ClusterScope,
		},
		{
			Name:  model.AuthResource,
			Scope: model.ClusterScope,
		},
		{
			Name:  model.NamespaceResource,
			Scope: model.ClusterScope,
		},
		// {
		// 	Name:  model.KubeDeployment,
		// 	Scope: model.NamespaceScope,
		// },
		// {
		// 	Name:  model.KubeStatefulset,
		// 	Scope: model.NamespaceScope,
		// },
		// {
		// 	Name:  model.KubeDaemonset,
		// 	Scope: model.NamespaceScope,
		// },
		// {
		// 	Name:  model.KubePod,
		// 	Scope: model.NamespaceScope,
		// },
		// {
		// 	Name:  model.KubeService,
		// 	Scope: model.NamespaceScope,
		// },
		// {
		// 	Name:  model.KubeIngress,
		// 	Scope: model.NamespaceScope,
		// },
	}
	if err := r.RBAC().CreateResources(resources, clause.OnConflict{DoNothing: true}); err != nil {
		return err
	}

	// create default group
	groups := []model.Group{
		{
			Name:     model.RootGroup,
			Kind:     model.SystemGroup,
			Describe: "system root group",
		},
		{
			Name:     model.AuthenticatedGroup,
			Kind:     model.SystemGroup,
			Describe: "system group contains all authenticated user",
		},
		{
			Name:     model.UnAuthenticatedGroup,
			Kind:     model.SystemGroup,
			Describe: "system group contains all unauthenticated user",
		},
	}
	if err := r.Group().CreateGroups(groups, clause.OnConflict{DoNothing: true}); err != nil {
		return err
	}

	return nil
}
