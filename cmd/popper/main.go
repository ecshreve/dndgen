//go:build ignore
// +build ignore

package main

import (
	"context"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.Info("Starting dndgen/popper...")

	ctx := context.Background()

	// Get command-line arguments
	dbname := os.Args[1]
	if dbname == "" {
		dbname = "dev_default.db"
	}

	log.Info("Using database", "dbname", dbname)
	
	dburl := "file:" + dbname + "?_fk=1"
	client, err := ent.Open(dialect.SQLite, dburl)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Creating schema...")
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	log.Info("Populating database...")
	p := popper.NewPopper(ctx, client)
	if err := p.PopulateAll(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("DB population completed.")
}
