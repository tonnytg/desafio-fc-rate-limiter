package main

import (
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/config"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/entity"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/infra/database"
	"github.com/tonnytg/desafio-fc-rate-limiter/pkg/middleware"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	redisRepo := database.NewRedisRepository(cfg)
	rateLimitService := entity.NewRateLimitService(
		redisRepo,
		cfg.RateLimitIP,
		cfg.RateLimitToken,
		cfg.RefillInterval,
		cfg.TokensPerRefill,
		cfg.MaxTokensPerIP,
		cfg.MaxTokensPerToken,
		cfg.BlockTime,
	)

	middleware.StartMiddlware(rateLimitService)
}
