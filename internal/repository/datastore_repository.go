package repository

import (
	"context"
	"log"
	datastore "multiapi_golang/datastore"
	domain "multiapi_golang/domain"

	gcdatastore "cloud.google.com/go/datastore"
)

type RequestRepository struct{}

func NewRequestRepository() *RequestRepository {
	return &RequestRepository{}
}

func (r *RequestRepository) Save(request *domain.Request) error {
	client := datastore.GetDatastoreClient()
	ctx := context.Background()

	key := gcdatastore.IncompleteKey("Request", nil)
	_, err := client.Put(ctx, key, request)
	if err != nil {
		log.Printf("Error saving request to Datastore: %v", err)
		return err
	}

	return nil
}
