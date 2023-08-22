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

// PopulateClass populates the Class entities from the JSON data files.
func (p *Popper) PopulateClass(ctx context.Context) error {
	fpath := "internal/popper/data/Class.json"
	var v []ent.Class

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		_, err := p.Client.Class.Create().SetClass(&vv).Save(ctx)
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


// CleanUp clears all entities from the database.
func (p *Popper) CleanUp(ctx context.Context) error {
	p.Client.Equipment.Delete().ExecX(ctx)
	p.Client.Weapon.Delete().ExecX(ctx)
	p.Client.Armor.Delete().ExecX(ctx)
	p.Client.Gear.Delete().ExecX(ctx)
	p.Client.Vehicle.Delete().ExecX(ctx)
	p.Client.Tool.Delete().ExecX(ctx)

	log.Infof("deleted all entities for type Equipment and subtypes")

	p.Client.Proficiency.Delete().ExecX(ctx)
	log.Infof("deleted all entities for type Proficiency")

	
	if _, err := p.Client.AbilityScore.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all AbilityScore entities")
	}
	log.Infof("deleted all entities for type AbilityScore")

	
	if _, err := p.Client.Class.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all Class entities")
	}
	log.Infof("deleted all entities for type Class")

	
	if _, err := p.Client.Race.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all Race entities")
	}
	log.Infof("deleted all entities for type Race")

	
	if _, err := p.Client.Skill.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all Skill entities")
	}
	log.Infof("deleted all entities for type Skill")

	

	return nil
}

// PopulateAll populates all entities from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	var start int

	start = p.Client.Equipment.Query().CountX(ctx)
	if err := p.PopulateEquipment(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Equipment entities")
	}
	log.Infof("created %d entities for type Equipment", p.Client.Equipment.Query().CountX(ctx)-start)
	log.Infof("created %d entities for type Weapon", p.Client.Weapon.Query().CountX(ctx)-start)
	log.Infof("created %d entities for type Armor", p.Client.Armor.Query().CountX(ctx)-start)
	log.Infof("created %d entities for type Gear", p.Client.Gear.Query().CountX(ctx)-start)
	log.Infof("created %d entities for type Tool", p.Client.Tool.Query().CountX(ctx)-start)

	
	start = p.Client.AbilityScore.Query().CountX(ctx)
	if err := p.PopulateAbilityScore(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}
	log.Infof("created %d entities for type AbilityScore", p.Client.AbilityScore.Query().CountX(ctx)-start)
	
	start = p.Client.Class.Query().CountX(ctx)
	if err := p.PopulateClass(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}
	log.Infof("created %d entities for type Class", p.Client.Class.Query().CountX(ctx)-start)
	
	start = p.Client.Race.Query().CountX(ctx)
	if err := p.PopulateRace(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}
	log.Infof("created %d entities for type Race", p.Client.Race.Query().CountX(ctx)-start)
	
	start = p.Client.Skill.Query().CountX(ctx)
	if err := p.PopulateSkill(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}
	log.Infof("created %d entities for type Skill", p.Client.Skill.Query().CountX(ctx)-start)
	

	if err := p.PopulateProficiency(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate Proficiency entities")
	}

	return nil
}
