package popper

import (
	"context"

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

func (p *Popper) PopAll() error {
	ctx := context.Background()

	start := p.Client.AbilityScore.Query().CountX(ctx)
	if err := p.PopulateAbilityScores(); err != nil {
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

func (p *Popper) PopulateAbilityScores() error {
	fpath := "internal/popper/data/5e-SRD-Ability-Scores.json"
	var v []ent.AbilityScore

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, as := range v {
		_, err := p.Client.AbilityScore.Create().SetAbilityScore(&as).Save(context.Background())
		if ent.IsConstraintError(err) {
			log.Warnf("constraint failed, skipping %s", as.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create ability score %s", as.Indx)
		}
	}

	return nil
}

func (p *Popper) PopulateSkills(ctx context.Context) error {
	fpath := "internal/popper/data/5e-SRD-Skills.json"
	var v []ent.Skill

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Skill.Create().SetSkill(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Warnf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}

func (p *Popper) PopulateAlignments(ctx context.Context) error {
	fpath := "internal/popper/data/5e-SRD-Alignments.json"
	var v []ent.Alignment

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Alignment.Create().SetAlignment(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Warnf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}
