package dndgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ecshreve/dndgen/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

func (r *queryResolver) AbilityScores(ctx context.Context) ([]*ent.AbilityScore, error) {
	fmt.Println("queryResolver.AbilityScores")
	return r.Client.AbilityScore.Query().All(ctx)
}

func (r *queryResolver) Armors(ctx context.Context) ([]*ent.Armor, error) {
	fmt.Println("queryResolver.Armors")
	return r.Client.Armor.Query().All(ctx)
}

func (r *queryResolver) Classes(ctx context.Context) ([]*ent.Class, error) {
	fmt.Println("queryResolver.Classes")
	return r.Client.Class.Query().All(ctx)
}

func (r *queryResolver) EquipmentSlice(ctx context.Context) ([]*ent.Equipment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Gears(ctx context.Context) ([]*ent.Gear, error) {
	fmt.Println("queryResolver.Gears")
	return r.Client.Gear.Query().All(ctx)
}

func (r *queryResolver) Races(ctx context.Context) ([]*ent.Race, error) {
	fmt.Println("queryResolver.Races")
	return r.Client.Race.Query().All(ctx)
}

func (r *queryResolver) Skills(ctx context.Context) ([]*ent.Skill, error) {
	fmt.Println("queryResolver.Skills")
	return r.Client.Skill.Query().All(ctx)
}

func (r *queryResolver) Tools(ctx context.Context) ([]*ent.Tool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Weapons(ctx context.Context) ([]*ent.Weapon, error) {
	fmt.Println("queryResolver.Weapons")
	return r.Client.Weapon.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
