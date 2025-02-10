package repository

import (
	"internal/domain"
	"log"
	"pkg/db"
)

// LimitRepository lida com o acesso ao banco de dados para recuperar limites
type LimitRepository struct{}

// NewLimitRepository cria uma nova instância do LimitRepository
func NewLimitRepository() *LimitRepository {
	return &LimitRepository{}
}

// GetAll recupera todos os limites armazenados no banco de dados
func (r *LimitRepository) GetAll() ([]*domain.Limit, error) {
	// Usando a conexão singleton configurada no pkg/db
	db := db.GetDB()
	rows, err := db.Query("SELECT id, name, value FROM limits")
	if err != nil {
		log.Printf("Error fetching limits: %v", err)
		return nil, err
	}
	defer rows.Close()

	var limits []*domain.Limit
	for rows.Next() {
		var limit domain.Limit
		if err := rows.Scan(&limit.ID, &limit.Name, &limit.Value); err != nil {
			log.Printf("Error scanning limit: %v", err)
			return nil, err
		}
		limits = append(limits, &limit)
	}

	return limits, nil
}
