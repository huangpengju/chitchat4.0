package middleware

import (
	"fmt"
	"net/http"

	"chitchat4.0/pkg/authorization"
	"chitchat4.0/pkg/common"
	"chitchat4.0/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Context中提取user
		user := common.GetUser(c)

		if user == nil {
			user = &model.User{}
		}

		// 从 Context 中的Keys 提取 [requestInfo]（http请求中的部分信息）
		ri := common.GetRequestInfo(c)
		if ri == nil {
			common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("failed to get request info"))
			c.Abort()
			return
		}
		// 是否资源请求
		if ri.IsResourceRequest {
			// 资源
			resource := ri.Resource
			// 授权 user 和 user 对应的请求（待研究：具体是如何把当前的请求，放入到对应的group中）
			ok, err := authorization.Authorize(user, ri)

			if err != nil {
				// 授权出错
				common.ResponseFailed(c, http.StatusInternalServerError, err)
				c.Abort()
				return
			}
			// 输出Info日志
			// 授权 user admin(1), 命名空间是啥？（待研究），请求资源(待研究)，请求name（待研究），请求的方法，true
			logrus.Infof("authorize user [%s(%d)], namespace [%s] resource [%s(%s)] verb [%s], result: %t",
				user.Name, user.ID, ri.Namespace, ri.Resource, ri.Name, ri.Verb, ok)

			// 授权失败
			if !ok {
				if user.Name == "" {
					common.ResponseFailed(c, http.StatusUnauthorized, nil)
				} else {
					// 禁止用户[%s]用于命名空间%s中的资源%s
					common.ResponseFailed(c, http.StatusForbidden, fmt.Errorf("user [%s] is forbidden for resource %s in namespace %s", user.Name, resource, ri.Namespace))
				}
				c.Abort()
				return
			}
		}
		c.Next()

	}
}
