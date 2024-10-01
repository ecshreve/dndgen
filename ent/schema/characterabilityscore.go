package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/charmbracelet/log"

	gen "github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/hook"
	"github.com/ecshreve/dndgen/internal/utils"
)

type CharacterAbilityScore struct {
	ent.Schema
}

// Indexes
func (CharacterAbilityScore) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "ability_score_id").Unique(),
	}
}

func (CharacterAbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("score").Positive(),
		field.Int("modifier").Range(-5, 10),
		field.Int("character_id"),
		field.Int("ability_score_id"),
	}
}

func (CharacterAbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Unique().
			Required().
			Field("character_id"),
		edge.To("ability_score", AbilityScore.Type).
			Unique().
			Required().
			Field("ability_score_id"),
		edge.To("character_skills", CharacterSkill.Type),
	}
}

func (CharacterAbilityScore) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			func(next ent.Mutator) ent.Mutator {
				return hook.CharacterAbilityScoreFunc(func(ctx context.Context, m *gen.CharacterAbilityScoreMutation) (ent.Value, error) {
					log.Info("CharacterAbilityScore mutation -- set modifier hook")
					score, ok := m.Score()
					if !ok {
						return nil, errors.New("score is not set")
					}

					newMod := utils.AbilityScoreModifier(score)
					m.SetModifier(newMod)
					return next.Mutate(ctx, m)
				})
			},
			hook.HasFields("score"),
		),
	}
}
