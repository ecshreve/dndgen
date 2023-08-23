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

// Popper is a populator for the ent database.
type Popper struct {
	Client   *ent.Client
	IdToIndx map[int]string
	IndxToId map[string]int
}

// NewPopper creates a new Popper.
func NewPopper(ctx context.Context, client *ent.Client) *Popper {
	if client == nil {
		log.Warn("client is nil, creating new client")
		client, _ = ent.Open(dialect.SQLite, "file:dev?mode=memory&_fk=1")
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatal(err)
		}
	}

	return &Popper{
		Client:   client,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
	}
}

// PopulateAll populates all entities from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	if err := p.PopulateAllGen(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate all gen")
	}

	start := len(p.IdToIndx)
	if err := p.PopulateEquipment(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate equipment")
	}
	log.Infof("created %d entities for type Equipment", len(p.IdToIndx)-start)

	start = len(p.IdToIndx)
	if err := p.PopulateProficiency(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate proficiency")
	}
	log.Infof("created %d entities for type Proficiency", len(p.IdToIndx)-start)

	log.Info("done populating all entities")
	return nil
}
