package butlerai

import (
	"fmt"
	config2 "github.com/shrikanthcodes/butler-ai/backend/docs/config"
	"log"

	"gorm.io/gorm"
)

func init() {
	// Initialize DB connection
	str, err := config2.InitDB("Hoi")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println(str)

	var db *gorm.DB
	// Run migrations
	err = config2.RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully!")
}
