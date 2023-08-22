package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// EquipmentCategory holds the schema definition for the EquipmentCategory entity.
type EquipmentCategory struct {
	ent.Schema
}

func (EquipmentCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the EquipmentCategory.
func (EquipmentCategory) Fields() []ent.Field {
	return nil
}

// Edges of the EquipmentCategory.
func (EquipmentCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("category"),
	}
}
