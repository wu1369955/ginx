package middlewares

import "github.com/gin-gonic/gin"

// SetupMiddlewares 设置所有中间件
func SetupMiddlewares(router *gin.Engine) {
	router.Use(Logger())
	router.Use(RequestID())
	router.Use(Cors())
	router.Use(Session())
	router.Use(HTTPSRedirect())
	router.Use(Compress())
	router.Use(RateLimit())
	router.Use(ErrorHandler())
}
