//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/migrate"
	"github.com/ecshreve/dndgen/internal/popper"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
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
		schema.WithForeignKeys(true),
	}
	migrationName := ""
	if len(os.Args) != 2 {
		log.Warn("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>\ngenerating one for this migration plan'")
		migrationName = fmt.Sprintf("_GEN_%d", time.Now().Unix())
	} else {
		migrationName = os.Args[1]
	}

	if err := seed(ctx, dir, migrationName); err != nil {
		log.Fatalf("failed seeding database: %v", err)
	}

	err = migrate.NamedDiff(ctx, "sqlite://dev.db?_fk=1", migrationName, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}

func seed(ctx context.Context, dir *atlas.LocalDir, migrationName string) error {
	w := &schema.DirWriter{Dir: dir}

	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))

	p := &popper.Popper{
		Client: client,
	}

	if err := p.PopulateAll(ctx); err != nil {
		return oops.Wrapf(err, "failed populating database")
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		migrationName,
		"seed data to the database.",
	)
}
