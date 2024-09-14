package seeder

import (
	"context"
	"fmt"

	"builder/ent"
	"builder/ent/migrate"

	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

//go:generate go run gen.go

func NewClient(dbUrl string) (*ent.Client, error) {
	ctx := context.Background()
	log.Info("Creating client")

	client, err := ent.Open(
		"sqlite3",
		dbUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		return nil, fmt.Errorf("error creating schema: %w", err)
	}

	if err := SeedAll(ctx, client); err != nil {
		return nil, fmt.Errorf("error seeding data: %w", err)
	}

	return client, nil
}
