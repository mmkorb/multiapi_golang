package config

// Config armazena todas as configurações da aplicação
type Config struct {
	RabbitMQ               RabbitMQConfig
	DB                     DBConfig
	Datastore              DatastoreConfig
	PollingIntervalSeconds int
	EnableDatastore        bool
}

// RabbitMQConfig contém as configurações de conexão do RabbitMQ
type RabbitMQConfig struct {
	URI   string
	Queue string
}

// DBConfig contém as configurações do banco de dados
type DBConfig struct {
	URI string
}

// DatastoreConfig contém as configurações do Google Datastore
type DatastoreConfig struct {
	ProjectID string
}

// AppConfig contém a configuração global da aplicação
var AppConfig = &Config{
	RabbitMQ: RabbitMQConfig{
		URI:   "amqp://guest:guest@localhost:5672/",
		Queue: "requestQueue",
	},
	DB: DBConfig{
		URI: "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable",
	},
	Datastore: DatastoreConfig{
		ProjectID: "your-google-project-id",
	},
	EnableDatastore: false,
}
