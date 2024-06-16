package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tonnytg/desafio-fc-rate-limiter/config"
	"github.com/tonnytg/desafio-fc-rate-limiter/limiter"
	"github.com/tonnytg/desafio-fc-rate-limiter/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := config.NewRedisClient(cfg)
	rateLimiter := limiter.NewRateLimiter(client, cfg.RateLimitIP, cfg.RateLimitToken, cfg.RefillInterval, cfg.TokensPerRefill, cfg.MaxTokensPerIP, cfg.MaxTokensPerToken, cfg.BlockTime)

	r := gin.Default()
	r.Use(middleware.RateLimiterMiddleware(rateLimiter))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.Run(":8080")
}
