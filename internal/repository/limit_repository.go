package repository

import (
	"database/sql"
	"log"
	"multiapi_golang/internal/domain"
	"multiapi_golang/pkg/db"
)

// LimitRepository define métodos para acessar a tabela "limit"
type LimitRepository struct {
	db *sql.DB
}

// NewLimitRepository cria uma instância do repositório com conexão ao banco
func NewLimitRepository() *LimitRepository {
	return &LimitRepository{db: db.DB}
}

// GetAll retorna todos os registros da tabela "limit"
func (r *LimitRepository) GetAll() ([]domain.Limit, error) {
	query := `SELECT id, nome, valor FROM limit`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Erro ao buscar todos os limites: %v", err)
		return nil, err
	}
	defer rows.Close()

	var limits []domain.Limit
	for rows.Next() {
		var limit domain.Limit
		if err := rows.Scan(&limit.ID, &limit.Nome, &limit.Valor); err != nil {
			log.Printf("Erro ao escanear linha: %v", err)
			return nil, err
		}
		limits = append(limits, limit)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erro ao percorrer linhas: %v", err)
		return nil, err
	}

	return limits, nil
}

// Save insere um novo registro na tabela "limit"
func (r *LimitRepository) Save(limit domain.Limit) error {
	query := `INSERT INTO limit (nome, valor) VALUES ($1, $2) RETURNING id`

	err := r.db.QueryRow(query, limit.Nome, limit.Valor).Scan(&limit.ID)
	if err != nil {
		log.Printf("Erro ao inserir limite: %v", err)
		return err
	}

	return nil
}
