package main

import (
	"context"

	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	p := popper.NewPopper(ctx)
	if err := p.PopulateAll(ctx); err != nil {
		log.Error(err)
	}
}
