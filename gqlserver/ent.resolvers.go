package dndgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/ecshreve/dndgen/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

func (r *queryResolver) AbilityScores(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AbilityScoreOrder, where *ent.AbilityScoreWhereInput) (*ent.AbilityScoreConnection, error) {
	fmt.Println("queryResolver.AbilityScores")
	return r.Client.AbilityScore.Query().Paginate(ctx, after, first, before, last, ent.WithAbilityScoreOrder(orderBy), ent.WithAbilityScoreFilter(where.Filter))
}

func (r *queryResolver) Armors(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.ArmorOrder, where *ent.ArmorWhereInput) (*ent.ArmorConnection, error) {
	fmt.Println("queryResolver.Armors")
	return r.Client.Armor.Query().Paginate(ctx, after, first, before, last, ent.WithArmorOrder(orderBy), ent.WithArmorFilter(where.Filter))
}

func (r *queryResolver) Classes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.ClassOrder, where *ent.ClassWhereInput) (*ent.ClassConnection, error) {
	fmt.Println("queryResolver.Classes")
	return r.Client.Class.Query().Paginate(ctx, after, first, before, last, ent.WithClassOrder(orderBy), ent.WithClassFilter(where.Filter))
}

func (r *queryResolver) DamageTypes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.DamageTypeOrder, where *ent.DamageTypeWhereInput) (*ent.DamageTypeConnection, error) {
	return r.Client.DamageType.Query().Paginate(ctx, after, first, before, last, ent.WithDamageTypeOrder(orderBy), ent.WithDamageTypeFilter(where.Filter))
}

func (r *queryResolver) EquipmentSlice(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.EquipmentOrder, where *ent.EquipmentWhereInput) (*ent.EquipmentConnection, error) {
	fmt.Println("queryResolver.EquipmentSlice")
	return r.Client.Equipment.Query().Paginate(ctx, after, first, before, last, ent.WithEquipmentOrder(orderBy), ent.WithEquipmentFilter(where.Filter))
}

func (r *queryResolver) Gears(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.GearOrder, where *ent.GearWhereInput) (*ent.GearConnection, error) {
	fmt.Println("queryResolver.Gears")
	return r.Client.Gear.Query().Paginate(ctx, after, first, before, last, ent.WithGearOrder(orderBy), ent.WithGearFilter(where.Filter))
}

func (r *queryResolver) Languages(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.LanguageOrder, where *ent.LanguageWhereInput) (*ent.LanguageConnection, error) {
	fmt.Println("queryResolver.Languages")
	return r.Client.Language.Query().Paginate(ctx, after, first, before, last, ent.WithLanguageOrder(orderBy), ent.WithLanguageFilter(where.Filter))
}

func (r *queryResolver) MagicSchools(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MagicSchoolOrder, where *ent.MagicSchoolWhereInput) (*ent.MagicSchoolConnection, error) {
	fmt.Println("queryResolver.MagicSchools")
	return r.Client.MagicSchool.Query().Paginate(ctx, after, first, before, last, ent.WithMagicSchoolOrder(orderBy), ent.WithMagicSchoolFilter(where.Filter))
}

func (r *queryResolver) Proficiencies(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.ProficiencyOrder, where *ent.ProficiencyWhereInput) (*ent.ProficiencyConnection, error) {
	fmt.Println("queryResolver.Proficiencies")
	return r.Client.Proficiency.Query().Paginate(ctx, after, first, before, last, ent.WithProficiencyOrder(orderBy), ent.WithProficiencyFilter(where.Filter))
}

func (r *queryResolver) Races(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.RaceOrder, where *ent.RaceWhereInput) (*ent.RaceConnection, error) {
	fmt.Println("queryResolver.Races")
	return r.Client.Race.Query().Paginate(ctx, after, first, before, last, ent.WithRaceOrder(orderBy), ent.WithRaceFilter(where.Filter))
}

func (r *queryResolver) Rules(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.RuleOrder, where *ent.RuleWhereInput) (*ent.RuleConnection, error) {
	fmt.Println("queryResolver.Rules")
	return r.Client.Rule.Query().Paginate(ctx, after, first, before, last, ent.WithRuleOrder(orderBy), ent.WithRuleFilter(where.Filter))
}

func (r *queryResolver) RuleSections(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.RuleSectionOrder, where *ent.RuleSectionWhereInput) (*ent.RuleSectionConnection, error) {
	fmt.Println("queryResolver.RuleSections")
	return r.Client.RuleSection.Query().Paginate(ctx, after, first, before, last, ent.WithRuleSectionOrder(orderBy), ent.WithRuleSectionFilter(where.Filter))
}

func (r *queryResolver) Skills(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SkillOrder, where *ent.SkillWhereInput) (*ent.SkillConnection, error) {
	fmt.Println("queryResolver.Skills")
	return r.Client.Skill.Query().Paginate(ctx, after, first, before, last, ent.WithSkillOrder(orderBy), ent.WithSkillFilter(where.Filter))
}

func (r *queryResolver) Subraces(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SubraceOrder, where *ent.SubraceWhereInput) (*ent.SubraceConnection, error) {
	fmt.Println("queryResolver.Subraces")
	return r.Client.Subrace.Query().Paginate(ctx, after, first, before, last, ent.WithSubraceOrder(orderBy), ent.WithSubraceFilter(where.Filter))
}

func (r *queryResolver) Tools(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.ToolOrder, where *ent.ToolWhereInput) (*ent.ToolConnection, error) {
	fmt.Println("queryResolver.Tools")
	return r.Client.Tool.Query().Paginate(ctx, after, first, before, last, ent.WithToolOrder(orderBy), ent.WithToolFilter(where.Filter))
}

func (r *queryResolver) Traits(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.TraitOrder, where *ent.TraitWhereInput) (*ent.TraitConnection, error) {
	fmt.Println("queryResolver.Traits")
	return r.Client.Trait.Query().Paginate(ctx, after, first, before, last, ent.WithTraitOrder(orderBy), ent.WithTraitFilter(where.Filter))
}

func (r *queryResolver) Vehicles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.VehicleOrder, where *ent.VehicleWhereInput) (*ent.VehicleConnection, error) {
	fmt.Println("queryResolver.Vehicles")
	return r.Client.Vehicle.Query().Paginate(ctx, after, first, before, last, ent.WithVehicleOrder(orderBy), ent.WithVehicleFilter(where.Filter))
}

func (r *queryResolver) Weapons(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.WeaponOrder, where *ent.WeaponWhereInput) (*ent.WeaponConnection, error) {
	fmt.Println("queryResolver.Weapons")
	return r.Client.Weapon.Query().Paginate(ctx, after, first, before, last, ent.WithWeaponOrder(orderBy), ent.WithWeaponFilter(where.Filter))
}

func (r *queryResolver) WeaponDamages(ctx context.Context) ([]*ent.WeaponDamage, error) {
	fmt.Println("queryResolver.WeaponDamages")
	return r.Client.WeaponDamage.Query().All(ctx)
}

func (r *queryResolver) WeaponProperties(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.WeaponPropertyOrder, where *ent.WeaponPropertyWhereInput) (*ent.WeaponPropertyConnection, error) {
	fmt.Println("queryResolver.WeaponProperties")
	return r.Client.WeaponProperty.Query().Paginate(ctx, after, first, before, last, ent.WithWeaponPropertyOrder(orderBy), ent.WithWeaponPropertyFilter(where.Filter))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
