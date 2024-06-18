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
	clientRedis *ClientRedis
}

type RedisRepositoryInterface interface {
	Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error)
	Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) // Este método deve ser exportado
}

func NewRedisRepository(config *config.Config) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     config.ClientAddr,
		Password: config.ClientPassword,
		DB:       config.ClientDB,
	})

	return &RedisRepository{
		clientRedis: &ClientRedis{
			client: client,
			config: config,
		},
	}
}

func (rr *RedisRepository) Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {
	return rr.Allow(ctx, key, limit, maxTokens)
}

func (rr *RedisRepository) AllowIP(ctx context.Context, ip string) (bool, error) {
	return true, nil
}
func (rr *RedisRepository) AllowToken(ctx context.Context, token string) (bool, error) {
	return true, nil
}

func (rr *RedisRepository) Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {
	pipe := rr.clientRedis.client.TxPipeline()

	// Get the current number of tokens and the last refill time
	tokens, err := rr.clientRedis.client.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false, err
	}

	// If there are no tokens, set the initial token count to the limit
	if err == redis.Nil {
		tokens = limit
	}

	// If the current tokens are less than the limit, refill tokens
	if tokens < maxTokens {
		pipe.IncrBy(ctx, key, int64(rr.clientRedis.config.TokensPerRefill))
		tokens += rr.clientRedis.config.TokensPerRefill
		if tokens > maxTokens {
			tokens = maxTokens
		}
	}

	// Decrement the token count
	pipe.Decr(ctx, key)
	pipe.Expire(ctx, key, time.Duration(rr.clientRedis.config.RefillInterval)*time.Second)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	tokens--
	return tokens >= 0, nil
}
