package schema

import "entgo.io/ent"

// Trait holds the schema definition for the Trait entity.
type Trait struct {
	ent.Schema
}

// Mixin of the Trait.
func (Trait) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
