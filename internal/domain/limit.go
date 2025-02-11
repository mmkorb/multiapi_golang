package domain

import (
	"encoding/json"
	"fmt"
)

// Limit representa um registro na tabela "limit"
type Limit struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Valor int    `json:"valor"`
}

// ToString converte a struct Limit para uma string JSON
func (l Limit) ToString() string {
	limitJSON, err := json.Marshal(l)
	if err != nil {
		fmt.Println("Erro ao converter Limit para JSON:", err)
		return ""
	}
	return string(limitJSON)
}
