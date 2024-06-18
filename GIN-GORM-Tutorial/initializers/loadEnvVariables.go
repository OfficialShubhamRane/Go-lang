package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Environment variables loaded:")
	log.Printf("DB_URL: %s", os.Getenv("DB_URL"))
}
