package entity

import (
	"github.com/tonnytg/desafio-fc-rate-limiter/internal/config"
)

type Client struct {
	config *config.Config
}

type ClientInterface interface{}

func NewClient(config *config.Config) *Client {
	return &Client{
		config,
	}
}

func (c *Client) AddCreditToken() bool {
	return true
}

func (c *Client) RemoveCreditToken() bool {
	return true
}

func (c *Client) AddCreditsIP() bool {
	return true
}

func (c *Client) RemoveCreditsIP() bool {
	return true
}
