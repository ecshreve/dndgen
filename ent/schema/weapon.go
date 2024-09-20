package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Damage holds the schema definition for the Damage entity.
type Damage struct {
	ent.Schema
}

// Fields of the Damage.
func (Damage) Fields() []ent.Field {
	return []ent.Field{
		field.String("damage_dice"),
	}
}

// Edges of the Damage.
func (Damage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("damage_type", DamageType.Type).
			Unique(),
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
			Values("simple", "martial", "exotic", "other").
			Annotations(entgql.QueryField("weapon_category")),
		field.Enum("weapon_subcategory").
			Values("melee", "ranged", "other").
			Annotations(entgql.QueryField("weapon_subcategory")),
	}
}

// Edges of the Weapon.
func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("damage", Damage.Type).
			Unique(),
		edge.To("properties", Property.Type),
		edge.To("equipment", Equipment.Type).
			Unique(),
	}
}
