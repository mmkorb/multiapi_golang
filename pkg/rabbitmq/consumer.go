package rabbitmq

import (
	"configs"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Consumer é responsável por consumir mensagens do RabbitMQ
type Consumer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// NewConsumer cria uma nova instância do Consumer
func NewConsumer(channel *amqp.Channel, queueName string) (*Consumer, error) {
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

	return &Consumer{
		Channel: channel,
		Queue:   queue,
	}, nil
}

// Start começa a consumir mensagens da fila a cada intervalo configurado
func (c *Consumer) Start() {
	// Intervalo de polling configurado (em segundos)
	pollingInterval := time.Duration(configs.AppConfig.PollingIntervalSeconds) * time.Second

	// Inicia um loop infinito para buscar mensagens de acordo com o intervalo
	for {
		messages, err := c.Channel.Consume(
			c.Queue.Name, // Fila
			"",           // Consumer
			true,         // Auto ack
			false,        // Exclusivo
			false,        // Não compartilhar
			false,        // Não auto-delete
			nil,          // Argumentos
		)
		if err != nil {
			log.Fatalf("Error starting consumer: %v", err)
		}

		// Aguarda por mensagens por um intervalo configurado
		select {
		case message := <-messages:
			c.ProcessMessage(message.Body)
		case <-time.After(pollingInterval):
			log.Printf("No messages in %v, retrying...", pollingInterval)
		}
	}
}

// ProcessMessage processa cada mensagem recebida
func (c *Consumer) ProcessMessage(messageBody []byte) {
	log.Printf("Message received: %s", messageBody)
	// Lógica de processamento da mensagem
	// Aqui você pode chamar outro serviço ou executar algum trabalho específico
}
