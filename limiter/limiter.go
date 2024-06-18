package limiter

import (
	"context"
	"fmt"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/entity"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	client            entity.ClientRepositoryInterface
	rateLimitIP       int
	rateLimitToken    int
	blockTime         time.Duration
	refillInterval    time.Duration
	tokensPerRefill   int
	maxTokensPerIP    int
	maxTokensPerToken int
}

func NewRateLimiter(clientRepository entity.ClientRepositoryInterface, rateLimitIP, rateLimitToken, refillInterval, tokensPerRefill, maxTokensPerIP, maxTokensPerToken, blockTime int) *RateLimiter {
	return &RateLimiter{
		client:            clientRepository,
		rateLimitIP:       rateLimitIP,
		rateLimitToken:    rateLimitToken,
		refillInterval:    time.Duration(refillInterval) * time.Second,
		tokensPerRefill:   tokensPerRefill,
		maxTokensPerIP:    maxTokensPerIP,
		maxTokensPerToken: maxTokensPerToken,
		blockTime:         time.Duration(blockTime) * time.Second,
	}
}

func (rl *RateLimiter) AllowIP(ctx context.Context, ip string) (bool, error) {
	key := fmt.Sprintf("rl:ip:%s", ip)
	return rl.allow(ctx, key, rl.rateLimitIP, rl.maxTokensPerIP)
}

func (rl *RateLimiter) AllowToken(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf("rl:token:%s", token)
	return rl.allow(ctx, key, rl.rateLimitToken, rl.maxTokensPerToken)
}

func (rl *RateLimiter) allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {
	pipe := rl.client.TxPipeline()

	// Get the current number of tokens and the last refill time
	tokens, err := rl.client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}

	// If there are no tokens, set the initial token count to the limit
	if err == redis.Nil {
		tokens = limit
	}

	// If the current tokens are less than the limit, refill tokens
	if tokens < maxTokens {
		pipe.IncrBy(ctx, key, int64(rl.tokensPerRefill))
		tokens += rl.tokensPerRefill
		if tokens > maxTokens {
			tokens = maxTokens
		}
	}

	// Decrement the token count
	pipe.Decr(ctx, key)
	pipe.Expire(ctx, key, rl.refillInterval)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	tokens--
	return tokens >= 0, nil
}
