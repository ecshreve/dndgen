package schema

import (
	"entgo.io/ent"
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
	return nil
}

// Edges of the DamageType.
func (DamageType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapon_damage", WeaponDamage.Type).Ref("damage_type").Unique(),
	}
}

type WeaponDamage struct {
	ent.Schema
}

func (WeaponDamage) Fields() []ent.Field {
	return []ent.Field{
		field.String("dice"),
	}
}

func (WeaponDamage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("damage_type", DamageType.Type),
		edge.From("weapon", Weapon.Type).Ref("damage"),
	}
}
