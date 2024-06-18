package config

import (
	"os"
	"strconv"
)

type Config struct {
	ClientAddr        string
	ClientPassword    string
	ClientDB          int
	RateLimitIP       int
	RateLimitToken    int
	BlockTime         int
	RefillInterval    int
	TokensPerRefill   int
	MaxTokensPerIP    int
	MaxTokensPerToken int
}

func LoadConfig() (*Config, error) {

	db, err := strconv.Atoi(os.Getenv("CLIENT_DB"))
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
		ClientAddr:        os.Getenv("CLIENT_ADDR"),
		ClientPassword:    os.Getenv("CLIENT_PASSWORD"),
		ClientDB:          db,
		RateLimitIP:       rateLimitIP,
		RateLimitToken:    rateLimitToken,
		BlockTime:         blockTime,
		RefillInterval:    refillInterval,
		TokensPerRefill:   tokensPerRefill,
		MaxTokensPerIP:    maxTokensPerIP,
		MaxTokensPerToken: maxTokensPerToken,
	}, nil
}
