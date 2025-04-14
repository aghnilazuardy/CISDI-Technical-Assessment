package main

import (
	"cisdi-technical-assessment/REST/data-service/config"
	"cisdi-technical-assessment/REST/data-service/migrations"
	"cisdi-technical-assessment/REST/data-service/routes"
	"fmt"
	"log"
)

func main() {
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
	log.Printf("Data service starting on %s", addr)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
