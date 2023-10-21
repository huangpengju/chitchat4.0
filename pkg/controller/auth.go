package controller

import (
	"encoding/json"
	"net/http"

	"chitchat4.0/pkg/authentication"
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

// @Summary Login | 登录
// @Description user login | 用户登录
// @Accept json
// @Produce json
// @Tags auth
// @Param user body model.AuthUser true "auth user info"
// @Success 200 {object} common.Response{data=model.JWTToken}
// @Router /api/v1/auth/token [post]
func (ac *AuthController) Login(c *gin.Context) {
	auser := new(model.AuthUser)
	if err := c.BindJSON(auser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	var user *model.User
	var err error
	// 判断授权类型是不是为空
	// if !oauth.IsEmptyAuthType(auser.AuthType) && auser.Name == "" {
	// 空
	// provider, err := ac.oauthManger.GetAuthProvider(auser.AuthType)
	// if err != nil{
	// 	common.ResponseFailed(c,http.StatusBadRequest,err)
	// 	return
	// }

	// 第三方登录
	// authToken, err := provider.GetToken(auser.AuthCode)
	// } else {
	// 使用登录用户的name查询用户是否存在，然后对比登录用户密码和数据库用户密码
	user, err = ac.userService.Auth(auser)
	// }
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
	// 	c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	// 第一个参数 key;
	// 第二个参数 value ;
	// 第三个参数 过期时间.如果只想设置 Cookie 的保存路径而不想设置存活时间，可以在第三个 参数中传递 nil ;
	// 第四个参数 cookie 的路径 ;
	// 第五个参数 cookie 的路径 Domain 作用域 本地调试配置成 localhost , 正式上线配置成域名 ;
	// 第六个参数是 secure ，当 secure 值为 true 时，cookie 在 HTTP 中是无效，在 HTTPS 中 才有效 ;
	// 第七个参数 httpOnly，表示 cookie 是否可以通过 js代码进行操作，为true时不能被js获取,是微软对 COOKIE 做的扩展。如果在 COOKIE 中设置了“httpOnly”属性， 则通过程序（JS 脚本、applet 等）将无法读取到 COOKIE 信息，防止 XSS 攻击产生;
	if auser.SetCookie {
		c.SetCookie(common.CookieTokenName, token, 3600*24, "/", "localhost", true, true)
		c.SetCookie(common.CookieLoginUser, string(userJson), 36000*24, "/", "localhost", true, false)
	}
	common.ResponseSuccess(c, model.JWTToken{
		Token:    token,
		Describe: "set token in Authorization Header,[Authorization:Bearer {token}]",
	})
}

func (ac *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/auth/user", ac.Register)
	api.POST("/auth/token", ac.Login)
}

func (ac *AuthController) Name() string {
	return "Authentication"
}
