package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CharacterSkill struct {
	ent.Schema
}

func (CharacterSkill) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("proficient").Default(false),
	}
}

func (CharacterSkill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("character_skills").
			Unique().
			Required(),
		edge.To("skill", Skill.Type).
			Unique().
			Required(),
		edge.To("character_ability_score", CharacterAbilityScore.Type).
			Unique().
			Required(),
		edge.To("character_proficiency", CharacterProficiency.Type).
			Unique(),
	}
}

// Hooks

// func (CharacterSkill) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return hook.CharacterSkillFunc(func(ctx context.Context, m *gen.CharacterSkillMutation) (ent.Value, error) {
// 					log.Info("CharacterSkill mutation -- set proficient hook")
// 					_, ok := m.CharacterProficiencyID()
// 					m.SetProficient(ok)

// 					return next.Mutate(ctx, m)
// 				})
// 			},
// 			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
// 		),
// 	}
// }

func (CharacterSkill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
