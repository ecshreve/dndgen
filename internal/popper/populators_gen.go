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
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}

// PopulateAlignment populates the Alignment entities from the JSON data files.
func (p *Popper) PopulateAlignment(ctx context.Context) error {
	fpath := "internal/popper/data/Alignment.json"
	var v []ent.Alignment

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Alignment.Create().SetAlignment(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) error {
	fpath := "internal/popper/data/Skill.json"
	var v []ent.Skill

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Skill.Create().SetSkill(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}

// PopulateProficiency populates the Proficiency entities from the JSON data files.
func (p *Popper) PopulateProficiency(ctx context.Context) error {
	fpath := "internal/popper/data/Proficiency.json"
	var v []ent.Proficiency

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Proficiency.Create().SetProficiency(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}

// PopulateMagicSchool populates the MagicSchool entities from the JSON data files.
func (p *Popper) PopulateMagicSchool(ctx context.Context) error {
	fpath := "internal/popper/data/MagicSchool.json"
	var v []ent.MagicSchool

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.MagicSchool.Create().SetMagicSchool(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
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
	
	start = p.Client.Alignment.Query().CountX(ctx)
	if err := p.PopulateAlignment(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Alignment entities")
	}
	log.Infof("created %d entities for type Alignment", p.Client.Alignment.Query().CountX(ctx)-start)
	
	start = p.Client.Skill.Query().CountX(ctx)
	if err := p.PopulateSkill(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}
	log.Infof("created %d entities for type Skill", p.Client.Skill.Query().CountX(ctx)-start)
	
	start = p.Client.Proficiency.Query().CountX(ctx)
	if err := p.PopulateProficiency(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Proficiency entities")
	}
	log.Infof("created %d entities for type Proficiency", p.Client.Proficiency.Query().CountX(ctx)-start)
	
	start = p.Client.MagicSchool.Query().CountX(ctx)
	if err := p.PopulateMagicSchool(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate MagicSchool entities")
	}
	log.Infof("created %d entities for type MagicSchool", p.Client.MagicSchool.Query().CountX(ctx)-start)
	

	return nil
}
