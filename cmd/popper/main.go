package main

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file://dev.db?_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	p := popper.NewPopper(ctx, client)
	if err := p.CleanUp(ctx); err != nil {
		log.Error(err)
	}

	if err := p.PopulateAllGen(ctx); err != nil {
		log.Error(err)
	}

	if err := p.PopulateEquipment(ctx); err != nil {
		log.Error(err)
	}
	if err := p.PopulateProficiency(ctx); err != nil {
		log.Error(err)
	}
}
