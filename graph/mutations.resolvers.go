package generated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
)

// CreateCharacter is the resolver for the createCharacter field.
func (r *mutationResolver) CreateCharacter(ctx context.Context, input ent.CreateCharacterInput) (*ent.Character, error) {
	cl := r.Client

	ch, err := cl.Character.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	as, err := utils.HandleCreateCharacterAbilityScores(ctx, cl, ch)
	if err != nil {
		return nil, err
	}
	log.Info("Character ability scores created", "ability_scores", as)

	skills, err := utils.HandleCharacterSkills(ctx, cl, ch, as)
	if err != nil {
		return nil, err
	}
	log.Info("Character skills created", "skills", skills)

	err = utils.HandleCharacterProficiencies(ctx, cl, ch, skills)
	if err != nil {
		return nil, err
	}
	log.Info("Character proficiencies created")

	return ch, nil
}

// UpdateCharacter is the resolver for the updateCharacter field.
func (r *mutationResolver) UpdateCharacter(ctx context.Context, id int, input ent.UpdateCharacterInput) (*ent.Character, error) {
	log.Info("Updating character", "id", id, "input", input)

	return r.Client.Character.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
