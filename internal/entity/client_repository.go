package entity

import "context"

type ClientRepository struct {
	client *Client
}

type ClientRepositoryInterface interface {
	Save(ctx context.Context, key string, limit int, maxTokens int) (bool, error)
}

func NewClientRepository(client *Client) *ClientRepository {
	return &ClientRepository{
		client: client,
	}
}
