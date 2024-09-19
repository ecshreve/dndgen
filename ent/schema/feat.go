package schema

import "entgo.io/ent"

// Feat holds the schema definition for the Feat entity.
type Feat struct {
	ent.Schema
}

// Mixin implements the ent.Mixin for the Feat.
func (Feat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
