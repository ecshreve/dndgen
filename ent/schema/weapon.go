package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Weapon holds the schema definition for the Weapon entity.
type Weapon struct {
	ent.Schema
}

// Fields of the Weapon.
func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("weapon_category").
			Values(
				"simple",
				"martial",
				"exotic",
				"other",
			),
		field.Enum("weapon_subcategory").
			Values(
				"melee",
				"ranged",
				"other",
			),
		field.Int("range_normal").
			Optional(),
		field.Int("range_long").
			Optional(),
		field.Int("throw_range_normal").
			Optional(),
		field.Int("throw_range_long").
			Optional(),
		field.String("damage_dice").
			Optional(),
	}
}

// Edges of the Weapon.
func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("properties", Property.Type),
		edge.To("damage_type", DamageType.Type).
			Unique(),
		edge.From("equipment", Equipment.Type).
			Ref("weapon").
			Unique().
			Required(),
	}
}
