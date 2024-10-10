package schema

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/charmbracelet/log"
	gen "github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/hook"
	"github.com/ecshreve/dndgen/internal/utils"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
		field.Int("age").Positive().Default(25),
		field.Int("level").Positive().Default(1).Annotations(
			entgql.OrderField("LEVEL"),
		),
		field.Int("proficiency_bonus").Positive().Default(2),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("race", Race.Type).Unique(),
		edge.To("class", Class.Type).Unique(),
		edge.To("alignment", Alignment.Type).Unique(),
		// edge.To("traits", Trait.Type),
		// edge.To("languages", Language.Type),
		// edge.To("proficiencies", Proficiency.Type).
		// 	Through("character_proficiencies", CharacterProficiency.Type).
		// 	Annotations(
		// 		entgql.QueryField(),
		// 	),
		edge.To("character_ability_scores", CharacterAbilityScore.Type),
		edge.To("character_skills", CharacterSkill.Type),
		edge.To("character_proficiencies", CharacterProficiency.Type),
	}
}

// Annotations returns the annotations for the Character.
func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Character) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			func(next ent.Mutator) ent.Mutator {
				return hook.CharacterFunc(func(ctx context.Context, m *gen.CharacterMutation) (ent.Value, error) {
					log.Info("Character mutation -- set proficiency bonus hook")
					level, ok := m.Level()
					if !ok {
						return nil, errors.New("level is not set")
					}
					bonus := utils.LevelProficiencyBonus(level)
					m.SetProficiencyBonus(bonus)
					return next.Mutate(ctx, m)
				})
			},
			hook.HasFields("level"),
		),
	}
}
