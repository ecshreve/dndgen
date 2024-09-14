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
	log.Info("Creating client")

	client, err := ent.Open(
		"sqlite3",
		dbUrl,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		return nil, fmt.Errorf("error creating schema: %w", err)
	}

	if err := Seed(client); err != nil {
		return nil, fmt.Errorf("error seeding database: %w", err)
	}

	return client, nil
}
func Seed(client *ent.Client) error {
	ctx := context.Background()
	// Seed the database
	_, err := SeedAbilityScore(ctx, client)
	if err != nil {
		return fmt.Errorf("error seeding ability scores: %w", err)
	}
	_, err = SeedAlignment(ctx, client)
	if err != nil {
		return fmt.Errorf("error seeding alignments: %w", err)
	}
	_, err = SeedClass(ctx, client)
	if err != nil {
		return fmt.Errorf("error seeding classes: %w", err)
	}
	_, err = SeedRace(ctx, client)
	if err != nil {
		return fmt.Errorf("error seeding races: %w", err)
	}

	return nil
}
