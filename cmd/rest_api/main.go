package rest_api

import (
	"log"
	config "multiapi_golang/configs"
	"multiapi_golang/pkg/datastore"
	"multiapi_golang/pkg/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartAPI() {
	// Inicializa o banco de dados PostgreSQL
	db.InitDB()

	// Inicializa o Google Datastore se estiver ativado
	if config.AppConfig.EnableDatastore {
		datastore.InitDatastore()
	} else {
		log.Println("Google Datastore está desativado.")
	}

	// Criando o router Gin
	router := gin.Default()

	// Rota do Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rota simples de teste
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API REST rodando!"})
	})

	// Inicia o servidor Gin na porta 8080
	log.Println("🚀 API REST iniciada na porta 8080")
	router.Run(":8080")
}
