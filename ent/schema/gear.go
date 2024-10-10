package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Gear holds the schema definition for the Gear entity.
type Gear struct {
	ent.Schema
}

// Fields of the Gear.
func (Gear) Fields() []ent.Field {
	return []ent.Field{
		field.String("gear_category"),
		field.Strings("desc").Optional(),
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
