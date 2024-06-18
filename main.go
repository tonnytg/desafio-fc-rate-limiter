package main

import (
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/config"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/infra/database"
	"github.com/tonnytg/desafio-fc-rate-limiter/pkg/middleware"
	"log"

	"github.com/tonnytg/desafio-fc-rate-limiter/limiter"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//repo := database.NewDatabaseRepository(cfg)
	//serv := entity.NewClientService(repo)

	// TODO: fazer algo com o service

	clientRepository := database.NewRedisRepository(cfg)
	rateLimiter := limiter.NewRateLimiter(clientRepository, cfg.RateLimitIP, cfg.RateLimitToken, cfg.RefillInterval, cfg.TokensPerRefill, cfg.MaxTokensPerIP, cfg.MaxTokensPerToken, cfg.BlockTime)

	middleware.StartMiddlware(rateLimiter)
}
