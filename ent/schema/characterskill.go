package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CharacterSkill struct {
	ent.Schema
}

// Annotations returns the annotations for the CharacterSkill.
//
//	func (CharacterSkill) Annotations() []schema.Annotation {
//		// return []schema.Annotation{
//		// 	field.ID("character_id", "skill_id"),
//		// 	entgql.QueryField(),
//		// }
//	}
//
// Index returns the index for the CharacterSkill.
func (CharacterSkill) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "skill_id").Unique(),
	}
}

func (CharacterSkill) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("proficient").
			Default(false),
		field.Int("modifier").
			Default(0),
		field.Int("character_id"),
		field.Int("skill_id"),
	}
}

func (CharacterSkill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Unique().
			Required().
			Field("character_id"),
		edge.To("skill", Skill.Type).
			Unique().
			Required().
			Field("skill_id"),
	}
}
