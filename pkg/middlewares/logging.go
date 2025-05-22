package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func (c * gin.Context){
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Printf("Logger...[%s] %s %s (%v)", c.Request.Method, c.Request.URL.Path, c.ClientIP(), duration)
	}
}