package boostrap

import (
	"log"

	"github.com/joho/godotenv"
)

func Run() error {
	// Load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load database
	db := newDatabase()

	return nil
}
