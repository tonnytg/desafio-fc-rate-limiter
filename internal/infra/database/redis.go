package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/config"
	"time"
)

type ClientRedis struct {
	client *redis.Client
	config *config.Config
}

type RedisRepository struct {
	clientRedis ClientRedis
}

type RedisRepositoryInterface interface {
	Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error)
}

func NewRedisRepository(config *config.Config) *RedisRepository {

	rr := RedisRepository{
		ClientRedis{
			redis.NewClient(&redis.Options{
				Addr:     config.ClientAddr,
				Password: config.ClientPassword,
				DB:       config.ClientDB,
			}),
			config,
		}}

	return &rr
}

func (cr *ClientRedis) Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {
	pipe := cr.client.TxPipeline()

	// Get the current number of tokens and the last refill time
	tokens, err := cr.client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}

	// If there are no tokens, set the initial token count to the limit
	if err == redis.Nil {
		tokens = limit
	}

	// If the current tokens are less than the limit, refill tokens
	if tokens < maxTokens {
		pipe.IncrBy(ctx, key, int64(cr.config.TokensPerRefill))
		tokens += cr.config.TokensPerRefill
		if tokens > maxTokens {
			tokens = maxTokens
		}
	}

	// Decrement the token count
	pipe.Decr(ctx, key)
	pipe.Expire(ctx, key, time.Duration(cr.config.RefillInterval)*time.Second)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	tokens--
	return tokens >= 0, nil
}
