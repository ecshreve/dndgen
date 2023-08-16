package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ecshreve/dndgen/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) AbilityScores(ctx context.Context) ([]*ent.AbilityScore, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Alignments(ctx context.Context) ([]*ent.Alignment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Classes(ctx context.Context) ([]*ent.Class, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Conditions(ctx context.Context) ([]*ent.Condition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DamageTypes(ctx context.Context) ([]*ent.DamageType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EquipmentSlice(ctx context.Context) ([]*ent.Equipment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EquipmentCategories(ctx context.Context) ([]*ent.EquipmentCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Languages(ctx context.Context) ([]*ent.Language, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MagicSchools(ctx context.Context) ([]*ent.MagicSchool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Proficiencies(ctx context.Context) ([]*ent.Proficiency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Races(ctx context.Context) ([]*ent.Race, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Skills(ctx context.Context) ([]*ent.Skill, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
