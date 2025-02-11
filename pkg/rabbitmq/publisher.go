package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Publisher é responsável por enviar mensagens para o RabbitMQ
type Publisher struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func NewPublisher(rabbitURI, queueName string) (*Publisher, error) {
	// Conecta ao RabbitMQ
	conn, err := amqp.Dial(rabbitURI)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return nil, err
	}

	// Abre um canal
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Erro ao abrir canal no RabbitMQ: %v", err)
		conn.Close()
		return nil, err
	}

	// Declara a fila (caso não exista)
	queue, err := ch.QueueDeclare(
		queueName, // Nome da fila
		true,      // Durável
		false,     // Excluída ao encerrar RabbitMQ
		false,     // Exclusiva
		false,     // Sem auto-delete
		nil,       // Argumentos
	)
	if err != nil {
		log.Printf("Erro ao declarar a fila: %v", err)
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &Publisher{
		Connection: conn,
		Channel:    ch,
		Queue:      queue,
	}, nil
}

// Publish envia uma mensagem para a fila
func (p *Publisher) Publish(message string) error {
	err := p.Channel.Publish(
		"",           // Exchange (vazio para usar a fila diretamente)
		p.Queue.Name, // Routing Key (nome da fila)
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Erro ao publicar mensagem: %v", err)
	}
	return err
}

// Close fecha a conexão e o canal do RabbitMQ
func (p *Publisher) Close() {
	if err := p.Channel.Close(); err != nil {
		log.Printf("Erro ao fechar canal: %v", err)
	}
	if err := p.Connection.Close(); err != nil {
		log.Printf("Erro ao fechar conexão: %v", err)
	}
}
