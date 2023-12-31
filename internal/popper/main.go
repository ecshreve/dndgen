//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	client, err := ent.Open(dialect.SQLite, "file:dev.db?_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	p := popper.NewPopper(ctx, client)
	// _, err = p.PopulateCoin(ctx)
	// _, err = p.PopulateAbilityScore(ctx)
	// _, err = p.PopulateSkill(ctx)
	// _, err = p.PopulateLanguage(ctx)
	// _, err = p.PopulateDamageType(ctx)
	// _, err = p.PopulateWeaponProperty(ctx)
	// err = p.PopulateEquipment(ctx)
	// _, err = p.PopulateClass(ctx)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := p.PopulateAll(ctx); err != nil {
		log.Fatal(err)
	}
}
