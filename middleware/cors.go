package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type" }
	config.AllowOrigins = []string{"http://localhost:80"}
	config.MaxAge = 12 * time.Hour
	config.AllowCredentials = true
	return cors.New(config)
}