package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"chitchat4.0/pkg/authentication"
	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/repository"
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware 验证Token中间件，同时把 user 存储到 gin的Context中
func AuthenticationMiddleware(jwtService *authentication.JWTService, userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		// header 获取 token
		token, _ := getTokenFromAuthorizationHeader(c)

		if token == "" {
			// request 获取 token
			token, _ = getTokenFromCookie(c)
		}
		// 解析 Token ,获取当前 user 信息
		user, _ := jwtService.ParseToken(token)

		// 判断用户
		if user != nil {
			// 使用 user.id 查询数据库，确认用户
			user, err := userRepo.GetUserByID(user.ID)
			if err != nil {
				common.ResponseFailed(c, http.StatusInternalServerError, fmt.Errorf("failed to get user"))
				c.Abort()
				return
			}

			// 在 gin 的 Context 中设置 user,后续可以使用 Context 中的 user
			common.SetUser(c, user)
		}
		c.Next()

	}
}

// getTokenFromCookie 在Request中的获取Cookie，
// c.Request.Cookie
func getTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie(common.CookieTokenName)
}

// getTokenFromAuthorizationHeader 在Header中的获取 Authorization 中的 Token
func getTokenFromAuthorizationHeader(c *gin.Context) (string, error) {
	// Header中的获取Cookie
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", nil
	}
	token := strings.Fields(auth)
	// Bearer {eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjcsIm5hbWUiOiLlkajkuIkiLCJpc3MiOiJocGouaW8iLCJleHAiOjE2OTgzMDYyMjMsIm5iZiI6MTY5NzcwMDQyMywianRpIjoiMjcifQ.WJGBl9WXwcBKCGOW7wnOCmmjZX-NCWvxsLZ4Vf55bQk}
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return "", fmt.Errorf("authorization header invaild")
	}
	return token[1], nil
}
