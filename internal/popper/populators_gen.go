package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScore(ctx context.Context) ([]*ent.AbilityScoreCreate, error) {
	fpath := "data/AbilityScore.json"
	var v []AbilityScoreWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.AbilityScoreCreate, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(p)
	}

	return creates, nil
}

// PopulateClass populates the Class entities from the JSON data files.
func (p *Popper) PopulateClass(ctx context.Context) ([]*ent.ClassCreate, error) {
	fpath := "data/Class.json"
	var v []ClassWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.ClassCreate, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(p)
	}

	return creates, nil
}

// PopulateRace populates the Race entities from the JSON data files.
func (p *Popper) PopulateRace(ctx context.Context) ([]*ent.RaceCreate, error) {
	fpath := "data/Race.json"
	var v []RaceWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.RaceCreate, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(p)
	}

	return creates, nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) ([]*ent.SkillCreate, error) {
	fpath := "data/Skill.json"
	var v []SkillWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.SkillCreate, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(p)
	}

	return creates, nil
}

// PopulateLanguage populates the Language entities from the JSON data files.
func (p *Popper) PopulateLanguage(ctx context.Context) ([]*ent.LanguageCreate, error) {
	fpath := "data/Language.json"
	var v []LanguageWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.LanguageCreate, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(p)
	}

	return creates, nil
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
	
	if _, err := p.Client.Language.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all Language entities")
	}
	log.Infof("deleted all entities for type Language")
	
	return nil
}

// PopulateAllGen populates all entities generated from the JSON data files.
func (p *Popper) PopulateAllGen(ctx context.Context) error {
	
	createsAbilityScore, err := p.PopulateAbilityScore(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}
	p.Client.AbilityScore.CreateBulk(createsAbilityScore...).ExecX(ctx)
	log.Infof("created %d entities for type AbilityScore", len(createsAbilityScore))
	
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
	
	return nil
}
