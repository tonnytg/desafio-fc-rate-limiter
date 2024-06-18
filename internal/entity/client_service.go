package entity

import "github.com/tonnytg/desafio-fc-rate-limiter/internal/infra/database"

type ClientService struct {
	repo ClientRepositoryInterface
}

type ClientSerivceInterface interface{}

func NewClientService(dr database.DatabaseRepository) *ClientService {
	return &ClientService{
		dr,
	}
}
