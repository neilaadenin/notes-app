package database

import (
	"log"
	"os"
	"strings"
	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load .env file - coba dari root dan dari backend folder
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not loaded, using system environment variables")
		}
	}

	// Ambil DATABASE_URL dari .env
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is missing in .env")
	}

	// Log database yang digunakan (tanpa password)
	logDSN := dsn
	if strings.Contains(logDSN, "@") {
		parts := strings.Split(logDSN, "@")
		if len(parts) > 0 {
			// Hide password in log
			userPart := strings.Split(parts[0], ":")
			if len(userPart) > 2 {
				logDSN = strings.Join([]string{userPart[0], "***", parts[1]}, "@")
			}
		}
	}
	log.Printf("Connecting to database: %s", logDSN)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate tabel
	err = db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Connected to database successfully!")
}
