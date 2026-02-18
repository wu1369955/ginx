package midware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//t := time.Now()
		c.Set("example", "example")
		// 请求前
		c.Next()
		// 请求后
		//latency := time.Since(t)
		//log.Println(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
