package schema

import (
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
			),
		field.String("capacity").Optional(),
		field.Strings("desc").Optional(),
		field.Float("speed_quantity").Optional(),
		field.Enum("speed_units").Optional().
			Values(
				"miles_per_hour",
				"feet_per_round",
			),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("vehicle").
			Unique().
			Required(),
	}
}
