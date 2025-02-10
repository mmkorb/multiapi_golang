package main

import (
	"log"
	"multiapi_golang/cmd/rabbitmq_consumer"
	"multiapi_golang/cmd/rest_api"
)

func main() {
	// Inicia a API REST em uma goroutine
	go rest_api.StartAPI()

	// Inicia o consumidor RabbitMQ em outra goroutine
	go rabbitmq_consumer.StartConsumer()

	log.Println("✅ Serviços iniciados! API REST e RabbitMQ Consumer estão rodando...")

	// Mantém o programa rodando
	select {}
}
