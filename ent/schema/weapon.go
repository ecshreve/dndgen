package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type WeaponProperty struct {
	ent.Schema
}

func (WeaponProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (WeaponProperty) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("desc"),
	}
}

func (WeaponProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapons", Weapon.Type).
			Ref("weapon_properties"),
	}
}

//==========================================================
// Weapon
//==========================================================

type Weapon struct {
	ent.Schema
}

// Mixin of the Weapon.
func (Weapon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("equipment_id"),
		field.String("weapon_category"),
		field.String("weapon_range"),
	}
}

func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("weapon").
			Unique().Required().
			Field("equipment_id"),
		edge.To("damage_type", DamageType.Type).Through("weapon_damage", WeaponDamage.Type),
		edge.To("weapon_properties", WeaponProperty.Type),
	}
}

func (Weapon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
