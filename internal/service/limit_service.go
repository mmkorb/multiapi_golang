package service

import (
	"internal/domain"
	"internal/repository"
)

// LimitService is the service that handles the business logic for limits
type LimitService struct {
	Repository repository.LimitRepository
}

// NewLimitService creates a new instance of LimitService
func NewLimitService(repo repository.LimitRepository) *LimitService {
	return &LimitService{
		Repository: repo,
	}
}

// GetLimits returns all limits stored in the database
func (s *LimitService) GetLimits() ([]*domain.Limit, error) {
	return s.Repository.GetAll()
}
