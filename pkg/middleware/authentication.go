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

func AuthenticationMiddleware(jwtService *authentication.JWTService, userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 token
		token, _ := getTokenFromAuthorizationHeader(c)
		if token == "" {
			token, _ = getTokenFromCookie(c)
		}
		// 解析 Token
		user, _ := jwtService.ParseToken(token)

		// 判断用户
		if user != nil {
			user, err := userRepo.GetUserByID(user.ID)
			if err != nil {
				common.ResponseFailed(c, http.StatusInternalServerError, fmt.Errorf("failed to get user"))
				c.Abort()
				return
			}
			// 在 content 中设置 user,在修改user信息时，会使用到
			common.SetUser(c, user)
		}
		c.Next()
	}
}

// getTokenFromCookie 在 Cookie 中获取 token
func getTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie("token")
}

// getTokenFromAuthorizationHeader 在 Authorization 中获取 Token
func getTokenFromAuthorizationHeader(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", nil
	}
	token := strings.Fields(auth)
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return "", fmt.Errorf("authorization header invaild")
	}
	return token[1], nil
}
