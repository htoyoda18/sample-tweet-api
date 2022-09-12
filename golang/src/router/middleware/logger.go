package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Setup()

		c.Next()
	}
}
