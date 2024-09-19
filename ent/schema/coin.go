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

// Mixin implements the ent.Mixin for the Coin.
func (Coin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
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
		edge.To("equipment_costs", EquipmentCost.Type).
			StorageKey(edge.Column("coin_id")).
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
	}
}
