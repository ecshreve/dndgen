//go:build ignore

package main

import (
	"builder/ent/migrate"
	"builder/ent/migrate/migratedata"
	"context"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.SQLite),          // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithGlobalUniqueID(true),
	}
	if len(os.Args) != 2 {
		log.Fatalf("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	if os.Args[1] == "all" {
		for _, v := range migratedata.SeedersToRun {
			err = v(ctx, dir)
			if err != nil {
				log.Fatal("failed running data migration", "error", err)
			}
		}
		log.Info("data migrations complete")
	}

	err = migrate.NamedDiff(ctx, "sqlite://ent?mode=memory&_fk=1", os.Args[1], opts...)
	if err != nil {
		log.Fatal("failed generating migration file", "error", err)
	}

	log.Info("migration file generated")

}
