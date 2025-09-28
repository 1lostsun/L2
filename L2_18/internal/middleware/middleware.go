package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Force400Middleware : форс четырехсотой ошибки на все ошибки от 400 до 499
func Force400Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		if status >= 400 && status < 500 {
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
