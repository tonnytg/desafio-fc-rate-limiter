package entity

import "context"

type ClientRepository struct {
	client *Client
}

type ClientRepositoryInterface interface {
	Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error)
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}
