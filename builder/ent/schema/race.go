package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
	return []ent.Field{
		field.Int("speed").NonNegative(),
		field.Enum("size").Values("Tiny", "Small", "Medium", "Large", "Huge", "Gargantuan"),
		field.String("size_description"),
		field.String("age"),
	}
}

// Edges of the Race.
func (Race) Edges() []ent.Edge {
	return []ent.Edge{}
}
