package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CharacterAbilityScore struct {
	ent.Schema
}

func (CharacterAbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("score").Positive(),
		field.Int("modifier"),
	}
}

func (CharacterAbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("character_ability_scores").
			Unique().
			Required().
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("ability_score", AbilityScore.Type).
			Unique().
			Required(),
		edge.From("character_skills", CharacterSkill.Type).
			Ref("character_ability_score"),
	}
}

// func (CharacterAbilityScore) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.CharacterAbilityScoreFunc(func(ctx context.Context, m *gen.CharacterAbilityScoreMutation) (ent.Value, error) {
// 					log.Info("CharacterAbilityScore mutation -- set modifier hook")
// 					score, ok := m.Score()
// 					if !ok {
// 						return nil, errors.New("score is not set")
// 					}

// 					newMod := utils.AbilityScoreModifier(score)
// 					m.SetModifier(newMod)
// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			ent.OpCreate|ent.OpUpdateOne|ent.OpUpdateOne,
// 		),
// 	}
// }

// Annotations
func (CharacterAbilityScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
