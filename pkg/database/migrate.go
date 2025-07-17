package database

import (
	"learning-companion/internal/model"
	"log"
)

func Migrate() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}
	
	if err := DB.AutoMigrate(&model.RefreshToken{}); err != nil {
		log.Fatalf("could not migrate refresh token model: %v", err)
	}
	log.Println("Database migration successful")
}
