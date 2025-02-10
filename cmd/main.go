package main

import (
	"configs"
	"internal/repository"
	"internal/service"
	"log"
	"pkg/datastore"
	"pkg/db"
	"pkg/rabbitmq"
)

func main() {
	// Carregar a configuração
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Inicializa as conexões
	db.InitDB() // Inicializa o banco de dados
	if config.EnableDatastore {
		// Inicializa o Google Datastore se estiver habilitado
		datastore.InitDatastore(config.Datastore.ProjectID)
	} else {
		log.Println("Google Datastore is disabled.")
	}

	err = rabbitmq.InitRabbitMQ(config.RabbitMQ.URI) // Inicializa o RabbitMQ
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	// Cria os repositórios
	limitRepo := repository.NewLimitRepository()
	requestRepo := repository.NewRequestRepository()
	rabbitMQRepo := repository.NewRabbitMQRepository()

	// Cria os serviços
	limitService := service.NewLimitService(limitRepo)
	requestService := service.NewRequestService(requestRepo)

	// Exemplo de uso do Publisher para enviar mensagens
	channel := rabbitmq.GetRabbitMQChannel() // Obtém o canal do RabbitMQ
	publisher, err := rabbitmq.NewPublisher(channel, config.RabbitMQ.Queue)
	if err != nil {
		log.Fatalf("Error creating publisher: %v", err)
	}
	message := "New request added"
	err = publisher.Publish(message)
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	// Exemplo de uso do Consumer para consumir mensagens com intervalo configurado
	consumer, err := rabbitmq.NewConsumer(channel, config.RabbitMQ.Queue)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}

	go consumer.Start() // Inicia o consumidor em uma goroutine para consumir as mensagens

	// A aplicação continuará rodando, processando as mensagens
	// Se necessário, adicione lógica adicional para finalizar ou monitorar a execução
	select {} // O programa ficará aqui aguardando por novas mensagens
}
