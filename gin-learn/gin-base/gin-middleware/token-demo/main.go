package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//去header里面拿token(参数param token:xxxxxx)
		token := c.Request.Header.Get("token")
		fmt.Println("获取token", token)
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "身份验证不通过",
			})
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(AuthMiddleware())
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "home 页面",
		})
	})
	r.Run(":8000")
}
