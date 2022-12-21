package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/book/:id", GetBookDetailHandler) // url参数  /book/1  c.Param  用的不多
	r.GET("/user", GetUserDetailHandler)     // c.Query ?a=b&a1=b1  参数输入
	// 也可以结合 /book/1?a=b
	r.Run(":8000")
}

func GetBookDetailHandler(c *gin.Context) {
	bookId := c.Param("id")
	username := c.Query("name")
	// c.String
	c.String(http.StatusOK, fmt.Sprintf("成功获取书籍详情：%s %s", bookId, username))
}

func GetUserDetailHandler(c *gin.Context) {
	username := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("成功获取用户详情：%s", username))
}
