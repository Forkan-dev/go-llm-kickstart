package main

import (
	"learning-companion/internal/api/validator"
	"learning-companion/internal/config"
	"learning-companion/internal/server"
	"learning-companion/pkg/database"
	"log"

	"github.com/gin-gonic/gin/binding"
	validatorV10 "github.com/go-playground/validator/v10"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Connect to the database
	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Register custom validators
	if v, ok := binding.Validator.Engine().(*validatorV10.Validate); ok {
		v.RegisterValidation("password", validator.NewPasswordValidator(&cfg.Validation.Password))
	}

	// Run migrations
	database.Migrate()

	// Seed the database
	database.Seed()

	// Start the server
	server.StartServer(cfg)
}
