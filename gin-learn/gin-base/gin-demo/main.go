package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化gin对象
	r := gin.Default()
	// 定义路由规则，gin.Context，封装了request和response
	r.GET("/api/gin-test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	//启动服务
	r.Run(":8000")
}
