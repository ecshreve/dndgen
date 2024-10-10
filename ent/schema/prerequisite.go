package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Prerequisite holds the schema definition for the Prerequisite entity.
type Prerequisite struct {
	ent.Schema
}

// Fields of the Prerequisite.
func (Prerequisite) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("prerequisite_type").
			StructTag(`json:"type,omitempty"`).
			Values(
				"level",
				"spell",
				"feature",
			),
		field.Int("level_value").
			StructTag(`json:"level,omitempty"`).
			Optional(),
		field.String("feature_value").
			StructTag(`json:"feature,omitempty"`).
			Optional(),
		field.String("spell_value").
			StructTag(`json:"spell,omitempty"`).
			Optional(),
	}
}

// Edges of the Prerequisite.
func (Prerequisite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Feature.Type).
			Ref("prerequisites").
			Unique(),
	}
}
