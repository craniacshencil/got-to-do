package handlers

import (
	"database/sql"
	"log"
	"os"

	"github.com/craniacshencil/got_to_do/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type ApiCfg struct {
	DB *database.Queries
}

func SetupDB() *ApiCfg {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("DB_URL was not found in .env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println("Couldn't connect to databse", err)
	}

	ApiConfig := &ApiCfg{
		DB: database.New(db),
	}

	return ApiConfig
}
