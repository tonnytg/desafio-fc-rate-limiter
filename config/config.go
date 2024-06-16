package config

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	RedisAddr         string
	RedisPassword     string
	RedisDB           int
	RateLimitIP       int
	RateLimitToken    int
	BlockTime         int
	RefillInterval    int
	TokensPerRefill   int
	MaxTokensPerIP    int
	MaxTokensPerToken int
}

func LoadConfig() (*Config, error) {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	rateLimitIP, err := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	if err != nil {
		return nil, err
	}

	rateLimitToken, err := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	if err != nil {
		return nil, err
	}

	blockTime, err := strconv.Atoi(os.Getenv("BLOCK_TIME"))
	if err != nil {
		return nil, err
	}

	refillInterval, err := strconv.Atoi(os.Getenv("REFILL_INTERVAL"))
	if err != nil {
		return nil, err
	}

	tokensPerRefill, err := strconv.Atoi(os.Getenv("TOKENS_PER_REFILL"))
	if err != nil {
		return nil, err
	}

	maxTokensPerIP, err := strconv.Atoi(os.Getenv("MAX_TOKENS_PER_IP"))
	if err != nil {
		return nil, err
	}

	maxTokensPerToken, err := strconv.Atoi(os.Getenv("MAX_TOKENS_PER_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &Config{
		RedisAddr:         os.Getenv("REDIS_ADDR"),
		RedisPassword:     os.Getenv("REDIS_PASSWORD"),
		RedisDB:           db,
		RateLimitIP:       rateLimitIP,
		RateLimitToken:    rateLimitToken,
		BlockTime:         blockTime,
		RefillInterval:    refillInterval,
		TokensPerRefill:   tokensPerRefill,
		MaxTokensPerIP:    maxTokensPerIP,
		MaxTokensPerToken: maxTokensPerToken,
	}, nil
}

func NewRedisClient(config *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
}
