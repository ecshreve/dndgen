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

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("INDX"),
			),
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Enum("equipment_category").
			NamedValues(
				"gear", "GEAR",
				"tool", "TOOL",
				"weapon", "WEAPON",
				"vehicle", "VEHICLE",
				"armor", "ARMOR",
			).
			Annotations(
				entgql.OrderField("EQUIPMENT_CATEGORY"),
			),
		field.Float("weight").Optional(),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cost", Cost.Type).
			Unique(),
		edge.To("gear", Gear.Type).
			Unique(),
		edge.To("tool", Tool.Type).
			Unique(),
		edge.To("weapon", Weapon.Type).
			Unique(),
		edge.To("vehicle", Vehicle.Type).
			Unique(),
		edge.To("armor", Armor.Type).
			Unique(),
		edge.From("proficiencies", Proficiency.Type).
			Ref("equipment"),
	}
}

// Annotations of the Equipment
func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("equipments"),
		entgql.RelayConnection(),
	}
}
