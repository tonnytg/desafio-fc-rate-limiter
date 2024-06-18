package entity

import "context"

type RateLimitRepository struct{}

type RateLimitRepositoryInterface interface {
	AllowIP(ctx context.Context, ip string) (bool, error)
	AllowToken(ctx context.Context, token string) (bool, error)
	Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) // Este método deve ser exportado
}

func NewRateLimitRepository() *RateLimitRepository {
	return &RateLimitRepository{}
}
