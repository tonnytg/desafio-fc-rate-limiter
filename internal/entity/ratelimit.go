package entity

import (
	"time"
)

type RateLimit struct {
	clientRepository  ClientRepositoryInterface
	rateLimitIP       int
	rateLimitToken    int
	blockTime         time.Duration
	refillInterval    time.Duration
	tokensPerRefill   int
	maxTokensPerIP    int
	maxTokensPerToken int
}

type RateLimitInterface interface{}

func NewRateLimit(clientRepository ClientRepositoryInterface, rateLimitIP, rateLimitToken, refillInterval, tokensPerRefill, maxTokensPerIP, maxTokensPerToken, blockTime int) *RateLimit {
	return &RateLimit{
		clientRepository:  clientRepository,
		rateLimitIP:       rateLimitIP,
		rateLimitToken:    rateLimitToken,
		refillInterval:    time.Duration(refillInterval) * time.Second,
		tokensPerRefill:   tokensPerRefill,
		maxTokensPerIP:    maxTokensPerIP,
		maxTokensPerToken: maxTokensPerToken,
		blockTime:         time.Duration(blockTime) * time.Second,
	}
}
