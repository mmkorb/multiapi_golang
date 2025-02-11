package db

import (
	"database/sql"
	"log"

	config "multiapi_golang/configs"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := config.AppConfig.DB.URI
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
}

func GetDB() *sql.DB {
	return DB
}
