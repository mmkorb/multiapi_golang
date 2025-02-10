package repository

import (
	"log"
	db "multiapi_golang/db"
	domain "multiapi_golang/domain"
)

type LimitRepository struct{}

func NewLimitRepository() *LimitRepository {
	return &LimitRepository{}
}

func (r *LimitRepository) GetAll() ([]*domain.Limit, error) {
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
