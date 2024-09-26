package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/charmbracelet/log"
	gen "github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/hook"
	"github.com/ecshreve/dndgen/internal/utils"
)

// AbilityScore holds the schema definition for the AbilityScore entity.
type AbilityScore struct {
	ent.Schema
}

func (AbilityScore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Fields of the AbilityScore.
func (AbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Annotations(
			entgql.OrderField("FULL_NAME"),
		),
	}
}

// Edges of the AbilityScore.
func (AbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("skills", Skill.Type),
		edge.From("classes", Class.Type).
			Ref("saving_throws"),
		edge.From("characters", Character.Type).
			Ref("ability_scores").
			Through("character_ability_scores", CharacterAbilityScore.Type),
		edge.From("race", Race.Type).
			Ref("ability_bonuses").
			Through("race_ability_bonuses", AbilityBonus.Type),
	}
}

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
	}
}

// Hooks
func (CharacterAbilityScore) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CharacterAbilityScoreFunc(
					func(ctx context.Context, m *gen.CharacterAbilityScoreMutation) (ent.Value, error) {
						if newScore, ok := m.Score(); ok {
							oldMod, _ := m.Modifier()
							newMod := utils.AbilityScoreModifier(newScore)
							log.Debug("updating modifier", "oldMod", oldMod, "newMod", newMod, "score", newScore)
							m.SetModifier(newMod)
						}
						return next.Mutate(ctx, m)
					},
				)
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
