package middleware

import (
	logger "github.com/1lostsun/L2/tree/main/L2_18/internal/pkg"
	"github.com/gin-gonic/gin"
	"time"
)

// GinLogger : middleware логгирования Gin функций
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		switch {
		case statusCode >= 500 && statusCode <= 599:
			logger.Error("%s %s %d %s (ip=%s)", method, path, statusCode, latency, clientIP)
		case statusCode >= 400 && statusCode <= 599:
			logger.Warn("%s %s %d %s (%s)", method, path, 400, latency, clientIP)
		default:
			logger.Info("%s %s %d %s (%s)", method, path, statusCode, latency, clientIP)
		}
	}
}
