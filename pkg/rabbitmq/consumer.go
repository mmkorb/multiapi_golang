package rabbitmq

import (
	"log"
	config "multiapi_golang/configs"

	"github.com/streadway/amqp"
)

// Consumer representa um consumidor RabbitMQ
type Consumer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

// GetConsumer cria e retorna um Consumer configurado
func GetConsumer(queueName string) (*Consumer, error) {
	// Conecta ao RabbitMQ
	conn, err := amqp.Dial(config.AppConfig.RabbitMQ.URI)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return nil, err
	}

	// Abre um canal no RabbitMQ
	channel, err := conn.Channel()
	if err != nil {
		log.Printf("Erro ao abrir canal no RabbitMQ: %v", err)
		return nil, err
	}

	// Declara a fila (se não existir)
	queue, err := channel.QueueDeclare(
		queueName,
		true,  // Durável
		false, // Excluída quando o RabbitMQ for encerrado
		false, // Exclusiva
		false, // Sem auto-delete
		nil,   // Argumentos
	)
	if err != nil {
		log.Printf("Erro ao declarar fila: %v", err)
		return nil, err
	}

	return &Consumer{Channel: channel, Queue: queue}, nil
}

// Start inicia o consumo de mensagens da fila de forma bloqueante
func (c *Consumer) Start() {
	// Consome as mensagens de forma bloqueante (aguarda por novas mensagens)
	messages, err := c.Channel.Consume(
		c.Queue.Name,
		"",    // Consumer
		true,  // Auto ack
		false, // Exclusivo
		false, // Não compartilhar
		false, // Não auto-delete
		nil,   // Argumentos
	)
	if err != nil {
		log.Fatalf("Erro ao iniciar consumidor: %v", err)
	}

	log.Println("🐰 Aguardando novas mensagens...")

	// Loop que processa as mensagens recebidas
	for message := range messages {
		c.ProcessMessage(message.Body)
	}
}

// ProcessMessage processa cada mensagem recebida
func (c *Consumer) ProcessMessage(messageBody []byte) {
	log.Printf("Mensagem recebida: %s", messageBody)
	// Lógica de processamento da mensagem
}
