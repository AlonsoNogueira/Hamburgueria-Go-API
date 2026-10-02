package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	database_url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(database_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro with conection database")
	}

	log.Println("Postgres Database conect")

	return db
}
