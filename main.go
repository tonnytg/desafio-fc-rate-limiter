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

	repo := database.NewDatabaseRepository(cfg) // No Limiter Database Cache all request will be allowed
	//repo := database.NewRedisRepository(cfg) // Has Limiter Database Cache some request will be blocked
	rateLimitService := entity.NewRateLimitService(
		repo,
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
