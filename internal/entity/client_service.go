package entity

type ClientService struct {
	repo ClientRepositoryInterface
}

type ClientServiceInterface interface{}

func NewClientService(repo ClientRepositoryInterface) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
