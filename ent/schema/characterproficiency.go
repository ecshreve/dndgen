package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CharacterProficiency struct {
	ent.Schema
}

// Indexes
func (CharacterProficiency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "proficiency_id").Unique(),
	}
}

func (CharacterProficiency) Fields() []ent.Field {
	return []ent.Field{
		field.Int("character_id"),
		field.Int("proficiency_id"),
	}
}

func (CharacterProficiency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Unique().
			Required().
			Field("character_id"),
		edge.To("proficiency", Proficiency.Type).
			Unique().
			Required().
			Field("proficiency_id"),
	}
}
