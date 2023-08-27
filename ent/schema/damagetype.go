package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DamageType holds the schema definition for the DamageType entity.
type DamageType struct {
	ent.Schema
}

func (DamageType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the DamageType.
func (DamageType) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("desc"),
	}
}

// Edges of the DamageType.
func (DamageType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapon_damage", WeaponDamage.Type).
			Ref("damage_type"),
	}
}

// Annotations of the DamageType.
func (DamageType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

type WeaponDamage struct {
	ent.Schema
}

func (WeaponDamage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weapon_id"),
		field.Int("damage_type_id"),
		field.String("dice"),
	}
}

func (WeaponDamage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapon", Weapon.Type).
			Ref("weapon_damage").
			Unique().
			Required().
			Field("weapon_id"),
		edge.To("damage_type", DamageType.Type).
			Unique().
			Required().
			Field("damage_type_id"),
	}
}

// Annotations of the WeaponDamage.
func (WeaponDamage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
