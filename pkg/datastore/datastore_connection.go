package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

var client *datastore.Client
var ctx context.Context

// InitDatastore inicializa a conexão com o Google Datastore
func InitDatastore(projectID string) {
	var err error
	ctx = context.Background()
	client, err = datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// GetDatastoreClient retorna a instância única do cliente do Google Datastore
func GetDatastoreClient() *datastore.Client {
	return client
}
