package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		// General
		field.String("name").Unique().NotEmpty(),
		field.Int("age").Positive(),
		// Attributes
		// field.Int("str").Range(1, 20),
		// field.Int("dex").Range(1, 20),
		// field.Int("con").Range(1, 20),
		// field.Int("int").Range(1, 20),
		// field.Int("wis").Range(1, 20),
		// field.Int("cha").Range(1, 20),
		// // Hit Points
		// field.Int("max_hit_points").Positive(),
		// field.Int("current_hit_points"),
		// // Experience Points
		// field.Int("experience_points").Default(0).NonNegative(),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{}
}
