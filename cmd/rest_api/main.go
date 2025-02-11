package rest_api

import (
	"log"
	"net/http"

	_ "multiapi_golang/cmd/rest_api/docs"
	config "multiapi_golang/configs"
	"multiapi_golang/internal/domain"
	"multiapi_golang/internal/repository"
	"multiapi_golang/internal/service"
	"multiapi_golang/pkg/datastore"
	"multiapi_golang/pkg/db"
	"multiapi_golang/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MultiAPI Golang
// @version 1.0
// @description API para gerenciamento de limites e integração com RabbitMQ
// @host localhost:8080
// @BasePath /
func StartAPI() {
	// Inicializa o banco de dados PostgreSQL
	db.InitDB()

	// Inicializa o Google Datastore se estiver ativado
	if config.AppConfig.EnableDatastore {
		datastore.InitDatastore()
	} else {
		log.Println("Google Datastore está desativado.")
	}

	// Inicializa o RabbitMQ Producer
	producer, err := rabbitmq.NewPublisher(config.AppConfig.RabbitMQ.URI, config.AppConfig.RabbitMQ.Queue)
	if err != nil {
		log.Fatalf("Erro ao inicializar produtor RabbitMQ: %v", err)
	}

	// Inicializa os serviços e repositórios
	limitRepo := repository.NewLimitRepository()
	limitService := service.NewLimitService(limitRepo, producer)

	// Criando o router Gin
	router := gin.Default()

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Healthcheck
	router.GET("/health", HealthCheck)

	// Endpoint para buscar todos os registros
	router.GET("/limits", func(c *gin.Context) {
		GetAllLimits(c, limitService)
	})

	// Endpoint para criar um novo registro e publicar no RabbitMQ
	router.POST("/limits", func(c *gin.Context) {
		CreateLimit(c, limitService)
	})

	// Inicia o servidor Gin na porta 8080
	log.Println("🚀 API REST iniciada na porta 8080")
	router.Run(":8080")
}

// HealthCheck verifica se a API está rodando
// @Summary Verifica a saúde da API
// @Description Retorna um status simples indicando que a API está rodando
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// GetAllLimits retorna todos os registros da tabela "limit"
// @Summary Retorna todos os limites cadastrados
// @Description Obtém todos os registros da tabela "limit"
// @Tags limits
// @Produce json
// @Success 200 {array} domain.Limit
// @Failure 500 {object} map[string]string
// @Router /limits [get]
func GetAllLimits(c *gin.Context, limitService *service.LimitService) {
	limits, err := limitService.GetAllLimits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar limites"})
		return
	}
	c.JSON(http.StatusOK, limits)
}

// CreateLimit cria um novo registro e publica no RabbitMQ
// @Summary Cria um novo limite
// @Description Insere um novo limite no banco de dados e publica no RabbitMQ
// @Tags limits
// @Accept json
// @Produce json
// @Param limit body domain.Limit true "Dados do limite"
// @Success 201 {object} domain.Limit
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /limits [post]
func CreateLimit(c *gin.Context, limitService *service.LimitService) {
	var newLimit domain.Limit
	if err := c.ShouldBindJSON(&newLimit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := limitService.CreateLimit(newLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar limite"})
		return
	}

	c.JSON(http.StatusCreated, newLimit)
}
