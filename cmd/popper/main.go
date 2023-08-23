package main

import (
	"context"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	if err := seed(ctx); err != nil {
		log.Fatalf("failed seeding database: %v", err)
	}
}

func seed(ctx context.Context) error {
	dir, _ := atlas.NewLocalDir("ent/migrate/migrations")

	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.SQLite, w)))
	// if err := client.Schema.Create(ctx); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	p := popper.NewPopper(ctx, client)

	p.PopulateAbilityScore(ctx)
	// w.FlushChange("ability_score", "seed data to the database.")

	p.PopulateSkill(ctx)
	// w.FlushChange("skill", "seed data to the database.")

	p.PopulateClass(ctx)
	// w.FlushChange("class", "seed data to the database.")

	p.PopulateRace(ctx)
	// w.FlushChange("race", "seed data to the database.")

	p.PopulateLanguage(ctx)
	// w.FlushChange("language", "seed data to the database.")

	// p.PopulateEquipment(ctx)
	// // w.FlushChange("equipment", "seed data to the database.")

	// p.PopulateProficiency(ctx)
	// w.FlushChange("proficiency", "seed data to the database.")
	w.FlushChange("seed", "seed data to the database.")

	return nil
}
