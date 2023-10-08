package common

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/utils/trace"
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

// GetTrace 获取痕迹
func GetTrace(c *gin.Context) *trace.Trace {
	if c == nil {
		return nil
	}
	val, ok := c.Get(TraceContextKey)
	if !ok {
		return nil
	}
	trace, ok := val.(*trace.Trace)
	if !ok {
		return nil
	}
	return trace
}

// TraceStep 痕迹步骤
func TraceStep(c *gin.Context, msg string, fields ...trace.Field) {
	trace := GetTrace(c)
	if trace != nil {
		trace.Step(msg, fields...)
	}
}
