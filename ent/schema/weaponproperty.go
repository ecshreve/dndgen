package schema

import "entgo.io/ent"

// WeaponProperty holds the schema definition for the WeaponProperty entity.
type WeaponProperty struct {
	ent.Schema
}

// Mixin of the WeaponProperty.
func (WeaponProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
