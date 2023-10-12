package repository

import (
	"strconv"

	"chitchat4.0/pkg/database"
	"chitchat4.0/pkg/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	userCreateField = []string{"name", "email", "password", "avatar"}
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

// Create 实现 user 插入数据到数据库，
// 是使用 *userRepository 接收器定义的方法，
// 作用：实现了 user 仓库接口的 Create 方法
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
		return nil, err
	}

	// 缓存用户
	// u.setCacheUser(user)
	if err := u.setCacheUser(user); err != nil {
		logrus.Errorf("redis 无法设置用户：%v", err)
	}

	return user, nil
}

func (u *userRepository) GetUserByID(id uint) (*model.User, error) {

	user := new(model.User)

	if err := u.db.Omit("Password").First(user, id).Error; err != nil {
		return nil, err
	}
	if err := u.setCacheUser(user); err != nil {
		logrus.Errorf("无法设置用户：%v", err)
	}
	return user, nil
}

// setCacheUser 缓存 user
func (u *userRepository) setCacheUser(user *model.User) error {
	if user == nil {
		return nil
	}
	// 参数1：表名:id
	// 参数2：user的id
	// 参数3：user
	return u.rdb.HSet(user.CacheKey(), strconv.Itoa(int(user.ID)), user)
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{})
}
