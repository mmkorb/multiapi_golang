package service

import (
	"internal/domain"
	"internal/repository"
)

// RequestService is the service that handles the business logic for requests
type RequestService struct {
	Repository repository.RequestRepository
}

// NewRequestService creates a new instance of RequestService
func NewRequestService(repo repository.RequestRepository) *RequestService {
	return &RequestService{
		Repository: repo,
	}
}

// CreateRequest creates a new request and saves it in Datastore
func (s *RequestService) CreateRequest(description string, status string) (*domain.Request, error) {
	// Create the request
	request := domain.NewRequest(0, description, status)

	// Save the request
	err := s.Repository.Save(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
