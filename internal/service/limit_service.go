package service

import (
	"multiapi_golang/internal/domain"
	"multiapi_golang/internal/repository"
	"multiapi_golang/pkg/rabbitmq"
)

// LimitService representa o serviço de limites
type LimitService struct {
	repo     *repository.LimitRepository
	producer *rabbitmq.Publisher
}

// NewLimitService cria uma nova instância do serviço
func NewLimitService(repo *repository.LimitRepository, producer *rabbitmq.Publisher) *LimitService {
	return &LimitService{repo: repo, producer: producer}
}

// GetAllLimits retorna todos os registros da tabela "limit"
func (s *LimitService) GetAllLimits() ([]domain.Limit, error) {
	return s.repo.GetAll()
}

// CreateLimit insere um novo registro e publica no RabbitMQ
func (s *LimitService) CreateLimit(limit domain.Limit) error {
	// Salva no banco de dados
	if err := s.repo.Save(limit); err != nil {
		return err
	}

	// Publica no RabbitMQ
	err := s.producer.Publish(limit.ToString())
	return err
}
