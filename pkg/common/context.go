package common

import (
	"chitchat4.0/pkg/model"
	"chitchat4.0/pkg/utils/request"
	"chitchat4.0/pkg/utils/trace"
	"github.com/gin-gonic/gin"
)

// SetUser 在 Context 中设置 user
func SetUser(c *gin.Context, user *model.User) {
	if c == nil || user == nil {
		return
	}
	c.Set(UserContextKey, user)
}

// GetUser 判断 Context 中是否有 user
func GetUser(c *gin.Context) *model.User {
	if c == nil {
		return nil
	}
	// c.Get 返回给定键的值和true,不存在返回nil和false
	// UserContextKey 表示 user
	val, ok := c.Get(UserContextKey)
	if !ok {
		return nil
	}

	// val 进行类型断言
	user, ok := val.(*model.User)
	if !ok {
		return nil
	}

	return user
}

// GetTrace 获取追踪钥匙
func GetTrace(c *gin.Context) *trace.Trace {
	if c == nil {
		return nil
	}
	// Get返回给定键的值，即:(value, true)。如果值不存在，则返回(nil, false)
	val, ok := c.Get(TraceContextKey)
	if !ok {
		return nil
	}
	// val 进行类型断言，此处是指针类型断言
	trace, ok := val.(*trace.Trace)
	if !ok {
		return nil
	}
	return trace
}

func SetTrace(c *gin.Context, t *trace.Trace) {
	if c == nil || t == nil {
		return
	}
	c.Set(TraceContextKey, t)
}

// TraceStep 追踪步骤，
// 参数1：c *gin.Context，
// 参数2：start create user，
// 参数3： trace.Field{Key   string	Value interface{}}
func TraceStep(c *gin.Context, msg string, fields ...trace.Field) {
	trace := GetTrace(c)
	if trace != nil {
		trace.Step(msg, fields...)
	}
}

func SetRequestInfo(c *gin.Context, ri *request.RequestInfo) {
	if c == nil || ri == nil {
		return
	}

	c.Set(RequestInfoContextKey, ri)
}
