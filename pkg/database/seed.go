package database

import (
	"learning-companion/internal/model"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// Check if users already exist
	var count int64
	DB.Model(&model.User{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	user := model.User{
		Uuid:      uuid.New().String(),
		Username:  "admin",
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
		Password:  string(hashedPassword),
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Fatalf("could not seed database: %v", err)
	}

	log.Println("Database seeding successful")
}
