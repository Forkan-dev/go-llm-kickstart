package database

import (
	"learning-companion/internal/model/admin"
	"learning-companion/internal/model/auth"
	"learning-companion/internal/model/quiz"
	"learning-companion/internal/model/quizattempt"
	"learning-companion/internal/model/user"
	"log"
)

func Migrate() {
	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}

	if err := DB.AutoMigrate(&auth.RefreshToken{}); err != nil {
		log.Fatalf("could not migrate refresh token model: %v", err)
	}

	if err := DB.AutoMigrate(&quiz.Subject{}); err != nil {
		log.Fatalf("could not migrate subject model: %v", err)
	}

	if err := DB.AutoMigrate(&quiz.Topic{}); err != nil {
		log.Fatalf("could not migrate Topic model: %v", err)
	}

	if err := DB.AutoMigrate(&quiz.Quiz{}); err != nil {
		log.Fatalf("could not migrate Quiz model: %v", err)
	}

	if err := DB.AutoMigrate(&quiz.Question{}); err != nil {
		log.Fatalf("could not migrate Question model: %v", err)
	}

	if err := DB.AutoMigrate(&quizattempt.QuizAttempt{}); err != nil {
		log.Fatalf("could not migrate quiz attempt model: %v", err)
	}

	if err := DB.AutoMigrate(&quizattempt.QuizAttemptAnswer{}); err != nil {
		log.Fatalf("could not migrate quiz attempt answer model: %v", err)
	}

	if err := DB.AutoMigrate(&admin.Admin{}); err != nil {
		log.Fatalf("could not migrate admin model: %v", err)
	}

	log.Println("Database migration successful")
}
