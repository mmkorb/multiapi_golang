package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Postgres driver
)

// DB é a variável global que mantém a conexão com o banco de dados
var DB *sql.DB

// InitDB inicializa a conexão com o banco de dados e deve ser chamada apenas uma vez no início
func InitDB() {
	var err error
	// Exemplo de string de conexão, substitua pelos seus dados reais
	connStr := "user=username dbname=mydb sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
}

// GetDB retorna a instância única da conexão com o banco de dados
func GetDB() *sql.DB {
	return DB
}
