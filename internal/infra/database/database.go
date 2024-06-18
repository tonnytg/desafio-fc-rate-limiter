package database

import (
	"context"
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/config"
	"log"
)

type Database struct {
	client *Database
	config *config.Config
}

type DatabaseRepository struct {
	database Database
}

type DatabaseRepositoryInterface interface{}

func NewDatabaseRepository(config *config.Config) *DatabaseRepository {

	rr := DatabaseRepository{
		Database{
			&Database{},
			config,
		}}

	return &rr
}

func (dr *DatabaseRepository) Allow(ctx context.Context, key string, limit int, maxTokens int) (bool, error) {

	log.Println("DatabaseRepository received:", ctx, key, limit, maxTokens)
	return true, nil
}
