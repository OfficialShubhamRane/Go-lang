package main

import (
	"GIN-GORM-Tutorial/initializers"
	"GIN-GORM-Tutorial/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is not established")
	}
	// Migrate the schema
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
