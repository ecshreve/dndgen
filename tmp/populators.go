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

type SkillWrapper struct {
	Base
	AbilityScore IDWrapper `json:"ability_score,omitempty"`
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) error {
	fpath := "internal/popper/data/Skill.json"
	var v []SkillWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, wrapper := range v {
		ss := ent.Skill{
			Name: wrapper.Name,
			Desc: wrapper.Desc,
		}
		_, err := p.Client.Skill.Create().
			SetSkill(&ss).
			Save(ctx)
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

type ClassWrapper struct {
	Base
	HitDie       int         `json:"hit_die"`
	SavingThrows []IDWrapper `json:"saving_throws"`
}

// PopulateClass populates the Class entities from the JSON data files.
func (p *Popper) PopulateClass(ctx context.Context) error {
	fpath := "internal/popper/data/Class.json"
	var v []ClassWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, wrapper := range v {
		vv := ent.Class{
			Name:   wrapper.Name,
			HitDie: wrapper.HitDie,
		}

		_, err := p.Client.Class.Create().
			SetClass(&vv).
			Save(ctx)
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

// type ProficiencyWrapper struct {
// 	Base
// 	Category string      `json:"type"`
// 	Classes  []IDWrapper `json:"classes,omitempty"`
// 	Races    []IDWrapper `json:"races,omitempty"`
// }

// // PopulateProficiency populates the Proficiency entities from the JSON data files.
// func (p *Popper) PopulateProficiency(ctx context.Context) error {
// 	fpath := "internal/popper/data/Proficiency.json"
// 	var v []ProficiencyWrapper

// 	if err := LoadJSONFile(fpath, &v); err != nil {
// 		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
// 	}

// 	for _, ww := range v {
// 		vv := ent.Proficiency{
// 			ID:       ww.ID,
// 			Name:     ww.Name,
// 			Category: ww.Category,
// 		}
// 		_, err := p.Client.Proficiency.Create().
// 			SetProficiency(&vv).
// 			AddClassIDs(GetIDStrings(ww.Classes)...).
// 			AddRaceIDs(GetIDStrings(ww.Races)...).
// 			Save(ctx)
// 		if ent.IsConstraintError(err) {
// 			log.Debugf("constraint failed, skipping %s", vv.ID)
// 			continue
// 		}
// 		if err != nil {
// 			return oops.Wrapf(err, "unable to create entity %s", vv.ID)
// 		}
// 	}

// 	return nil
// }

// PopulateRace populates the Race entities from the JSON data files.
func (p *Popper) PopulateRace(ctx context.Context) error {
	fpath := "internal/popper/data/Race.json"
	var v []ent.Race

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Race.Create().SetRace(&vv).Save(ctx)
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

	start = p.Client.Class.Query().CountX(ctx)
	if err := p.PopulateClass(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}
	log.Infof("created %d entities for type Class", p.Client.Class.Query().CountX(ctx)-start)

	start = p.Client.Race.Query().CountX(ctx)
	if err := p._PopulateRace(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}
	log.Infof("created %d entities for type Race", p.Client.Race.Query().CountX(ctx)-start)

	return nil
}
