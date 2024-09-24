package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WeaponRange is the schema of a WeaponRange
type WeaponRange struct {
	ent.Schema
}

// Fields of the WeaponRange.
func (WeaponRange) Fields() []ent.Field {
	return []ent.Field{
		field.Int("range_normal").
			Optional(),
		field.Int("range_long").
			Optional(),
		field.Int("throw_range_normal").
			Optional(),
		field.Int("throw_range_long").
			Optional(),
	}
}

// Weapon holds the schema definition for the Weapon entity.
type Weapon struct {
	ent.Schema
}

// Fields of the Weapon.
func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("weapon_category").
			Values("simple", "martial", "exotic", "other"),
		field.Enum("weapon_subcategory").
			Values("melee", "ranged", "other"),
	}
}

// Edges of the Weapon.
func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("damage", Damage.Type).
			Unique(),
		edge.To("properties", Property.Type),
		edge.From("equipment", Equipment.Type).
			Ref("weapon").
			Unique().
			Required(),
		edge.To("weapon_range", WeaponRange.Type).
			Unique(),
	}
}

// Annotations of the Weapon
func (Weapon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("weapons"),
	}
}
