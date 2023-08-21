package main

import (
	"context"

	_ "github.com/ecshreve/dndgen/internal/log"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/oops"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	p := popper.NewPopper(ctx)
	start := p.Client.AbilityScore.Query().CountX(ctx)
	if err := p.PopulateAbilityScore(ctx); err != nil {
		log.Fatal(oops.Wrapf(err, "unable to populate AS entities"))
	}
	log.Infof("created %d entities for type As", p.Client.AbilityScore.Query().CountX(ctx)-start)

	start = p.Client.Skill.Query().CountX(ctx)
	if err := p.PopulateSkill(ctx); err != nil {
		log.Fatal(oops.Wrapf(err, "unable to populate Skill entities"))
	}
	log.Infof("created %d entities for type Skill", p.Client.Skill.Query().CountX(ctx)-start)

	// if err := p.PopulateAll(ctx); err != nil {
	// 	log.Error(err)
	// }
}
