package controller

import (
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"chitchat4.0/pkg/utils/trace"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) Controller {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) List(c *gin.Context) {

}

// @Summary 创建 user
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
	}
	// user 邮箱为空时，设置默认邮箱
	u.userService.Default(user)
	common.TraceStep(c, "start create user", trace.Field{"user", user.Name})
	defer common.TraceStep(c, "create user done", trace.Field{"user", user.Name})
	user, err := u.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
	}
	common.ResponseSuccess(c, user)
}

func (u *UserController) Get(c *gin.Context) {

}

func (u *UserController) Update(c *gin.Context) {

}

func (u *UserController) Delete(c *gin.Context) {

}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	// api.GET("/users", u.List)
	api.POST("/users", u.Create)
	// api.GET("/users/:id", u.Get)
	// api.PUT("/users:id", u.Update)
	// api.DELETE("/users/:id", u.Delete)
}

func (u *UserController) Name() string {
	return "User"
}
