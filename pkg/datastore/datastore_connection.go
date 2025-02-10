package datastore

import (
	"context"
	"log"

	config "multiapi_golang/configs"

	"cloud.google.com/go/datastore"
)

var client *datastore.Client
var ctx context.Context

func InitDatastore() {
	var err error
	ctx = context.Background()
	client, err = datastore.NewClient(ctx, config.AppConfig.Datastore.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func GetDatastoreClient() *datastore.Client {
	return client
}
