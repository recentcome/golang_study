package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//1创建路由
	r := gin.Default()
	//2、绑定路由规则执行的函数 gin.Context
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	r.GET("/reg/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post ok")
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(200, name+" is "+action)
	})
	//3、监听端口 默认在8080
	r.GET("/reg/put")
	r.Run(":8000")
}
