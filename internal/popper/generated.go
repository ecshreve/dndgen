// GENERATED BY popgen
// DO NOT EDIT
package popper

import (
	"context"
	"fmt"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
	
	"github.com/charmbracelet/log"
)

// PopulateAll populates all the entities from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	
	if _, err := p.PopulateAbilityScore(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateSkill(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateLanguage(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateAlignment(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateDamageType(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateFeature(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateFeat(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateCondition(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateProperty(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateMagicSchool(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateRuleSection(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateRule(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateTrait(ctx); err != nil {
		return err
	}
	
	if _, err := p.PopulateCoin(ctx); err != nil {
		return err
	}
	
	return nil
}


// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScore(ctx context.Context) ([]*ent.AbilityScore, error) {
	fpath := fmt.Sprintf("%s/AbilityScore.json", p.DataDir)
	var v []ent.AbilityScore

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.AbilityScoreCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.AbilityScore.Create().SetAbilityScore(&vv)
	}

	created, err := p.Client.AbilityScore.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "AbilityScore")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill(ctx context.Context) ([]*ent.Skill, error) {
	fpath := fmt.Sprintf("%s/Skill.json", p.DataDir)
	var v []ent.Skill

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.SkillCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Skill.Create().SetSkill(&vv)
	}

	created, err := p.Client.Skill.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Skill")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateLanguage populates the Language entities from the JSON data files.
func (p *Popper) PopulateLanguage(ctx context.Context) ([]*ent.Language, error) {
	fpath := fmt.Sprintf("%s/Language.json", p.DataDir)
	var v []ent.Language

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.LanguageCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Language.Create().SetLanguage(&vv)
	}

	created, err := p.Client.Language.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Language")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateAlignment populates the Alignment entities from the JSON data files.
func (p *Popper) PopulateAlignment(ctx context.Context) ([]*ent.Alignment, error) {
	fpath := fmt.Sprintf("%s/Alignment.json", p.DataDir)
	var v []ent.Alignment

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.AlignmentCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Alignment.Create().SetAlignment(&vv)
	}

	created, err := p.Client.Alignment.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Alignment")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateDamageType populates the DamageType entities from the JSON data files.
func (p *Popper) PopulateDamageType(ctx context.Context) ([]*ent.DamageType, error) {
	fpath := fmt.Sprintf("%s/DamageType.json", p.DataDir)
	var v []ent.DamageType

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.DamageTypeCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.DamageType.Create().SetDamageType(&vv)
	}

	created, err := p.Client.DamageType.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "DamageType")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateFeature populates the Feature entities from the JSON data files.
func (p *Popper) PopulateFeature(ctx context.Context) ([]*ent.Feature, error) {
	fpath := fmt.Sprintf("%s/Feature.json", p.DataDir)
	var v []ent.Feature

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.FeatureCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Feature.Create().SetFeature(&vv)
	}

	created, err := p.Client.Feature.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Feature")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateFeat populates the Feat entities from the JSON data files.
func (p *Popper) PopulateFeat(ctx context.Context) ([]*ent.Feat, error) {
	fpath := fmt.Sprintf("%s/Feat.json", p.DataDir)
	var v []ent.Feat

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.FeatCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Feat.Create().SetFeat(&vv)
	}

	created, err := p.Client.Feat.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Feat")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateCondition populates the Condition entities from the JSON data files.
func (p *Popper) PopulateCondition(ctx context.Context) ([]*ent.Condition, error) {
	fpath := fmt.Sprintf("%s/Condition.json", p.DataDir)
	var v []ent.Condition

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.ConditionCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Condition.Create().SetCondition(&vv)
	}

	created, err := p.Client.Condition.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Condition")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateProperty populates the Property entities from the JSON data files.
func (p *Popper) PopulateProperty(ctx context.Context) ([]*ent.Property, error) {
	fpath := fmt.Sprintf("%s/Property.json", p.DataDir)
	var v []ent.Property

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.PropertyCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Property.Create().SetProperty(&vv)
	}

	created, err := p.Client.Property.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Property")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateMagicSchool populates the MagicSchool entities from the JSON data files.
func (p *Popper) PopulateMagicSchool(ctx context.Context) ([]*ent.MagicSchool, error) {
	fpath := fmt.Sprintf("%s/MagicSchool.json", p.DataDir)
	var v []ent.MagicSchool

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.MagicSchoolCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.MagicSchool.Create().SetMagicSchool(&vv)
	}

	created, err := p.Client.MagicSchool.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "MagicSchool")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateRuleSection populates the RuleSection entities from the JSON data files.
func (p *Popper) PopulateRuleSection(ctx context.Context) ([]*ent.RuleSection, error) {
	fpath := fmt.Sprintf("%s/RuleSection.json", p.DataDir)
	var v []ent.RuleSection

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.RuleSectionCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.RuleSection.Create().SetRuleSection(&vv)
	}

	created, err := p.Client.RuleSection.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "RuleSection")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateRule populates the Rule entities from the JSON data files.
func (p *Popper) PopulateRule(ctx context.Context) ([]*ent.Rule, error) {
	fpath := fmt.Sprintf("%s/Rule.json", p.DataDir)
	var v []ent.Rule

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.RuleCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Rule.Create().SetRule(&vv)
	}

	created, err := p.Client.Rule.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Rule")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateTrait populates the Trait entities from the JSON data files.
func (p *Popper) PopulateTrait(ctx context.Context) ([]*ent.Trait, error) {
	fpath := fmt.Sprintf("%s/Trait.json", p.DataDir)
	var v []ent.Trait

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.TraitCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Trait.Create().SetTrait(&vv)
	}

	created, err := p.Client.Trait.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Trait")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateCoin populates the Coin entities from the JSON data files.
func (p *Popper) PopulateCoin(ctx context.Context) ([]*ent.Coin, error) {
	fpath := fmt.Sprintf("%s/Coin.json", p.DataDir)
	var v []ent.Coin

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.CoinCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Coin.Create().SetCoin(&vv)
	}

	created, err := p.Client.Coin.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Coin")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}
