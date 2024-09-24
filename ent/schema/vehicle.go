package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("vehicle_category").
			Values(
				"mounts_and_other_animals",
				"tack_harness_and_drawn_vehicles",
				"waterborne",
			).
			Annotations(entgql.QueryField("vehicle_category")),
		field.String("capacity").Optional(),
		field.Float("speed_quantity").Optional(),
		field.Enum("speed_units").Optional().
			Values(
				"miles_per_hour",
				"feet_per_round",
			).
			Annotations(entgql.QueryField("speed_units")),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type).
			Unique(),
	}
}
