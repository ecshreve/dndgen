package main

import (
	"builder/seeder"
	"context"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting builder")

	client, err := seeder.NewClient("dev.db", "_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	// Query the database
	classes, err := client.Class.Query().All(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Classes", "classes", classes)
}
