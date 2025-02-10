package rabbitmq_consumer

import (
	"log"
	config "multiapi_golang/configs"
	rabbitmq "multiapi_golang/pkg/rabbitmq"
)

func StartConsumer() {
	consumer, err := rabbitmq.GetConsumer(config.AppConfig.RabbitMQ.Queue)
	if err != nil {
		log.Fatalf("Erro ao obter consumidor RabbitMQ: %v", err)
	}

	log.Println("🐰 RabbitMQ Consumer iniciado... aguardando mensagens.")
	consumer.Start()
}
