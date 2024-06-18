package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tonnytg/desafio-fc-rate-limiter/limiter"
	"net/http"
)

func RateLimiterMiddleware(rl *limiter.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")

		var allowed bool
		var err error

		if token != "" {
			allowed, err = rl.AllowToken(c.Request.Context(), token)
		} else {
			allowed, err = rl.AllowIP(c.Request.Context(), ip)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "you have reached the maximum number of requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func StartMiddlware(rateLimiter *limiter.RateLimiter) {
	r := gin.Default()
	r.Use(RateLimiterMiddleware(rateLimiter))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.Run(":8080")
}
