package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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

func (DamageType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapons", Weapon.Type).
			Ref("damage_type"),
	}
}
