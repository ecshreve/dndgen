package schema

import (
	"entgo.io/ent"
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
	}
}
