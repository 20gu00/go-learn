package gin_redirect

// 现在重定向一般前端做

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 永久 临时
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.Run(":8000") // 访问localhost:8000 -> 百度
}
