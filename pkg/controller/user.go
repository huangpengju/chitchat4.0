package controller

import (
	"net/http"
	"strconv"

	"chitchat4.0/pkg/authorization"
	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"chitchat4.0/pkg/utils/trace"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserController 用户控制器，
// userService 字段表示 user 服务接口
type UserController struct {
	userService service.UserService
}

// NewUserController 创建 user 控制器，
// 用于实现用 user 服务接口
func NewUserController(userService service.UserService) Controller {
	return &UserController{
		userService: userService,
	}
}

// @Summary Create user | 创建用户
// @Description 创建用户并存储
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.CreatedUser true "user 信息"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users [post]
func (u *UserController) Create(c *gin.Context) {
	// 准备一个空的结构模型 createdUser
	createdUser := new(model.CreatedUser)
	// 把接收到的数据绑定到 createdUser
	if err := c.BindJSON(createdUser); err != nil {
		// 数据绑定失败时，做出响应
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	// 把 createdUser 的值赋值给 User 的结构模型
	user := createdUser.GetUser()
	// 验证 user 的用户名和密码
	if err := u.userService.Validate(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	// user 邮箱为空时，设置默认邮箱
	u.userService.Default(user)

	// 追踪步骤：开启创建 user
	common.TraceStep(c, "start create user", trace.Field{Key: "user", Value: user.Name})
	defer common.TraceStep(c, "create user done", trace.Field{Key: "user", Value: user.Name})
	user, err := u.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		// return  需要确认是否添加
	}
	common.ResponseSuccess(c, user)
}

// @Summary Get user | 获取单个用户
// @Description 获取用户并保存
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [get]
func (u *UserController) Get(c *gin.Context) {
	// 调用用户服务
	user, err := u.userService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, user)
}

// @Summary List user | 用户列表
// @Description 获取用户列表并存储
// @Produce json
// @Tags user
// @Security JWT
// @Success 200 {object} common.Response{data=model.Users}
// @Router /api/v1/users [get]
func (u *UserController) List(c *gin.Context) {
	common.TraceStep(c, "start list user")
	users, err := u.userService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.TraceStep(c, "list user done")
	common.ResponseSuccess(c, users)
}

// @Summary Update user | 修改用户信息
// @Description 修改用户信息并保存
// @Accept json
// @Produce json
// @Tags user
// @Security JWT
// @Param user body model.UpdatedUser true "user info"
// @Param id path  int true "user id"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/users/{id} [put]
func (u *UserController) Update(c *gin.Context) {
	// GetUser 获取 Context(当前登录) 中的 user
	loginUser := common.GetUser(c)

	// ①无 user 登录, 不修改 直接返回
	if loginUser == nil {
		common.ResponseFailed(c, http.StatusForbidden, nil)
		return
	}

	// ②有 user 登录
	// 看是不是自己修改自己，

	// 不是自己修改自己
	if strconv.Itoa(int(loginUser.ID)) != c.Param("id") {
		// 先去获取user的group等全部信息
		loginUser, err := u.userService.Get(strconv.Itoa(int(loginUser.ID)))
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}
		// 如果user是管理员，最终允许修改，如果user不是管理员，不允许修改，返回
		if !authorization.IsClusterAdmin(loginUser) {
			common.ResponseFailed(c, http.StatusForbidden, nil)
			return
		}
	}

	// 是自己修改自己，直接进行修改

	// 要修改的 user 信息
	new := new(model.UpdatedUser)
	if err := c.BindJSON(new); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	logrus.Infof("已获取修改的 user: %#v", new.Name)

	common.TraceStep(c, "start update user", trace.Field{Key: "user", Value: new.Name})
	defer common.TraceStep(c, "update user done", trace.Field{Key: "user", Value: new.Name})

	// 修改user信息
	user, err := u.userService.Update(c.Param("id"), new.GetUser())
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	common.ResponseSuccess(c, user)
}

// @Summary Delete user | 删除 user
// @Description Delete user and stroage | 删除 user 和存储
// @Produce json
// @Tags user
// @Security JWT
// @Param id path int true "user id"
// @Success 200 {object} common.Response
// @Router /api/v1/users/{id} [delete]
func (u *UserController) Delete(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil || (strconv.Itoa(int(user.ID))) != c.Param("id") && !authorization.IsClusterAdmin(user) {
		common.ResponseFailed(c, http.StatusBadRequest, nil)
		return
	}

	if err := u.userService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, nil)

}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/users", u.List)          // 用户列表
	api.POST("/users", u.Create)       // 创建用户
	api.GET("/users/:id", u.Get)       // 查询某个用户
	api.PUT("/users/:id", u.Update)    // 修改用户信息
	api.DELETE("/users/:id", u.Delete) // 删除 user
}

/**
 * @description: Name()返回控制器的名称
 * @return {*}
 */
func (u *UserController) Name() string {
	return "User"
}
