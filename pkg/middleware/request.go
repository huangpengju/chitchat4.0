package middleware

import (
	"net/http"

	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/utils/request"
	"github.com/gin-gonic/gin"
)

// RequestInfoMiddleware 是筛选需要的 http.Request 请求处理信息
func RequestInfoMiddleware(resolver request.RequestInfoResolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 把 http.Request 中的部分数据转移到自定结构体 RequestInfo 中，并返回
		ri, err := resolver.NewRequestInfo(c.Request)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}
		// 把 ri(RequestInfo) 存储到 Context 的Keys[RequestInfo]中
		common.SetRequestInfo(c, ri)

		c.Next()
	}

}
