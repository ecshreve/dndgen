// Populates a sqlite database file with data.
//
// Usage:
//
//	go run cmd/popper/main.go dev.db
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/ecshreve/dndgen/ent/runtime"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetCallerFormatter(func(file string, line int, _ string) string {
		return fmt.Sprintf("%s:%d", strings.Replace(file, "/home/eric/repos/dndgen/", "", 1), line)
	})
	log.SetTimeFormat("15:04:05")
	styles := log.DefaultStyles()
	styles.Keys["ceq"] = lipgloss.NewStyle().Background(lipgloss.Color("204"))
	styles.Keys["weapon"] = lipgloss.NewStyle().Background(lipgloss.Color("120"))
	styles.Keys["armor"] = lipgloss.NewStyle().Background(lipgloss.Color("135"))
	styles.Keys["tool"] = lipgloss.NewStyle().Background(lipgloss.Color("184"))
	styles.Keys["gear"] = lipgloss.NewStyle().Background(lipgloss.Color("130"))
	styles.Keys["vehicle"] = lipgloss.NewStyle().Background(lipgloss.Color("166"))
	styles.Keys["choice"] = lipgloss.NewStyle().Background(lipgloss.Color("202"))
	styles.Keys["proficiency"] = lipgloss.NewStyle().Background(lipgloss.Color("204"))
	log.SetStyles(styles)
	log.Info("Starting dndgen/popper...")

	ctx := context.Background()

	dbname := "dev.db"
	// Get command-line arguments
	if len(os.Args) == 2 {
		dbname = os.Args[1]
	}

	log.Info("Using database", "dbname", dbname)

	// Check if the database is already populated
	if _, err := os.Stat(dbname); err == nil {
		log.Info("Database already populated skipping...")
		return
	}

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
	p := popper.NewPopper(ctx, client, "data")
	if err := p.PopulateAll(ctx); err != nil {
		log.Fatal(err)
	}
	if err := p.PopulateCustom(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("DB population completed.")
}
