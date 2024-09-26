package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Coin holds the schema definition for the Coin entity.
type Coin struct {
	ent.Schema
}

// Mixin returns the Coin mixed-in schema.
func (Coin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Fields of the Coin.
func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.Float("gold_conversion_rate"),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("costs", Cost.Type).
			Ref("coin").
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
	}
}
