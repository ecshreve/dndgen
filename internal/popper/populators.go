package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScore(ctx context.Context) error {
	fpath := "internal/popper/data/AbilityScore.json"
	var v []ent.AbilityScore

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.AbilityScore.Create().SetAbilityScore(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.ID)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.ID)
		}
	}

	return nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) error {
	fpath := "internal/popper/data/Skill.json"
	var v []Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, wrapper := range v {
		ss := ent.Skill{
			ID:             wrapper.ID,
			Name:           wrapper.Name,
			Desc:           wrapper.Desc,
			AbilityScoreID: wrapper.AbilityScore.ID,
		}
		_, err := p.Client.Skill.Create().SetSkill(&ss).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", ss.ID)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", ss.ID)
		}
	}

	return nil
}

// PopulateAll populates all entities from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	var start int

	start = p.Client.AbilityScore.Query().CountX(ctx)
	if err := p.PopulateAbilityScore(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}
	log.Infof("created %d entities for type AbilityScore", p.Client.AbilityScore.Query().CountX(ctx)-start)

	start = p.Client.Skill.Query().CountX(ctx)
	if err := p.PopulateSkill(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}
	log.Infof("created %d entities for type Skill", p.Client.Skill.Query().CountX(ctx)-start)

	return nil
}
