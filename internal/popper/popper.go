package popper

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

//go:generate go run gen.go

// Popper is a populator for the ent database.
type Popper struct {
	Client     *ent.Client
	Reader     *ent.Client
	IdToIndx   map[int]string
	IndxToId   map[string]int
	Context    *context.Context
	PathPrefix string
}

// NewPopper creates a new Popper.
func NewPopper(ctx context.Context, client *ent.Client) *Popper {
	if client == nil {
		client, _ = ent.Open(dialect.SQLite, "file:dev.db?_fk=1")
	}
	reader, _ := ent.Open(dialect.SQLite, "file:dev.db?_fk=1")

	return &Popper{
		Client:   client,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
		Context:  &ctx,
		Reader:   reader,
	}
}

// NewTestPopper creates a new Popper for testing.
func NewTestPopper(ctx context.Context) *Popper {
	client, err := ent.Open("sqlite3", "file:testdb?mode=memory&_fk=1")
	if err != nil {
		log.Error(err)
		return nil
	}
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	return &Popper{
		Client:     client,
		Reader:     client,
		IdToIndx:   map[int]string{},
		IndxToId:   map[string]int{},
		PathPrefix: "dndgen/data/",
	}
}

// PopAll populates all entities generated from the JSON data files.
func (p *Popper) PopAll(ctx context.Context) error {
	createsAbilityScore, err := p.PopulateAbilityScore(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}
	p.Client.AbilityScore.CreateBulk(createsAbilityScore...).ExecX(ctx)
	log.Infof("created %d entities for type AbilityScore", len(createsAbilityScore))

	createsSkill, err := p.PopulateSkill(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}
	p.Client.Skill.CreateBulk(createsSkill...).ExecX(ctx)
	log.Infof("created %d entities for type Skill", len(createsSkill))

	createsLanguage, err := p.PopulateLanguage(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Language entities")
	}
	p.Client.Language.CreateBulk(createsLanguage...).ExecX(ctx)
	log.Infof("created %d entities for type Language", len(createsLanguage))

	createsClass, err := p.PopulateClass(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}
	p.Client.Class.CreateBulk(createsClass...).ExecX(ctx)
	log.Infof("created %d entities for type Class", len(createsClass))

	createsRace, err := p.PopulateRace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}
	p.Client.Race.CreateBulk(createsRace...).ExecX(ctx)
	log.Infof("created %d entities for type Race", len(createsRace))

	return nil
}
