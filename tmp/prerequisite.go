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
		field.Int("minimum"),
	}
}

// Edges of the Prerequisite.
func (Prerequisite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_score", AbilityScore.Type),
	}
}
