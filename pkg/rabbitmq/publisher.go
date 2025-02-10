package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Publisher é responsável por enviar mensagens para o RabbitMQ
type Publisher struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// NewPublisher cria uma nova instância do Publisher
func NewPublisher(channel *amqp.Channel, queueName string) (*Publisher, error) {
	// Declara a fila (se não existir)
	queue, err := channel.QueueDeclare(
		queueName, // Nome da fila
		true,      // Durável
		false,     // Excluída quando o RabbitMQ for encerrado
		false,     // Exclusiva
		false,     // Sem auto-delete
		nil,       // Argumentos
	)
	if err != nil {
		log.Printf("Error declaring queue: %v", err)
		return nil, err
	}

	return &Publisher{
		Channel: channel,
		Queue:   queue,
	}, nil
}

// Publish envia uma mensagem para a fila do RabbitMQ
func (p *Publisher) Publish(message string) error {
	err := p.Channel.Publish(
		"",           // Exchange
		p.Queue.Name, // Fila
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Error publishing message: %v", err)
		return err
	}
	log.Printf("Message sent: %s", message)
	return nil
}
