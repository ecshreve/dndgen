package popper

import (
	"context"
	"errors"

	"entgo.io/ent/dialect"
	"github.com/ecshreve/dndgen/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Popper struct {
	Client *ent.Client
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

func (p *Popper) PopulateAbilityScores() error {
	fpath := "internal/popper/data/5e-SRD-Ability-Scores.json"
	var v []ent.AbilityScore

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	errCount := 0
	for _, as := range v {
		_, err := p.Client.AbilityScore.Create().SetAbilityScore(&as).Save(context.Background())
		if err != nil {
			if ent.IsConstraintError(err) {
				log.Warnf("constraint failed, skipping %s", as.Indx)
				errCount++
				continue
			}
			return oops.Wrapf(err, "unable to create ability score %+v", as.Indx)
		}
	}

	if errCount > 0 {
		return errors.New("failed to populate one or more entries - check logs")
	}

	return nil
}
