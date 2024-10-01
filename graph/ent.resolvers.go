package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/ecshreve/dndgen/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

// AbilityScores is the resolver for the abilityScores field.
func (r *queryResolver) AbilityScores(ctx context.Context) ([]*ent.AbilityScore, error) {
	return r.Client.AbilityScore.Query().All(ctx)
}

// Alignments is the resolver for the alignments field.
func (r *queryResolver) Alignments(ctx context.Context) ([]*ent.Alignment, error) {
	return r.Client.Alignment.Query().All(ctx)
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.CharacterOrder, where *ent.CharacterWhereInput) (*ent.CharacterConnection, error) {
	return r.Client.Character.Query().Paginate(ctx, after, first, before, last,
		ent.WithCharacterOrder(orderBy),
		ent.WithCharacterFilter(where.Filter),
	)
}

// Classes is the resolver for the classes field.
func (r *queryResolver) Classes(ctx context.Context) ([]*ent.Class, error) {
	return r.Client.Class.Query().All(ctx)
}

// Coins is the resolver for the coins field.
func (r *queryResolver) Coins(ctx context.Context) ([]*ent.Coin, error) {
	return r.Client.Coin.Query().All(ctx)
}

// Conditions is the resolver for the conditions field.
func (r *queryResolver) Conditions(ctx context.Context) ([]*ent.Condition, error) {
	return r.Client.Condition.Query().All(ctx)
}

// DamageTypes is the resolver for the damageTypes field.
func (r *queryResolver) DamageTypes(ctx context.Context) ([]*ent.DamageType, error) {
	return r.Client.DamageType.Query().All(ctx)
}

// Equipments is the resolver for the equipments field.
func (r *queryResolver) Equipments(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.EquipmentOrder, where *ent.EquipmentWhereInput) (*ent.EquipmentConnection, error) {
	return r.Client.Equipment.Query().Paginate(ctx, after, first, before, last,
		ent.WithEquipmentOrder(orderBy),
		ent.WithEquipmentFilter(where.Filter),
	)
}

// Feats is the resolver for the feats field.
func (r *queryResolver) Feats(ctx context.Context) ([]*ent.Feat, error) {
	return r.Client.Feat.Query().All(ctx)
}

// Features is the resolver for the features field.
func (r *queryResolver) Features(ctx context.Context) ([]*ent.Feature, error) {
	return r.Client.Feature.Query().All(ctx)
}

// Languages is the resolver for the languages field.
func (r *queryResolver) Languages(ctx context.Context) ([]*ent.Language, error) {
	return r.Client.Language.Query().All(ctx)
}

// MagicSchools is the resolver for the magicSchools field.
func (r *queryResolver) MagicSchools(ctx context.Context) ([]*ent.MagicSchool, error) {
	return r.Client.MagicSchool.Query().All(ctx)
}

// Properties is the resolver for the properties field.
func (r *queryResolver) Properties(ctx context.Context) ([]*ent.Property, error) {
	return r.Client.Property.Query().All(ctx)
}

// Races is the resolver for the races field.
func (r *queryResolver) Races(ctx context.Context) ([]*ent.Race, error) {
	return r.Client.Race.Query().All(ctx)
}

// Rules is the resolver for the rules field.
func (r *queryResolver) Rules(ctx context.Context) ([]*ent.Rule, error) {
	return r.Client.Rule.Query().All(ctx)
}

// RuleSections is the resolver for the ruleSections field.
func (r *queryResolver) RuleSections(ctx context.Context) ([]*ent.RuleSection, error) {
	return r.Client.RuleSection.Query().All(ctx)
}

// Skills is the resolver for the skills field.
func (r *queryResolver) Skills(ctx context.Context) ([]*ent.Skill, error) {
	return r.Client.Skill.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
