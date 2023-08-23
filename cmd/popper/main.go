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

	cl, _ := ent.Open("sqlite3", "file:ent/migrate/file.db?_fk=1")
	p := popper.NewPopper(ctx, cl)

	if err := p.CleanUp(ctx); err != nil {
		log.Error(err)
	}

	if err := p.PopulateAll(ctx); err != nil {
		log.Error(err)
	}
}
