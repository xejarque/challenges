package main

import (
	"github.com/andreibthoughtworks/challenges/solution/internal/app"
	"github.com/andreibthoughtworks/challenges/solution/internal/infrastructure"
	"log"
)

func main() {
	// Create a new instance of the PostgresDB adapter
	dbAdapter, err := infrastructure.NewSqlDB(infrastructure.LocalDatabaseConfig())
	if err != nil {
		log.Fatalf("failed to create database adapter: %v", err)
	}

	// Create a new instance of the application
	app := app.NewApplication(dbAdapter)

	// Start the application
	if err := app.Run(); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}
