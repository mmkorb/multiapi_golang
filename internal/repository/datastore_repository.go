package repository

import (
	"context"
	"internal/domain"
	"log"
	"pkg/datastore"

	"cloud.google.com/go/datastore"
)

// RequestRepository lida com a persistência de solicitações no Google Datastore
type RequestRepository struct{}

// NewRequestRepository cria uma nova instância do RequestRepository
func NewRequestRepository() *RequestRepository {
	return &RequestRepository{}
}

// Save salva uma solicitação no Google Datastore
func (r *RequestRepository) Save(request *domain.Request) error {
	// Usando a conexão singleton configurada no pkg/datastore
	client := datastore.GetDatastoreClient()
	ctx := context.Background()

	key := datastore.IncompleteKey("Request", nil)
	_, err := client.Put(ctx, key, request)
	if err != nil {
		log.Printf("Error saving request to Datastore: %v", err)
		return err
	}

	return nil
}
