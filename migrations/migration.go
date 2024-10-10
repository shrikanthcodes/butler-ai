package migrations

import (
	model "github.com/shrikanthcodes/butler-ai/internal/entity"
	"log"

	"gorm.io/gorm"
)

// RunMigrations - function to run migrations in order
func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatalf("Failed to run users table migration: %v", err)
	}

	err = db.AutoMigrate(
		&model.Authentication{},
		&model.Writeup{},
		&model.Profile{},
		&model.Health{},
		&model.Diet{},
		&model.Inventory{},
		&model.Shopping{},
		&model.MealChoice{},
		&model.Goal{},
		&model.Conversation{},
		&model.Recipe{},
	)
	if err != nil {
		log.Fatalf("Failed to run other tables migration: %v", err)
	}

	return err
}
