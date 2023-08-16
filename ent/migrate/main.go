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
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/migrate"
	"github.com/kr/pretty"
	_ "github.com/mattn/go-sqlite3"
)

func SeedAbilityScores(dir *atlas.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))

	wrk := migrate.NewClient()
	q := `{"query":"query RawData {\n  abilityScores {\n    index\n    name\n    full_name\n    desc\n    skills {\n      index\n      name\n      desc\n    }\n  }\n}","variables":{}}`

	type vv struct {
		Data struct {
			AbilityScores []struct {
				Indx     string   `json:"index"`
				Name     string   `json:"name"`
				FullName string   `json:"full_name"`
				Desc     []string `json:"desc"`
				// Skills   []struct {
				// 	Index string `json:"index"`
				// 	Name  string `json:"name"`
				// 	Desc  string `json:"desc"`
				// } `json:"skills"`
			} `json:"abilityScores"`
		} `json:"data"`
	}

	var v vv
	wrk.MakeGQLQuery(q, &v)
	pretty.Print(v)

	scores := make([]*ent.AbilityScoreCreate, len(v.Data.AbilityScores))
	for i, score := range v.Data.AbilityScores {
		scores[i] = client.AbilityScore.Create().
			SetIndx(score.Indx).
			SetName(score.Name).
			SetFullName(score.FullName).
			SetDesc(strings.Join(score.Desc, "\n"))
	}

	// The statement that generates the INSERT statement.
	err := client.AbilityScore.CreateBulk(
		scores...,
	).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_ability_scores",
		"Add the initial data to the database.",
	)
}

func main() {
	// We need a name for the new migration file.
	if len(os.Args) < 2 {
		log.Fatalln("no name given")
	}
	// Create a local migration directory.
	dir, err := atlas.NewLocalDir("./ent/migrate/migrations")
	if err != nil {
		log.Fatalln(err)
	}
	// Load the graph.
	graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}

	drv, err := sql.Open(dialect.SQLite, "sqlite://file?mode=memory&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := SeedAbilityScores(dir); err != nil {
		log.Fatalln(err)
	}

	// Inspect the current database state and compare it with the graph.
	m, err := schema.NewMigrate(drv, schema.WithDir(dir))
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.NamedDiff(context.Background(), os.Args[1], tbls...); err != nil {
		log.Fatalln(err)
	}
}
