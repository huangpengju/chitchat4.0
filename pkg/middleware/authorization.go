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
		user := common.GetUser(c)

		if user == nil {
			user = &model.User{}
		}

		ri := common.GetRequestInfo(c)
		if ri == nil {
			common.ResponseFailed(c, http.StatusBadRequest, fmt.Errorf("failed to get request info"))
			c.Abort()
			return
		}

		if ri.IsResourceRequest {
			resource := ri.Resource
			// 授权
			ok, err := authorization.Authorize(user, ri)
			if err != nil {
				common.ResponseFailed(c, http.StatusInternalServerError, err)
				c.Abort()
				return
			}
			logrus.Infof("authorize user [%s(%d)], namespace [%s] resource [%s(%s)] verb [%s], result: %t",
				user.Name, user.ID, ri.Namespace, ri.Resource, ri.Name, ri.Verb, ok)

			if !ok {
				if user.Name == "" {
					common.ResponseFailed(c, http.StatusUnauthorized, nil)
				} else {
					common.ResponseFailed(c, http.StatusForbidden, fmt.Errorf("user [%s] is forbidden for resource %s in namespace %s", user.Name, resource, ri.Namespace))
				}
				c.Abort()
				return
			}
		}
		c.Next()

	}
}
