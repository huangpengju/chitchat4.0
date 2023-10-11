package middleware

import (
	"time"

	"chitchat4.0/pkg/common"
	utiltrace "chitchat4.0/pkg/utils/trace"
	"github.com/bombsimon/logrusr/v2"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace := utiltrace.New("Handler",
			logrusr.New(logrus.StandardLogger()),
			utiltrace.Field{Key: "method", Value: c.Request.Method},
			utiltrace.Field{Key: "path", Value: c.Request.URL.Path},
		)

		defer trace.LogIfLong(100 * time.Millisecond)

		common.SetTrace(c, trace)

		c.Next()
	}
}
