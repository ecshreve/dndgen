package main

import (
	"builder/seeder"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting builder")

	client, err := seeder.NewClient("file:dev.db?_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("done", "client", client)
}
