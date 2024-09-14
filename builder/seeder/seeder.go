package seeder

import (
	"context"
	"fmt"

	"builder/ent"
	"builder/ent/migrate"

	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

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
	// Seed the database

	// Create the races
	names := []string{
		"Dwarf",
		"Elf",
		"Gnome",
		"Halfling",
		"Human",
		// "Half-Elf",
		// "Half-Orc",
		"Tiefling",
	}
	for _, name := range names {
		if _, err := client.Race.Create().SetName(name).Save(context.Background()); err != nil {
			return fmt.Errorf("error creating race: %w", err)
		}
	}

	// Create the classes
	names = []string{
		"Barbarian",
		"Bard",
		"Cleric",
		"Druid",
		"Mage",
		"Monk",
		"Necromancer",
		"Paladin",
		"Ranger",
		"Rogue",
		"Sorcerer",
		"Warlock",
		"Warrior",
		"Wizard",
	}
	for _, name := range names {
		if _, err := client.Class.Create().SetName(name).Save(context.Background()); err != nil {
			return fmt.Errorf("error creating class: %w", err)
		}
	}

	return nil
}
