package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			Values("armor", "adventuring-gear", "mounts-and-vehicles", "tools", "weapon", "other"),
		field.Float("weight"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment_costs", EquipmentCost.Type).
			StorageKey(edge.Column("equipment_id")),
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
		field.Int("equipment_id"),
		field.Int("coin_id"),
	}
}

// Edges of the EquipmentCost.
func (EquipmentCost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coin", Coin.Type).
			Unique().
			Required().
			Field("coin_id"),
		edge.To("equipment", Equipment.Type).
			Unique().
			Required().
			Field("equipment_id"),
	}
}

// Index of the EquipmentCost.
func (EquipmentCost) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("equipment_id").Unique(),
	}
}
