package repository

import (
	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"gorm.io/gorm"
)

// userRepository 是定义的用户仓库（属于大仓库中的小仓库），
// userRepository 中的 db|rdb 是结构体类型
// 作用：userRepository 实现了 UserRepository 接口
type userRepository struct {
	db  *gorm.DB
	rdb *database.RedisDB
}

// newUserRepository 接受两个参数，
// 参数1 是 *gorm.DB 结构体，参数2 是 *database.RedisDB 结构体。
// 作用：newUserRepository 内部实现了对 userRepository 结构体赋值，
// 返回结果是 userRepository 结构体地址，类型是 UserRepository 用户仓库接口
func newUserRepository(db *gorm.DB, rdb *database.RedisDB) UserRepository {
	return &userRepository{
		db:  db,
		rdb: rdb,
	}
}

// 下面是 userRepository 结构体实现 UserRepository 接口的全部方法
//
// List 是使用 *userRepository 接收器定义的方法，
// 作用：实现了 UserRepository 仓库接口的 User 方法
func (u *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)
	// if err := u.db.Preload(model.UserAuthInfoAssociation).Preload(model.GroupAssociation).Preload("Roles").Order("name").Find(&users).Error; err != nil {
	// 	return nil, err
	// }
	return users, nil
}

// Create 是使用 *userRepository 接收器定义的方法，
// 作用：实现了 UserRepository 仓库接口的 Create 方法
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	// if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
	// 	return nil, err
	// }

	// u.setCacheUser(user)

	return user, nil
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
