//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/migrate"
	"github.com/kr/pretty"
	_ "github.com/mattn/go-sqlite3"
)

func Seed(dir *atlas.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))

	// reader, _ := ent.Open("sqlite3", "file:file.db?cache=shared&_fk=1")
	wrk := migrate.NewClient()
	q := `{"query":"query Data {\n  damageTypes {\n    index\n    name\n    desc\n  }\n}","variables":{}}`

	type wrapper struct {
		Data struct {
			Resources []struct {
				Indx string   `json:"index"`
				Name string   `json:"name"`
				Desc []string `json:"desc"`
			} `json:"damageTypes"`
		} `json:"data"`
	}
	var v wrapper
	wrk.MakeGQLQuery(q, &v)
	pretty.Print(v)

	cc := make([]*ent.DamageTypeCreate, len(v.Data.Resources))
	for i, dd := range v.Data.Resources {
		cc[i] = client.DamageType.Create().
			SetIndx(dd.Indx).
			SetName(dd.Name).
			SetDesc(strings.Join(dd.Desc, "\n"))
	}

	// The statement that generates the INSERT statement.
	err := client.DamageType.CreateBulk(
		cc...,
	).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_damage_types",
		"Add the initial data to the database.",
	)
}

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.SQLite),          // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	if err := Seed(dir); err != nil {
		log.Fatalf("failed seeding: %v", err)
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, "sqlite://file?mode=memory&_fk=1", os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
