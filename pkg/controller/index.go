package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Home
// @Description 返回后端主页 html 源代码
// @Produce html
// @Tags home
// @Router /index [get]
func Index(c *gin.Context) {
	c.Data(http.StatusOK, "text/html;charset=utf-8", []byte(
		`<html>
		<head>
			<title>后端 Server</title>
		</head>
		<body>
			<h1>Hello 黄鹏举</h1>
			<ul>
				<li><a href="/swagger/index.html">swagger</a></li>
				<li><a href="/metrics">metrics</a></li>
				<li><a href="/healthz">healthz</a></li>
				<li><a href="/">api list</a></li>
			  </ul>
			<hr>
			<center>版本/1.0</center>
		</body>
	<html>`))
}
