package main

import (
	"cisdi-technical-assessment/REST/auth-service/config"
	"cisdi-technical-assessment/REST/auth-service/migrations"
	"cisdi-technical-assessment/REST/auth-service/routes"
	"fmt"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Run database migrations
	err = migrations.RunMigrations(cfg)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Setup and run the router
	r := routes.SetupRouter(cfg)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Auth service starting on %s", addr)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
