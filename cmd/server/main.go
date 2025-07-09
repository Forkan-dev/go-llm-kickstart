package main

import (
	"learning-companion/internal/config"
	"learning-companion/internal/server"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Start the server
	server.StartServer(cfg)
}
