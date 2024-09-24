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
		edge.From("equipment_costs", EquipmentCost.Type).
			Ref("coin").Annotations(
			entgql.Skip(entgql.SkipAll),
		),
	}
}

// EquipmentCost holds the schema definition for the EquipmentCost entity.
type EquipmentCost struct {
	ent.Schema
}

// Fields of the EquipmentCost.
func (EquipmentCost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Default(1),
	}
}

// Edges of the EquipmentCost.
func (EquipmentCost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coin", Coin.Type).
			Unique().
			Required(),
		edge.From("equipment", Equipment.Type).
			Ref("equipment_costs").
			Unique(),
	}
}
