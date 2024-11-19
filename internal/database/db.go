package database

import (
	"Authentication_System/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	err = DB.AutoMigrate(&models.RevokedToken{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database connected and migrated successfully!")
	return DB
}
