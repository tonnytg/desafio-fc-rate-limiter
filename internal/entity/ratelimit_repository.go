package entity

import "context"

type RateLimitRepository struct{}

type RateLimitRepositoryInterface interface {
	Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error)
}

func NewRateLimitRepository() *RateLimitRepository {
	return &RateLimitRepository{}
}

func (rr *RateLimitRepository) Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {
	return true, nil
}
