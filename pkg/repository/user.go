package repository

import (
	"fmt"
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
// List 是使用 *userRepository 接收器定义的方法，
// 作用：实现了 UserRepository 仓库接口的 User 方法，用于获取用户列表
func (u *userRepository) List() (model.Users, error) {
	// 创建一个空的users
	users := make(model.Users, 0)
	if err := u.db.Order("id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create 实现 user 插入数据到数据库，
// 是使用 *userRepository 接收器定义的方法，
// 作用：实现了 user 仓库接口的 Create 方法，用于创建用户
func (u *userRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Select(userCreateField).Create(user).Error; err != nil {
		return nil, err
	}

	// 缓存用户信息
	// u.setCacheUser(user)
	if err := u.setCacheUser(user); err != nil {
		logrus.Errorf("redis 无法设置用户缓存：%v", err)
	}
	return user, nil
}

// Update 通过id修改用户，实现了修改用户的服务
func (u *userRepository) Update(user *model.User) (*model.User, error) {
	if err := u.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	u.rdb.HDel(user.CacheKey(), strconv.Itoa(int(user.ID)))
	return user, nil
}

func (u *userRepository) Delete(user *model.User) error {
	// 删掉授权信息
	err := u.db.Select(model.UserAuthInfoAssociation).Delete(user).Error
	if err != nil {
		return err
	}
	// 删掉Redis缓存
	u.rdb.HDel(user.CacheKey(), strconv.Itoa(int(user.ID)))
	return nil
}

// GetUserByID 通过ID获取用户，实现获取用户的服务
func (u *userRepository) GetUserByID(id uint) (*model.User, error) {
	// 创建一个空的user
	user := new(model.User)
	// Qmit 查询时省略password
	if err := u.db.Omit("Password").Preload(model.UserAuthInfoAssociation).Preload("Groups").Preload("Groups.Roles").Preload("Roles").First(user, id).Error; err != nil {
		return nil, err
	}
	// 设置用户的redis缓存
	if err := u.setCacheUser(user); err != nil {
		logrus.Errorf("获取单个用户后无法设置用户缓存：%v", err)
	}
	return user, nil
}

// 第三方登录查询user
func (u *userRepository) GetUserByAuthID(authType, authID string) (*model.User, error) {
	authInfo := new(model.AuthInfo)
	if err := u.db.Where("auth_type = ? and auth_id = ?", authType, authID).First(authInfo).Error; err != nil {
		return nil, err
	}

	return u.GetUserByID(authInfo.UserId)
}

// GetUserByName 通过name获取用户，实现获取用户的服务
func (u *userRepository) GetUserByName(name string) (*model.User, error) {
	user := new(model.User)
	// 关联 auth_infos | groups | group_roles | roles
	if err := u.db.Preload(model.UserAuthInfoAssociation).Preload("Groups").Preload("Groups.Roles").Preload("Roles").Where("name=?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// setCacheUser 缓存 user
func (u *userRepository) setCacheUser(user *model.User) error {
	if user == nil {
		return fmt.Errorf("%v", "注册成功，设置缓存是user为nil")
	}
	// 参数1：表名:id
	// 参数2：user的id
	// 参数3：user
	return u.rdb.HSet(user.CacheKey(), strconv.Itoa(int(user.ID)), user)
}

func (u *userRepository) Migrate() error {
	return u.db.AutoMigrate(&model.User{}, &model.AuthInfo{})
}
