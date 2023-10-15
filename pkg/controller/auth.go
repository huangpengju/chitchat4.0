package controller

import (
	"net/http"

	"chitchat4.0/pkg/authentication"
	"chitchat4.0/pkg/authentication/oauth"
	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService service.UserService
	jwtService  *authentication.JWTService
	// oauthManger *oauth.OAuthManger
}

func NewAuthController(userService service.UserService, jwtService *authentication.JWTService) Controller {
	return &AuthController{
		userService: userService,
		jwtService:  jwtService,
		// oauthManger: oauthManager,
	}
}

// @Summary Register user | 注册用户
// @Description Create user and storage
// @Accept json
// @Produce json
// @Tags auth
// @Param user body model.CreatedUser true "user info"
// @Success 200 {object} common.Response{data=model.User}
// @Router /api/v1/auth/user [post]
func (ac *AuthController) Register(c *gin.Context) {
	createdUser := new(model.CreatedUser)
	if err := c.BindJSON(createdUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	// CreatedUser 赋值给 User
	user := createdUser.GetUser()
	if err := ac.userService.Validate(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	// 默认值处理
	ac.userService.Default(user)
	user, err := ac.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		// return  需要确认是否添加
	}
	common.ResponseSuccess(c, user)
}

func (an *AuthController) Login(c *gin.Context) {
	auser := new(model.AuthUser)
	if err := c.BindJSON(auser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	var user *model.User
	var err error
	// 判断授权类型是不是为空
	if !oauth.IsEmptyAuthType(auser.AuthType) && auser.Name == "" {
		// 空
		// provider, err := ac.oauthManger.GetAuthProvider(auser.AuthType)
		// if err != nil{
		// 	common.ResponseFailed(c,http.StatusBadRequest,err)
		// 	return
		// }

		// 第三方登录
		authToken, err := provider.GetToken(auser.AuthCode)
	} else {
		// 账户密码登录
	}
}

func (ac *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/auth/user", ac.Register)
	api.POST("/auth/token", ac.Login)
}

func (ac *AuthController) Name() string {
	return "Authentication"
}
