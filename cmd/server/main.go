package main

import (
	"learning-companion/internal/config"
	"learning-companion/internal/server"
	"learning-companion/pkg/database"
	"log"
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

	// Run migrations
	database.Migrate()

	// Seed the database
	database.Seed()

	// Start the server
	server.StartServer(cfg)
}
