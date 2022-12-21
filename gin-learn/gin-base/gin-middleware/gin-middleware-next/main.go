package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 洋葱结构,handlerfunc->中间件next
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("log start")
		//直接跳转到下一个handler，handler可能是中间件，也可能是路由处理函数
		//c.Next()
		fmt.Println("log end")
		//直接退出,后面的handler就不执行了(中间件,handlerFunc)
		c.Abort()
	}
}

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("requestid start")
		c.Next()
		fmt.Println("requestid end")
	}
}

func main() {
	r := gin.Default()
	r.Use(Log(), RequestId())
	r.GET("/", func(c *gin.Context) {
		fmt.Println("app running...")
		c.JSON(http.StatusOK, "hello World!")
	})
	r.Run(":8000")
}
