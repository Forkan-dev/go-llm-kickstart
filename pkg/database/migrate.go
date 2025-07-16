package database

import (
	"learning-companion/internal/model"
	"log"
)

func Migrate() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}
	log.Println("Database migration successful")
}
