package config

import (
	"log"

	models "backend/pkg/models"

	"gorm.io/gorm"
)

// RunMigrations runs the necessary database migrations
func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}) // Migrate your models
	if err != nil {
		return err
	}
	log.Println("Database migrations ran successfully!")
	return nil
}
