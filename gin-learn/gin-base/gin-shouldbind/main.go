package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string `json:"username" binding:"required" form:"username"` // 不带请求体form-data(query ?参数) 带请求体application/json
	Password string `json:"password" binding:"required" form:"password"`
}

// form json都可以搭配各种bind使用
// form搭配bind等,其实就是类似? c.Query
func LoginHandler(c *gin.Context) {
	// 初始化,不初始化,结构体的零值是个成员的零值,复合类型
	login := &Login{}
	// GET请求一般使用上面的form标签(GET不够安全)
	// application/json这种请求，POST(登录) PUT(修改) DELETE(删除)都是用json标签
	// bind shouldbindjson
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "绑定参数失败" + err.Error(),
			"data": nil,
			"code": 90400,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "接口请求成功",
		"data": login,
		"code": 90200,
	})
}

func main() {
	r := gin.Default()
	r.POST("/login", LoginHandler)
	r.Run(":8000")
}
