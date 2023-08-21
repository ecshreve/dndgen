package dndgen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ecshreve/dndgen/ent"
)

func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AbilityScores(ctx context.Context) ([]*ent.AbilityScore, error) {
	fmt.Println("queryResolver.AbilityScores")
	return r.Client.AbilityScore.Query().All(ctx)
}

func (r *queryResolver) Skills(ctx context.Context) ([]*ent.Skill, error) {
	fmt.Println("queryResolver.Skills")
	return r.Client.Skill.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
