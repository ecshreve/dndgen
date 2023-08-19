package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScores(ctx context.Context) error {
	fpath := "internal/popper/data/5e-SRD-AbilityScores.json"
	var v []ent.AbilityScore

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.AbilityScore.Create().SetAbilityScore(&vv).Save(ctx)
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

// PopulateAlignment populates the Alignment entities from the JSON data files.
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

// PopulateSkill populates the Skill entities from the JSON data files.
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
