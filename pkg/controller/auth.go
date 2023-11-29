package controller

import (
	"encoding/json"
	"net/http"

	"chitchat4.0/pkg/authentication"
	"chitchat4.0/pkg/authentication/oauth"
	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService  service.UserService        // user 服务
	jwtService   *authentication.JWTService // jwt服务
	oauthManager *oauth.OAuthManager        // 授权管理
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

// @Summary Login | 登录
// @Description user login | 用户登录
// @Accept json
// @Produce json
// @Tags auth
// @Param user body model.AuthUser true "auth user info"
// @Success 200 {object} common.Response{data=model.JWTToken}
// @Router /api/v1/auth/token [post]
func (ac *AuthController) Login(c *gin.Context) {
	// 准备把登录参数与结构体进行绑定
	auser := new(model.AuthUser)
	if err := c.BindJSON(auser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	var user *model.User
	var err error

	// 根据授权类型，判断是第三方登录，还是user登录
	if !oauth.IsEmptyAuthType(auser.AuthType) && auser.Name == "" {
		// 第三方登录
		// AuthType 不为 "" 时 ，a 是 !false | Name 为 ""时，b 是 true

		// GetAuthProvider() 根据授权类型，返回 提供者 provider
		provider, err := ac.oauthManager.GetAuthProvider(auser.AuthType)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}
		authToken, err := provider.GetToken(auser.AuthCode)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}

		userInfo, err := provider.GetUserInfo(authToken)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			return
		}

		// 第三方登录（不存在用户，则注册）
		user, err = ac.userService.CreateOAuthUser(userInfo.User())

	} else {
		// AuthType 为 "" 时 ，a 是 !true | Name 为 ""时， b 是 true
		// AuthType 为 "" 时 ，a 是 !true | Name 不为 ""时，b 是 false
		//

		// AuthType 不为 "" 时 ，a 是 !false | Name 不为 ""时，b 是 false

		// 使用登录用户的name查询用户是否存在，然后对比登录用户密码和数据库用户密码
		user, err = ac.userService.Auth(auser)
	}
	if err != nil {
		common.ResponseFailed(c, http.StatusUnauthorized, err)
		return
	}
	// 创建 token
	token, err := ac.jwtService.CreateToken(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	// json 序列化
	userJson, err := json.Marshal(user) // userJson 包括除了密码之外的用户信息
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	// 设置cookie
	// c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	// name：cookie 的名称（必须）。
	// value：cookie 的值（必须）。
	// maxAge：cookie 的过期时间，以秒为单位。如果为负数，则表示会话 cookie（在浏览器关闭之后删除），如果为零，则表示立即删除 cookie（可选，默认值为-1）。
	// path：cookie 的路径。如果为空字符串，则使用当前请求的 URI 路径作为默认值（可选，默认值为空字符串）。
	// domain：cookie 的域名。如果为空字符串，则不设置域名（可选，默认值为空字符串）。
	// secure：指定是否仅通过 HTTPS 连接发送 cookie。如果为 true，则仅通过 HTTPS 连接发送 cookie；否则，使用 HTTP 或 HTTPS 连接都可以发送 cookie（可选，默认值为 false）。
	// httpOnly：指定 cookie 是否可通过 JavaScript 访问。如果为 true，则无法通过 JavaScript 访问 cookie；否则，可以通过 JavaScript 访问 cookie（可选，默认值为 true）。

	if auser.SetCookie {
		c.SetCookie(common.CookieTokenName, token, 3600*24, "/", "", true, true)
		c.SetCookie(common.CookieLoginUser, string(userJson), 3600*24, "/", "", true, false)
	}
	common.ResponseSuccess(c, model.JWTToken{
		Token:    token,
		Describe: "set token in Authorization Header,[Authorization:Bearer {token}]",
	})
	// Authorization添加规范，如：Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFkbWluIiwiaXNzIjoiaHBqLmlvIiwiZXhwIjoxNzAxMjU5NTIzLCJuYmYiOjE3MDA2NTM3MjMsImp0aSI6IjEifQ.XXlKUnJZn59RmuNaHDGb-UTxpCNJd_fq_0ol0WB0KIg
	// Bearer 为前缀不能少
}

// @Summary Logout | 退出
// @Description User logout | User退出
// @Produce json
// @Tags auth
// @Success 200 {object} common.Response
// @Router /api/v1/auth/token [delete]
func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie(common.CookieTokenName, "", -1, "/", "", true, true)
	c.SetCookie(common.CookieLoginUser, "", -1, "/", "", true, false)
	common.ResponseSuccess(c, nil)
}

func (ac *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/auth/user", ac.Register)  // 注册用户
	api.POST("/auth/token", ac.Login)    //  用户登录
	api.DELETE("/auth/token", ac.Logout) // 退出
}

func (ac *AuthController) Name() string {
	return "Authentication"
}
