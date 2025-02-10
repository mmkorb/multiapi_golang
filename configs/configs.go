package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Config armazena todas as configurações da aplicação
type Config struct {
	RabbitMQ               RabbitMQConfig  `json:"rabbitmq"`
	DB                     DBConfig        `json:"db"`
	Datastore              DatastoreConfig `json:"datastore"`
	PollingIntervalSeconds int             `json:"polling_interval_seconds"`
	EnableDatastore        bool            `json:"enable_datastore"` // Novo campo para controlar o uso do Datastore
}

// RabbitMQConfig contém as configurações de conexão do RabbitMQ
type RabbitMQConfig struct {
	URI   string `json:"uri"`
	Queue string `json:"queue"`
}

// DBConfig contém as configurações do banco de dados
type DBConfig struct {
	URI string `json:"uri"`
}

// DatastoreConfig contém as configurações do Google Datastore
type DatastoreConfig struct {
	ProjectID string `json:"project_id"`
}

// LoadConfig carrega as configurações a partir do arquivo config.json ou variáveis de ambiente
func LoadConfig() (*Config, error) {
	// Inicializa o viper
	viper.AddConfigPath(".")      // Caminho do arquivo de configuração
	viper.SetConfigName("config") // Nome do arquivo (config.json)
	viper.SetConfigType("json")   // Tipo de arquivo (json ao invés de yaml)
	viper.AutomaticEnv()          // Carrega variáveis de ambiente

	// Tenta ler o arquivo de configuração
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Erro ao ler o arquivo de configuração, usando variáveis de ambiente: %v", err)
	}

	// Mapeia as variáveis para a struct Config
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Erro ao mapear configurações: %v", err)
	}

	return &config, nil
}
