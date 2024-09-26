package main

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetCallerFormatter(func(file string, line int, _ string) string {
		return fmt.Sprintf("%s:%d", strings.Replace(file, "/home/eric/repos/dndgen/", "", 1), line)
	})
	log.SetPrefix("🛠️")
	log.SetTimeFormat("15:04:05")
	log.Info("Starting dndgen/creator...")

	ctx := context.Background()

	dbname := "dev.db"
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

	CreateCharacter(ctx, client)
	log.Info("Character created")
}

// CreateCharacter creates a character
func CreateCharacter(ctx context.Context, client *ent.Client) {
	character := client.Character.Create().
		SetName("Zeek").
		SetRace(client.Race.Query().
			Where(race.Name("Human")).
			FirstX(ctx)).
		SetClass(client.Class.Query().
			Where(class.Name("Rogue")).
			FirstX(ctx)).
		SaveX(ctx)

	log.Info("Character created", "character", character)
}
