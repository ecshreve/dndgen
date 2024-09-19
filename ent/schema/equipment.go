package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Mixin of the Equipment.
func (Equipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("equipment_category").
			Values("armor", "gear", "vehicles", "tools", "weapon", "other").
			Annotations(entgql.QueryField("equipment_category")),
		field.Float("weight"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment_costs", EquipmentCost.Type).
			Unique(),
	}
}

// Annotations of the Equipment
func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("equipments"),
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
