// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (as *AbilityScore) Skills(ctx context.Context) (result []*Skill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = as.NamedSkills(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = as.Edges.SkillsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = as.QuerySkills().All(ctx)
	}
	return result, err
}

func (a *Armor) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := a.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (c *Character) Race(ctx context.Context) (*Race, error) {
	result, err := c.Edges.RaceOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryRace().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Character) Class(ctx context.Context) (*Class, error) {
	result, err := c.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryClass().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Character) Alignment(ctx context.Context) (*Alignment, error) {
	result, err := c.Edges.AlignmentOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryAlignment().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Character) CharacterAbilityScores(ctx context.Context) (result []*CharacterAbilityScore, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCharacterAbilityScores(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CharacterAbilityScoresOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCharacterAbilityScores().All(ctx)
	}
	return result, err
}

func (c *Character) CharacterSkills(ctx context.Context) (result []*CharacterSkill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCharacterSkills(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CharacterSkillsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCharacterSkills().All(ctx)
	}
	return result, err
}

func (c *Character) CharacterProficiencies(ctx context.Context) (result []*CharacterProficiency, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCharacterProficiencies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CharacterProficienciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCharacterProficiencies().All(ctx)
	}
	return result, err
}

func (cas *CharacterAbilityScore) Character(ctx context.Context) (*Character, error) {
	result, err := cas.Edges.CharacterOrErr()
	if IsNotLoaded(err) {
		result, err = cas.QueryCharacter().Only(ctx)
	}
	return result, err
}

func (cas *CharacterAbilityScore) AbilityScore(ctx context.Context) (*AbilityScore, error) {
	result, err := cas.Edges.AbilityScoreOrErr()
	if IsNotLoaded(err) {
		result, err = cas.QueryAbilityScore().Only(ctx)
	}
	return result, err
}

func (cas *CharacterAbilityScore) CharacterSkills(ctx context.Context) (result []*CharacterSkill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cas.NamedCharacterSkills(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cas.Edges.CharacterSkillsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cas.QueryCharacterSkills().All(ctx)
	}
	return result, err
}

func (cp *CharacterProficiency) Character(ctx context.Context) (*Character, error) {
	result, err := cp.Edges.CharacterOrErr()
	if IsNotLoaded(err) {
		result, err = cp.QueryCharacter().Only(ctx)
	}
	return result, err
}

func (cp *CharacterProficiency) Proficiency(ctx context.Context) (*Proficiency, error) {
	result, err := cp.Edges.ProficiencyOrErr()
	if IsNotLoaded(err) {
		result, err = cp.QueryProficiency().Only(ctx)
	}
	return result, err
}

func (cp *CharacterProficiency) CharacterSkill(ctx context.Context) (*CharacterSkill, error) {
	result, err := cp.Edges.CharacterSkillOrErr()
	if IsNotLoaded(err) {
		result, err = cp.QueryCharacterSkill().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cs *CharacterSkill) Character(ctx context.Context) (*Character, error) {
	result, err := cs.Edges.CharacterOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QueryCharacter().Only(ctx)
	}
	return result, err
}

func (cs *CharacterSkill) Skill(ctx context.Context) (*Skill, error) {
	result, err := cs.Edges.SkillOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QuerySkill().Only(ctx)
	}
	return result, err
}

func (cs *CharacterSkill) CharacterAbilityScore(ctx context.Context) (*CharacterAbilityScore, error) {
	result, err := cs.Edges.CharacterAbilityScoreOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QueryCharacterAbilityScore().Only(ctx)
	}
	return result, err
}

func (cs *CharacterSkill) CharacterProficiency(ctx context.Context) (*CharacterProficiency, error) {
	result, err := cs.Edges.CharacterProficiencyOrErr()
	if IsNotLoaded(err) {
		result, err = cs.QueryCharacterProficiency().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Class) Proficiencies(ctx context.Context) (result []*Proficiency, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedProficiencies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ProficienciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryProficiencies().All(ctx)
	}
	return result, err
}

func (c *Class) ProficiencyOptions(ctx context.Context) (result []*ProficiencyChoice, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedProficiencyOptions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ProficiencyOptionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryProficiencyOptions().All(ctx)
	}
	return result, err
}

func (c *Class) StartingEquipment(ctx context.Context) (result []*EquipmentEntry, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedStartingEquipment(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.StartingEquipmentOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryStartingEquipment().All(ctx)
	}
	return result, err
}

func (c *Class) SavingThrows(ctx context.Context) (result []*AbilityScore, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedSavingThrows(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.SavingThrowsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QuerySavingThrows().All(ctx)
	}
	return result, err
}

func (c *Class) Characters(ctx context.Context) (result []*Character, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCharacters(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CharactersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCharacters().All(ctx)
	}
	return result, err
}

func (c *Cost) Coin(ctx context.Context) (*Coin, error) {
	result, err := c.Edges.CoinOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCoin().Only(ctx)
	}
	return result, err
}

func (c *Cost) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := c.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryEquipment().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (dt *DamageType) Weapons(ctx context.Context) (result []*Weapon, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = dt.NamedWeapons(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = dt.Edges.WeaponsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = dt.QueryWeapons().All(ctx)
	}
	return result, err
}

func (e *Equipment) Cost(ctx context.Context) (*Cost, error) {
	result, err := e.Edges.CostOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryCost().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Gear(ctx context.Context) (*Gear, error) {
	result, err := e.Edges.GearOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryGear().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Tool(ctx context.Context) (*Tool, error) {
	result, err := e.Edges.ToolOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryTool().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Weapon(ctx context.Context) (*Weapon, error) {
	result, err := e.Edges.WeaponOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryWeapon().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Vehicle(ctx context.Context) (*Vehicle, error) {
	result, err := e.Edges.VehicleOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryVehicle().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Armor(ctx context.Context) (*Armor, error) {
	result, err := e.Edges.ArmorOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryArmor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) EquipmentEntries(ctx context.Context) (result []*EquipmentEntry, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = e.NamedEquipmentEntries(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = e.Edges.EquipmentEntriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = e.QueryEquipmentEntries().All(ctx)
	}
	return result, err
}

func (ee *EquipmentEntry) Class(ctx context.Context) (result []*Class, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ee.NamedClass(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ee.Edges.ClassOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ee.QueryClass().All(ctx)
	}
	return result, err
}

func (ee *EquipmentEntry) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := ee.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = ee.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (f *Feature) Prerequisites(ctx context.Context) (result []*Prerequisite, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = f.NamedPrerequisites(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = f.Edges.PrerequisitesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = f.QueryPrerequisites().All(ctx)
	}
	return result, err
}

func (ge *Gear) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := ge.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = ge.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (l *Language) Race(ctx context.Context) (result []*Race, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedRace(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.RaceOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryRace().All(ctx)
	}
	return result, err
}

func (l *Language) Options(ctx context.Context) (result []*LanguageChoice, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedOptions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.OptionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryOptions().All(ctx)
	}
	return result, err
}

func (lc *LanguageChoice) Languages(ctx context.Context) (result []*Language, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = lc.NamedLanguages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = lc.Edges.LanguagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = lc.QueryLanguages().All(ctx)
	}
	return result, err
}

func (lc *LanguageChoice) Race(ctx context.Context) (*Race, error) {
	result, err := lc.Edges.RaceOrErr()
	if IsNotLoaded(err) {
		result, err = lc.QueryRace().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Prerequisite) Feature(ctx context.Context) (*Feature, error) {
	result, err := pr.Edges.FeatureOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryFeature().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Proficiency) Race(ctx context.Context) (result []*Race, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedRace(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.RaceOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryRace().All(ctx)
	}
	return result, err
}

func (pr *Proficiency) Class(ctx context.Context) (result []*Class, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedClass(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.ClassOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryClass().All(ctx)
	}
	return result, err
}

func (pc *ProficiencyChoice) Proficiencies(ctx context.Context) (result []*Proficiency, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pc.NamedProficiencies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pc.Edges.ProficienciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pc.QueryProficiencies().All(ctx)
	}
	return result, err
}

func (pc *ProficiencyChoice) Race(ctx context.Context) (*Race, error) {
	result, err := pc.Edges.RaceOrErr()
	if IsNotLoaded(err) {
		result, err = pc.QueryRace().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pc *ProficiencyChoice) Class(ctx context.Context) (*Class, error) {
	result, err := pc.Edges.ClassOrErr()
	if IsNotLoaded(err) {
		result, err = pc.QueryClass().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Property) Weapons(ctx context.Context) (result []*Weapon, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedWeapons(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.WeaponsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryWeapons().All(ctx)
	}
	return result, err
}

func (r *Race) Traits(ctx context.Context) (result []*Trait, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedTraits(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.TraitsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryTraits().All(ctx)
	}
	return result, err
}

func (r *Race) StartingProficiencies(ctx context.Context) (result []*Proficiency, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedStartingProficiencies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.StartingProficienciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryStartingProficiencies().All(ctx)
	}
	return result, err
}

func (r *Race) StartingProficiencyOptions(ctx context.Context) (result []*ProficiencyChoice, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedStartingProficiencyOptions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.StartingProficiencyOptionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryStartingProficiencyOptions().All(ctx)
	}
	return result, err
}

func (r *Race) AbilityBonuses(ctx context.Context) (result []*AbilityScore, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedAbilityBonuses(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.AbilityBonusesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryAbilityBonuses().All(ctx)
	}
	return result, err
}

func (r *Race) Languages(ctx context.Context) (result []*Language, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedLanguages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.LanguagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryLanguages().All(ctx)
	}
	return result, err
}

func (r *Race) LanguageOptions(ctx context.Context) (*LanguageChoice, error) {
	result, err := r.Edges.LanguageOptionsOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryLanguageOptions().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Race) Characters(ctx context.Context) (result []*Character, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedCharacters(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.CharactersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryCharacters().All(ctx)
	}
	return result, err
}

func (r *Rule) Sections(ctx context.Context) (result []*RuleSection, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedSections(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.SectionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QuerySections().All(ctx)
	}
	return result, err
}

func (rs *RuleSection) Rule(ctx context.Context) (*Rule, error) {
	result, err := rs.Edges.RuleOrErr()
	if IsNotLoaded(err) {
		result, err = rs.QueryRule().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Skill) AbilityScore(ctx context.Context) (*AbilityScore, error) {
	result, err := s.Edges.AbilityScoreOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryAbilityScore().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Tool) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := t.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (t *Trait) Race(ctx context.Context) (result []*Race, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedRace(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.RaceOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryRace().All(ctx)
	}
	return result, err
}

func (v *Vehicle) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := v.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = v.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (w *Weapon) Properties(ctx context.Context) (result []*Property, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedProperties(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.PropertiesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryProperties().All(ctx)
	}
	return result, err
}

func (w *Weapon) DamageType(ctx context.Context) (*DamageType, error) {
	result, err := w.Edges.DamageTypeOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryDamageType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Weapon) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := w.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryEquipment().Only(ctx)
	}
	return result, err
}
