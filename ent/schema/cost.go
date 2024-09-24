package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cost holds the schema definition for the Cost entity.
type Cost struct {
	ent.Schema
}

// Fields of the Cost.
func (Cost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Default(1),
	}
}

// Edges of the Cost.
func (Cost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coin", Coin.Type).
			Unique().
			Required(),
		edge.From("equipment", Equipment.Type).
			Ref("cost").
			Unique(),
	}
}
