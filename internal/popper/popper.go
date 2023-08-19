package popper

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

//go:generate go run gen.go

type Popper struct {
	Client   *ent.Client
	GenTypes []string
}

func NewPopper() *Popper {
	client, err := ent.Open(
		dialect.SQLite,
		// "file:dev?mode=memory&cache=shared&_fk=1",
		"file:file.db?_fk=1",
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}

	return &Popper{
		Client: client,
	}
}

func (p *Popper) PopAll() error {
	ctx := context.Background()

	start := p.Client.AbilityScore.Query().CountX(ctx)
	if err := p.PopulateAbilityScores(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate ability scores")
	}
	log.Infof("created %d entities for type AbilityScore", p.Client.AbilityScore.Query().CountX(ctx)-start)

	start = p.Client.Skill.Query().CountX(ctx)
	if err := p.PopulateSkills(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate skills")
	}
	log.Infof("created %d entities for type Skill", p.Client.Skill.Query().CountX(ctx)-start)

	start = p.Client.Alignment.Query().CountX(ctx)
	if err := p.PopulateAlignments(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate alignments")
	}
	log.Infof("created %d entities for type Alignment", p.Client.Alignment.Query().CountX(ctx)-start)

	return nil
}
