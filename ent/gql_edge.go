// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ab *AbilityBonus) AbilityScore(ctx context.Context) (*AbilityScore, error) {
	result, err := ab.Edges.AbilityScoreOrErr()
	if IsNotLoaded(err) {
		result, err = ab.QueryAbilityScore().Only(ctx)
	}
	return result, err
}

func (ab *AbilityBonus) Race(ctx context.Context) (*Race, error) {
	result, err := ab.Edges.RaceOrErr()
	if IsNotLoaded(err) {
		result, err = ab.QueryRace().Only(ctx)
	}
	return result, err
}

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

func (as *AbilityScore) AbilityBonuses(ctx context.Context) (result []*AbilityBonus, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = as.NamedAbilityBonuses(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = as.Edges.AbilityBonusesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = as.QueryAbilityBonuses().All(ctx)
	}
	return result, err
}

func (d *Damage) DamageType(ctx context.Context) (*DamageType, error) {
	result, err := d.Edges.DamageTypeOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryDamageType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) EquipmentCosts(ctx context.Context) (*EquipmentCost, error) {
	result, err := e.Edges.EquipmentCostsOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryEquipmentCosts().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ec *EquipmentCost) Coin(ctx context.Context) (*Coin, error) {
	result, err := ec.Edges.CoinOrErr()
	if IsNotLoaded(err) {
		result, err = ec.QueryCoin().Only(ctx)
	}
	return result, err
}

func (ec *EquipmentCost) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := ec.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = ec.QueryEquipment().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Language) Races(ctx context.Context) (result []*Race, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedRaces(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.RacesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryRaces().All(ctx)
	}
	return result, err
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

func (r *Race) AbilityBonuses(ctx context.Context) (result []*AbilityBonus, err error) {
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

func (w *Weapon) Damage(ctx context.Context) (*Damage, error) {
	result, err := w.Edges.DamageOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryDamage().Only(ctx)
	}
	return result, MaskNotFound(err)
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

func (w *Weapon) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := w.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryEquipment().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Weapon) WeaponRange(ctx context.Context) (*WeaponRange, error) {
	result, err := w.Edges.WeaponRangeOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryWeaponRange().Only(ctx)
	}
	return result, MaskNotFound(err)
}
