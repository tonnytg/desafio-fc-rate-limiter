package entity

type RateLimitService struct{}

type RateLimitServiceInterface interface{}

func NewRateLimitService() *RateLimitService {
	return &RateLimitService{}
}
