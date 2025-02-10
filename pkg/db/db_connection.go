package db

import (
	"database/sql"
	"fmt"
	"log"

	config "multiapi_golang/configs"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		config.AppConfig.DB.URI,
		"mydb",
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
}

func GetDB() *sql.DB {
	return DB
}
