package entity

import (
	"context"
	"fmt"
	"time"
)

type RateLimitService struct {
	repo              RateLimitRepositoryInterface
	rateLimitIP       int
	rateLimitToken    int
	blockTime         time.Duration
	refillInterval    time.Duration
	tokensPerRefill   int
	maxTokensPerIP    int
	maxTokensPerToken int
}

type RateLimitServiceInterface interface {
	AllowIP(ctx context.Context, ip string) (bool, error)
	AllowToken(ctx context.Context, token string) (bool, error)
}

func NewRateLimitService(repo RateLimitRepositoryInterface, rateLimitIP, rateLimitToken, refillInterval, tokensPerRefill, maxTokensPerIP, maxTokensPerToken, blockTime int) *RateLimitService {
	return &RateLimitService{
		repo:              repo,
		rateLimitIP:       rateLimitIP,
		rateLimitToken:    rateLimitToken,
		refillInterval:    time.Duration(refillInterval) * time.Second,
		tokensPerRefill:   tokensPerRefill,
		maxTokensPerIP:    maxTokensPerIP,
		maxTokensPerToken: maxTokensPerToken,
		blockTime:         time.Duration(blockTime) * time.Second,
	}
}

func (rl *RateLimitService) AllowIP(ctx context.Context, ip string) (bool, error) {
	key := fmt.Sprintf("rl:ip:%s", ip)
	return rl.repo.Allow(ctx, key, rl.rateLimitIP, rl.maxTokensPerIP)
}

func (rl *RateLimitService) AllowToken(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf("rl:token:%s", token)
	return rl.repo.Allow(ctx, key, rl.rateLimitToken, rl.maxTokensPerToken)
}
