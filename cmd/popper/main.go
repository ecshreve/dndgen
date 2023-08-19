package main

import (
	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"

	log "github.com/sirupsen/logrus"
)

func main() {
	p := popper.NewPopper()

	if err := p.PopAll(); err != nil {
		log.Error(err)
	}
}
