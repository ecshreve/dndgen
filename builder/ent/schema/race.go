package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Race holds the schema definition for the Race entity.
type Race struct {
	ent.Schema
}

func (Race) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Race.
func (Race) Fields() []ent.Field {
	return nil
}

// Edges of the Race.
func (Race) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("characters", Character.Type),
	}
}
