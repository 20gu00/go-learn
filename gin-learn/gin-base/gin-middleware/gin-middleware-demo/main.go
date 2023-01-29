package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件的语法
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我是一个中间件")
	}
}

func main() {
	r := gin.Default()
	// 中间件的全局注册
	//r.Use(MiddleWare())
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("执行了Get方法")
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	// 按顺序执行局部中间件如果由全局中间件先执行外部的全局中间件
	// 中间件局部注册
	r.GET("/hello2", MiddleWare, func(c *gin.Context) {
		fmt.Println("执行了Get方法")
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	r.Run(":8000")
}
