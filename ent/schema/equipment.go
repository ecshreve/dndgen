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
			Values(
				"armor",
				"gear",
				"vehicles",
				"tools",
				"weapon",
				"other",
			),
		field.Float("weight"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cost", Cost.Type).
			Unique(),
		edge.To("tool", Tool.Type).
			Unique(),
		edge.To("gear", Gear.Type).
			Unique(),
		edge.To("armor", Armor.Type).
			Unique(),
		edge.To("weapon", Weapon.Type).
			Unique(),
		edge.To("vehicle", Vehicle.Type).
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
		field.String("tool_category"),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("tool").
			Unique().
			Required(),
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
		edge.From("equipment", Equipment.Type).
			Ref("gear").
			Unique().
			Required(),
	}
}
