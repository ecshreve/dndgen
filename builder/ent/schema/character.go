package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		// General
		field.Int("level").Default(1).Positive(),
		field.String("alignment").Optional(),
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
	return []ent.Edge{
		edge.From("race", Race.Type).
			Ref("characters").
			Unique().
			Required(),
		edge.From("class", Class.Type).
			Ref("characters").
			Unique().
			Required(),
	}
}
