package initializers

import (
	"GIN-GORM-Tutorial/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Printf("DB: %v", DB)
	if DB == nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	log.Println("Database connection established")
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
