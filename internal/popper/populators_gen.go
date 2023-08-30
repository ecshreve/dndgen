package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScore(ctx context.Context) ([]*ent.AbilityScore, error) {
	fpath := "data/AbilityScore.json"
	var v []ent.AbilityScore

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.AbilityScoreCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.AbilityScore.Create().SetAbilityScore(&vv)
	}

	created, err := p.Client.AbilityScore.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save AbilityScore entities")
	}
	log.Infof("created %d entities for type AbilityScore", len(created))

	p.PopulateAbilityScoreEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) ([]*ent.Skill, error) {
	fpath := "data/Skill.json"
	var v []ent.Skill

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.SkillCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Skill.Create().SetSkill(&vv)
	}

	created, err := p.Client.Skill.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save Skill entities")
	}
	log.Infof("created %d entities for type Skill", len(created))

	p.PopulateSkillEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateLanguage populates the Language entities from the JSON data files.
func (p *Popper) PopulateLanguage(ctx context.Context) ([]*ent.Language, error) {
	fpath := "data/Language.json"
	var v []ent.Language

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.LanguageCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Language.Create().SetLanguage(&vv)
	}

	created, err := p.Client.Language.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save Language entities")
	}
	log.Infof("created %d entities for type Language", len(created))

	p.PopulateLanguageEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateDamageType populates the DamageType entities from the JSON data files.
func (p *Popper) PopulateDamageType(ctx context.Context) ([]*ent.DamageType, error) {
	fpath := "data/DamageType.json"
	var v []ent.DamageType

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.DamageTypeCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.DamageType.Create().SetDamageType(&vv)
	}

	created, err := p.Client.DamageType.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save DamageType entities")
	}
	log.Infof("created %d entities for type DamageType", len(created))

	p.PopulateDamageTypeEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateRace populates the Race entities from the JSON data files.
func (p *Popper) PopulateRace(ctx context.Context) ([]*ent.Race, error) {
	fpath := "data/Race.json"
	var v []ent.Race

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.RaceCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Race.Create().SetRace(&vv)
	}

	created, err := p.Client.Race.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save Race entities")
	}
	log.Infof("created %d entities for type Race", len(created))

	p.PopulateRaceEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateClass populates the Class entities from the JSON data files.
func (p *Popper) PopulateClass(ctx context.Context) ([]*ent.Class, error) {
	fpath := "data/Class.json"
	var v []ent.Class

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.ClassCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Class.Create().SetClass(&vv)
	}

	created, err := p.Client.Class.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save Class entities")
	}
	log.Infof("created %d entities for type Class", len(created))

	p.PopulateClassEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateWeaponProperty populates the WeaponProperty entities from the JSON data files.
func (p *Popper) PopulateWeaponProperty(ctx context.Context) ([]*ent.WeaponProperty, error) {
	fpath := "data/WeaponProperty.json"
	var v []ent.WeaponProperty

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.WeaponPropertyCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.WeaponProperty.Create().SetWeaponProperty(&vv)
	}

	created, err := p.Client.WeaponProperty.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save WeaponProperty entities")
	}
	log.Infof("created %d entities for type WeaponProperty", len(created))

	p.PopulateWeaponPropertyEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateMagicSchool populates the MagicSchool entities from the JSON data files.
func (p *Popper) PopulateMagicSchool(ctx context.Context) ([]*ent.MagicSchool, error) {
	fpath := "data/MagicSchool.json"
	var v []ent.MagicSchool

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.MagicSchoolCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.MagicSchool.Create().SetMagicSchool(&vv)
	}

	created, err := p.Client.MagicSchool.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save MagicSchool entities")
	}
	log.Infof("created %d entities for type MagicSchool", len(created))

	p.PopulateMagicSchoolEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateRuleSection populates the RuleSection entities from the JSON data files.
func (p *Popper) PopulateRuleSection(ctx context.Context) ([]*ent.RuleSection, error) {
	fpath := "data/RuleSection.json"
	var v []ent.RuleSection

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.RuleSectionCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.RuleSection.Create().SetRuleSection(&vv)
	}

	created, err := p.Client.RuleSection.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save RuleSection entities")
	}
	log.Infof("created %d entities for type RuleSection", len(created))

	p.PopulateRuleSectionEdges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

