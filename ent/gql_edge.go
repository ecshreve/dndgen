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

func (a *Armor) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := a.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryEquipment().Only(ctx)
	}
	return result, err
}

func (a *Armor) ArmorClass(ctx context.Context) (*ArmorClass, error) {
	result, err := a.Edges.ArmorClassOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryArmorClass().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ac *ArmorClass) Armor(ctx context.Context) (*Armor, error) {
	result, err := ac.Edges.ArmorOrErr()
	if IsNotLoaded(err) {
		result, err = ac.QueryArmor().Only(ctx)
	}
	return result, MaskNotFound(err)
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

func (d *Damage) DamageType(ctx context.Context) (*DamageType, error) {
	result, err := d.Edges.DamageTypeOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryDamageType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Equipment) Cost(ctx context.Context) (*Cost, error) {
	result, err := e.Edges.CostOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryCost().Only(ctx)
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

func (e *Equipment) Gear(ctx context.Context) (*Gear, error) {
	result, err := e.Edges.GearOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryGear().Only(ctx)
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

func (ge *Gear) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := ge.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = ge.QueryEquipment().Only(ctx)
	}
	return result, err
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

func (t *Tool) Equipment(ctx context.Context) (*Equipment, error) {
	result, err := t.Edges.EquipmentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryEquipment().Only(ctx)
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
	return result, err
}

func (w *Weapon) WeaponRange(ctx context.Context) (*WeaponRange, error) {
	result, err := w.Edges.WeaponRangeOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryWeaponRange().Only(ctx)
	}
	return result, MaskNotFound(err)
}
