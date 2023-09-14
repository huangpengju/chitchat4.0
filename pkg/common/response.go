package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResponseFailed(c *gin.Context, code int, err error) {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	if code == http.StatusUnauthorized && c.Request != nil {
		if val, err := c.Cookie(CookieTokenName); err == nil && val != "" {
			c.SetCookie(CookieTokenName, "", -1, "/", "", true, true)
			c.SetCookie(CookieLoginUser, "", -1, "/", "", true, true)
		}
	}
	var msg string
	if err != nil {
		msg = err.Error()
		user := GetUser(c) //这里
		var name string
		if user != nil {
			name = user.name
		}
		var url string
		if c.Request != nil {
			url = c.Request.URL.String()
		}
		logrus.Warnf("url:%s,user:%s,err:%v", url, name, msg)
	}
	NewResponse(c, code, nil, msg)
}
