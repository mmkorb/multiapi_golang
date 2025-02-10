package repository

import (
	"log"
	"pkg/rabbitmq"
)

// RabbitMQRepository lida com a comunicação com o RabbitMQ
type RabbitMQRepository struct{}

// NewRabbitMQRepository cria uma nova instância do RabbitMQRepository
func NewRabbitMQRepository() *RabbitMQRepository {
	return &RabbitMQRepository{}
}

// SendMessage envia uma mensagem para o RabbitMQ
func (r *RabbitMQRepository) SendMessage(queueName, message string) error {
	// Usando a conexão singleton configurada no pkg/rabbitmq
	channel := rabbitmq.GetRabbitMQChannel()

	err := rabbitmq.SendMessage(channel, queueName, message)
	if err != nil {
		log.Printf("Error sending message to RabbitMQ: %v", err)
		return err
	}

	return nil
}
