package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Armor holds the schema definition for the Armor entity.
type Armor struct {
	ent.Schema
}

// Fields of the Armor.
func (Armor) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("armor_category").
			Values(
				"light",
				"medium",
				"heavy",
				"shield",
			),
		field.Int("str_minimum"),
		field.Bool("stealth_disadvantage"),
		field.Int("ac_base").Positive(),
		field.Bool("ac_dex_bonus").Default(false),
		field.Int("ac_max_bonus").Default(0),
	}
}

// Edges of the Armor.
func (Armor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("armor").
			Unique().
			Required(),
	}
}
