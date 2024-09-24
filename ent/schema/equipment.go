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

// Tool holds the schema definition for the Tool entity.
type Tool struct {
	ent.Schema
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("tool_category").Optional(),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type).
			Unique(),
	}
}

// Gear holds the schema definition for the Gear entity.
type Gear struct {
	ent.Schema
}

// Fields of the Gear.
func (Gear) Fields() []ent.Field {
	return []ent.Field{
		field.String("gear_category"),
	}
}

// Edges of the Gear.
func (Gear) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type).
			Unique(),
	}
}
