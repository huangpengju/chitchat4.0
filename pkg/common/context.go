package common

import (
	"chitchat4.0/pkg/model"
	"github.com/gin-gonic/gin"
)

// GetUser 判断用户
func GetUser(c *gin.Context) *model.User {
	if c == nil {
		return nil
	}
	val, ok := c.Get(UserContextKey)
	if !ok {
		return nil
	}
	user, ok := val.(*model.User)
	if !ok {
		return nil
	}
	return user
}

func TraceStep(c *gin.Context, msg string, fields ...trace.Field)
